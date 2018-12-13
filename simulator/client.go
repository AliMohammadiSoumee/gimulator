package simulator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type Client struct {
	Addr string
}

var _ Gimulator = (*Client)(nil)

func (c *Client) Get(key Key, object *Object) error {
	req, err := http.NewRequest("GET", c.url("GET", key), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unsuccessful request")
	}

	if err := json.NewDecoder(resp.Body).Decode(object); err != nil {
		return err
	}

	return nil
}

func (c *Client) Find(filter Object) ([]Object, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(filter); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url("FIND", Key{}), buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessful request")
	}

	var objectList []Object
	if err := json.NewDecoder(resp.Body).Decode(&objectList); err != nil {
		return nil, err
	}

	return objectList, nil
}

func (c *Client) Set(object Object) error {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(object); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.url("SET", object.Key), buf)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unsuccessful request")
	}
	return nil
}

func (c *Client) Delete(key Key) error {
	req, err := http.NewRequest("DELETE", c.url("DELETE", key), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unsuccessful request")
	}
	return nil
}

func (c *Client) Watch(filter Object, ch chan Reconcile) error {
	ws, _, err := websocket.DefaultDialer.Dial(c.url("WATCH", Key{}), nil)
	if err != nil {
		return err
	}

	if err := ws.WriteJSON(filter); err != nil {
		ws.Close()
		return err
	}

	go func() {
		defer close(ch)
		defer ws.Close()
		for {
			var reconcile Reconcile
			err := ws.ReadJSON(&reconcile)
			if err != nil {
				continue
			}
			ch <- reconcile
		}
	}()

	return nil
}

func (c *Client) url(action string, key Key) string {
	var u url.URL
	switch strings.ToUpper(action) {
	case "GET", "SET", "DELETE":
		u = url.URL{
			Scheme: "http",
			Host:   c.Addr,
			Path:   fmt.Sprintf("/%s/%s/%s", key.Namespace, key.Type, key.Name),
		}
	case "FIND":
		u = url.URL{Scheme: "http", Host: c.Addr, Path: "/find"}
	case "WATCH":
		u = url.URL{
			Scheme: "ws",
			Host:   c.Addr,
			Path:   "/watch",
		}
	default:
		panic("unknown action")
	}
	return u.String()
}
