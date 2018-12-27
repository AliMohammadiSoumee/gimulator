package agent

import (
	"log"
	"github.com/alidadar7676/gimulator/simulator"
	"github.com/alidadar7676/gimulator/types"
	"runtime"
	"fmt"
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
	playerIntroObject := simulator.Object{
		Key: simulator.Key{
			Type:      types.PlayerIntroType,
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Value: types.PlayerIntro{},
	}
	if err := c.gimulator.Set(playerIntroObject); err != nil {
		log.Fatalf("can not create playerIntro %v: %v\n", playerIntroObject, err)
	}

	c.gimulator.Watch(simulator.Object{
		Key: simulator.Key{
			Type:      types.WorldType,
			Namespace: c.Namespace,
		}}, c.watcher)

	go func() {
		for r := range c.watcher {
			var world types.World
			if err := r.Object.Struct(&world); err != nil {
				log.Fatalf("object %v has not world type: %v\n", r.Object, err)
			}
			if world.Turn == c.Name {
				c.Act(world)
			}
		}
	}()
}

func (c *Controller) Act(world types.World) {
	move := run(world, c.Name)

	PrintMemory()

	actionObject := simulator.Object{
		Key: simulator.Key{
			Type:      types.ActionType,
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Value: types.Action{
			PlayerName: c.Name,
			From:       move.A,
			To:         move.B,
		},
	}

	err := c.gimulator.Set(actionObject)
	if err != nil {
		log.Fatalf("can not set action object %v: %v\n", actionObject, err)
	}
}

func PrintMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v\n", m.Alloc / 1024 / 1024)
}
