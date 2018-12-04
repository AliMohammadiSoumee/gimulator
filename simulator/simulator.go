package simulator

import (
	"net/http"
	"sync"

	"github.com/alidadar7676/gimulator/game"
	"github.com/alidadar7676/gimulator/types"
	"github.com/gorilla/websocket"
)

type Simulator struct {
	judger game.Judger

	conns []*websocket.Conn
	state types.World

	mu sync.Mutex
}

func (s *Simulator) Run() error {
	return http.ListenAndServe("0.0.0.0:8585", s)
}

func (s *Simulator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// FIXME: 1024
	conn, err := websocket.Upgrade(w, r, r.Header, 1024, 1024)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.mu.Lock()
	s.conns = append(s.conns, conn)
	s.mu.Unlock()
}

func (s *Simulator) broke(input interface{}) error {
	for _, conn := range s.conns {
		if err := conn.WriteJSON(input); err != nil {
			return err
		}
	}
	return nil
}
