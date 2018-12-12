package simulator

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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
	r.HandleFunc("/{key}/", h.handleGet).Methods("GET")
	r.HandleFunc("/{key}/", h.handleSet).Methods("POST")
	r.HandleFunc("/{key}/", h.handleDelete).Methods("DELETE")
	r.HandleFunc("/{key}/watch", h.handleWatch).Methods("GET")
	h.router = r
}

func (h *HTTPSimulator) handleGet(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	result, err := h.Gimulator.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *HTTPSimulator) handleSet(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	var object interface{}
	if err := json.NewDecoder(r.Body).Decode(&object); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Gimulator.Set(key, object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *HTTPSimulator) handleDelete(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	err := h.Gimulator.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *HTTPSimulator) handleWatch(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

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
			log.Println("SENDING", r)
			w.conn.WriteJSON(r)
		case <-t.C:
			w.conn.WriteMessage(websocket.PingMessage, []byte{})
		}
	}
}
