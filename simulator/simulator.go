package simulator

import (
	"fmt"
)

const (
	TaskBufferSize = 256
)

const (
	SetAction = "set"
	GetAction = "get"
)

type Reconcile struct {
	Action string
	Object Object
}

type Simulator struct {
	storage  map[Key]Object
	tasks    chan task
	watchers []watcher
	running  bool
}

var _ Gimulator = (*Simulator)(nil)

func NewSimulator() *Simulator {
	return &Simulator{
		storage:  make(map[Key]Object),
		tasks:    make(chan task, TaskBufferSize),
		watchers: make([]watcher, 0),
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
	obj, ok := result.value.(Object)
	if !ok {
		return nil, fmt.Errorf("unexpected result from get")
	}
	return &obj, result.err
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

func (s *Simulator) Watch(filter Object, ch chan Reconcile) error {
	<-s.send(msgWatch{filter: filter, ch: ch})
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
		s.watch(msg.filter, msg.ch)
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
		if matchObject(filter, object) {
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

func (s *Simulator) watch(filter Object, ch chan Reconcile) {
	s.watchers = append(s.watchers, watcher{
		filter: filter,
		ch:     ch,
	})
}

func (s *Simulator) reconcile(reconcile Reconcile) {
	for _, w := range s.watchers {
		if matchObject(w.filter, reconcile.Object) {
			select {
			case w.ch <- reconcile:
			default:
			}
		}
	}
}

type watcher struct {
	filter Object
	ch     chan Reconcile
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
	filter Object
	ch     chan Reconcile
}
