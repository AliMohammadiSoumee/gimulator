package main

import (
	"log"

	"github.com/alidadar7676/gimulator/simulator"
	"github.com/alidadar7676/gimulator/types"
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
	worldFilter := simulator.Object{
		Key: simulator.Key{
			Type:      types.WorldType,
			Namespace: c.Namespace,
		}}

	worlds, err := c.gimulator.Find(worldFilter)
	switch {
	case err != nil:
		return
	case len(worlds) > 1:
		return
	case len(worlds) == 1:
		var world types.World
		if err := worlds[0].Struct(&world); err != nil {
			return
		}
		c.watchWorld(world)
	}

	c.gimulator.Watch(worldFilter, c.watcher)

	go func() {
		for r := range c.watcher {
			var world types.World
			if err := r.Object.Struct(&world); err != nil {
				log.Printf("object %v is not world\n", r.Object)
				continue
			}
			c.watchWorld(world)
		}
	}()
}

func (c *Controller) watchWorld(world types.World) {
	d := worldDrawer{World: world}
	render(d)
	disableEvent = world.Turn != playerName
}

func (c *Controller) InitPlayer(playerName string) error {
	playerIntroObject := simulator.Object{
		Key: simulator.Key{
			Type:      types.PlayerIntroType,
			Name:      playerName,
			Namespace: c.Namespace,
		},
		Value: types.PlayerIntro{},
	}
	return c.gimulator.Set(playerIntroObject)
}
