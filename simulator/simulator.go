package simulator

import (
	"fmt"
)

const (
	TaskBufferSize = 256
)

type Key struct {
	Namespace string
	Type      string
	Name      string
}

type Object struct {
	Key
	Value interface{}
}

type Reconcile struct {
	Action string
	Object Object
}

type Simulator struct {
	storage map[Key]Object
	tasks   chan task
	watcher map[Key][]chan Reconcile
	running bool
}

var _ Gimulator = (*Simulator)(nil)

func NewSimulator() *Simulator {
	return &Simulator{
		storage: make(map[Key]Object),
		tasks:   make(chan task, TaskBufferSize),
		watcher: make(map[Key][]chan Reconcile),
	}
}

func (s *Simulator) Run() {
	go func() {
		s.running = true
		defer func() {
			s.running = false
		}()

		for t := range s.tasks {
			s.loop(t)
		}
	}()
}

func (s *Simulator) Get(key Key) (*Object, error) {
	result := <-s.send(msgGet{key: key})
	object, ok := result.value.(Object)
	if !ok {
		return nil, fmt.Errorf("unexpected result from get")
	}
	return &object, result.err
}

func (s *Simulator) Find(filter Object) ([]Object, error) {
	result := <-s.send(msgFind{filter: filter})
	if result.err != nil {
		return nil, result.err
	}
	objectList, ok := result.value.([]Object)
	if !ok {
		return nil, fmt.Errorf("unexpected result from find")
	}
	return objectList, nil
}

func (s *Simulator) Set(object Object) error {
	<-s.send(msgSet{key: object.Key, object: object})
	return nil
}

func (s *Simulator) Delete(key Key) error {
	result := <-s.send(msgDelete{key: key})
	return result.err
}

func (s *Simulator) Watch(key Key, ch chan Reconcile) error {
	<-s.send(msgWatch{key: key, ch: ch})
	return nil
}

func (s *Simulator) send(msg interface{}) chan result {
	if !s.running {
		panic("can not send message: simulator is not running")
	}

	ch := make(chan result)
	s.tasks <- task{
		input:  msg,
		result: ch,
	}
	return ch
}

func (s *Simulator) loop(t task) {
	var (
		result  interface{}
		err     error
		changed *Reconcile
	)

	switch msg := t.input.(type) {
	case msgGet:
		result, err = s.get(msg.key)
	case msgFind:
		result, err = s.find(msg.filter)
	case msgSet:
		s.set(msg.key, msg.object)
		changed = &Reconcile{
			Action: "set",
			Object: msg.object,
		}
		result, err = nil, nil
	case msgDelete:
		changed = &Reconcile{
			Action: "delete",
		}
		result, err = nil, s.delete(msg.key)
	case msgWatch:
		s.watch(msg.key, msg.ch)
		result, err = nil, nil
	default:
		result, err = nil, fmt.Errorf("undefined message type")
	}

	t.response(result, err)
	if changed != nil {
		s.reconcile(*changed)
	}
}

func (s *Simulator) get(key Key) (Object, error) {
	if object, exists := s.storage[key]; exists {
		return object, nil
	}
	return Object{}, fmt.Errorf("object with %v key does not exist", key)
}

func (s *Simulator) find(filter Object) ([]Object, error) {
	result := make([]Object, 0)
	for _, object := range s.storage {
		fmt.Println(object, "------", filter)
		if filter.Namespace != "" && filter.Namespace != object.Namespace {
			continue
		}
		if filter.Type != "" && filter.Type != object.Type {
			continue
		}
		if filter.Name != "" && filter.Name != object.Name {
			continue
		}
		if match(filter.Value, object.Value) {
			result = append(result, object)
		}
	}
	return result, nil
}

func (s *Simulator) set(key Key, object Object) {
	s.storage[key] = object
}

func (s *Simulator) delete(key Key) error {
	if _, exists := s.storage[key]; !exists {
		return fmt.Errorf("object with %v key does not exist", key)
	}
	delete(s.storage, key)
	return nil
}

func (s *Simulator) watch(key Key, ch chan Reconcile) {
	if _, exists := s.watcher[key]; !exists {
		s.watcher[key] = make([]chan Reconcile, 0)
	}
	s.watcher[key] = append(s.watcher[key], ch)
}

func (s *Simulator) reconcile(reconcile Reconcile) {
	channels, exists := s.watcher[reconcile.Object.Key]
	if !exists {
		return
	}
	for _, ch := range channels {
		select {
		case ch <- reconcile:
		default:
		}
	}
}

type result struct {
	value interface{}
	err   error
}

type task struct {
	input  interface{}
	result chan result
}

func (t task) response(value interface{}, err error) {
	select {
	case t.result <- result{value: value, err: err}:
	default:
	}
}

type msgGet struct {
	key Key
}

type msgSet struct {
	key    Key
	object Object
}

type msgDelete struct {
	key Key
}

type msgFind struct {
	filter Object
}

type msgWatch struct {
	key Key
	ch  chan Reconcile
}
