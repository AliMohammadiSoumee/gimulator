package main

import (
	"log"

	"github.com/alidadar7676/gimulator/simulator"
)

type Controller struct {
	Name      string
	Namespace string

	gimulator simulator.Gimulator
	watcher   chan simulator.Reconcile
}

func NewController(name, namespace string, gimulator simulator.Gimulator) *Controller {
	return &Controller{
		Name:      name,
		Namespace: namespace,
		gimulator: gimulator,
		watcher:   make(chan simulator.Reconcile, 1024),
	}
}

func (c *Controller) Run() {
	c.gimulator.Watch(simulator.Object{
		Key: simulator.Key{
			Type:      WorldType,
			Namespace: c.Namespace,
		},
	}, c.watcher)

	go func() {
		for r := range c.watcher {
			var world World
			if err := r.Object.Struct(&world); err != nil {
				log.Printf("object %v is not world\n", r.Object)
				continue
			}
			c.watchWorld(world)
		}
	}()
}

func (c *Controller) watchWorld(world World) {
	drawer.World = world
	render(drawer)
}
