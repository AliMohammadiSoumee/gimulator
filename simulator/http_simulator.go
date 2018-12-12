package simulator

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type HTTPSimulator struct {
	Gimulator Gimulator
	router    *mux.Router
}

var _ http.Handler = (*HTTPSimulator)(nil)

func NewHTTPSimulator(gimulator Gimulator) *HTTPSimulator {
	h := &HTTPSimulator{
		Gimulator: gimulator,
	}
	h.setRouter()
	return h
}

func (h *HTTPSimulator) ListenAndServe(bind string) error {
	if h.router == nil {
		h.setRouter()
	}
	return http.ListenAndServe(bind, h)
}

func (h *HTTPSimulator) setRouter() {
	r := mux.NewRouter()

	// Get
	r.HandleFunc("/{namespace}/{type}/{name}", h.handleGet).Methods("GET")

	// Find
	r.HandleFunc("/{namespace}/{type}/find", h.handleFind).Methods("POST")
	r.HandleFunc("/{namespace}/find", h.handleFind).Methods("POST")
	r.HandleFunc("/find", h.handleFind).Methods("POST")

	// Set
	r.HandleFunc("/{namespace}/{type}/{name}", h.handleSet).Methods("POST")

	// Delete
	r.HandleFunc("/{namespace}/{type}/{name}", h.handleDelete).Methods("DELETE")

	// Watch
	r.HandleFunc("/{namespace}/{type}/{name}/watch", h.handleWatch).Methods("GET")

	h.router = r
}

func (h *HTTPSimulator) handleGet(w http.ResponseWriter, r *http.Request) {
	key := Key{
		Namespace: mux.Vars(r)["namespace"],
		Type:      mux.Vars(r)["type"],
		Name:      mux.Vars(r)["name"],
	}
	result, err := h.Gimulator.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *HTTPSimulator) handleFind(w http.ResponseWriter, r *http.Request) {
	var filter Object
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	objectList, err := h.Gimulator.Find(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(objectList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *HTTPSimulator) handleSet(w http.ResponseWriter, r *http.Request) {
	var object Object
	if err := json.NewDecoder(r.Body).Decode(&object); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Gimulator.Set(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *HTTPSimulator) handleDelete(w http.ResponseWriter, r *http.Request) {
	key := Key{
		Namespace: mux.Vars(r)["namespace"],
		Type:      mux.Vars(r)["type"],
		Name:      mux.Vars(r)["name"],
	}

	err := h.Gimulator.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *HTTPSimulator) handleWatch(w http.ResponseWriter, r *http.Request) {
	key := Key{
		Namespace: mux.Vars(r)["namespace"],
		Type:      mux.Vars(r)["type"],
		Name:      mux.Vars(r)["name"],
	}

	ch := make(chan Reconcile, 32)
	err := h.Gimulator.Watch(key, ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	watcher := &HTTPWatcher{
		conn: conn,
		ch:   ch,
	}
	go watcher.Run()
}

func (h *HTTPSimulator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type HTTPWatcher struct {
	conn *websocket.Conn
	ch   chan Reconcile
}

func (w *HTTPWatcher) Run() {
	defer w.conn.Close()
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()

	for {
		select {
		case r := <-w.ch:
			w.conn.WriteJSON(r)
		case <-t.C:
			w.conn.WriteMessage(websocket.PingMessage, []byte{})
		}
	}
}
