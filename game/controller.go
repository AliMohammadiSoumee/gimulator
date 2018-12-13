package game

import (
	"fmt"
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
		watcher:   make(chan simulator.Reconcile, 32),
	}
}

func (c *Controller) Run() {
	watchingList := []simulator.Object{
		simulator.Object{
			Key: simulator.Key{
				Type:      PlayerIntroType,
				Namespace: c.Namespace,
			},
		},
		simulator.Object{
			Key: simulator.Key{
				Type:      ActionType,
				Namespace: c.Namespace,
			},
		},
	}
	for _, item := range watchingList {
		err := c.gimulator.Watch(item, c.watcher)
		if err != nil {
			log.Fatalf("Error watching %v: %v", item, err)
		}
	}

	go func() {
		for r := range c.watcher {
			err := c.reconcile(r)
			if err != nil {
				log.Printf("Error reconcile %v: %v\n", r, err)
			}
		}
	}()
}

func (c *Controller) reconcile(r simulator.Reconcile) error {
	switch {
	case r.Object.Type == PlayerIntroType:
		return c.playerJoined()
	case r.Object.Type == ActionType:
		r.Object.Struct(&Action{})
		return c.playerActed(r.Object)
	}
	return nil
}

func (c *Controller) playerJoined() error {
	worldKey := simulator.Key{
		Type:      WorldType,
		Name:      c.Name,
		Namespace: c.Namespace,
	}
	err := c.gimulator.Get(worldKey, &simulator.Object{})
	if err == nil {
		// World already created
		return nil
	}

	players, err := c.gimulator.Find(simulator.Object{
		Key: simulator.Key{
			Namespace: c.Namespace,
			Type:      PlayerIntroType,
		}})
	if err != nil {
		return err
	}

	if len(players) < 2 {
		return nil
	}

	playerName1 := players[0].Name
	playerName2 := players[1].Name

	world := NewWorld(playerName1, playerName2)
	worldObject := simulator.Object{
		Key: simulator.Key{
			Type:      WorldType,
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Value: world,
	}

	err = c.gimulator.Set(worldObject)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) playerActed(actionObject simulator.Object) error {
	action, ok := actionObject.Value.(Action)
	if !ok {
		return fmt.Errorf("type of %v is not Action", actionObject)
	}

	worldKey := simulator.Key{
		Name:      c.Name,
		Namespace: c.Namespace,
		Type:      WorldType,
	}
	worldObject := simulator.Object{}
	if err := c.gimulator.Get(worldKey, &worldObject); err != nil {
		return fmt.Errorf("can not get world: %v", err)
	}
	var world World
	if err := worldObject.Struct(&world); err != nil {
		return fmt.Errorf("type of %v is not world: %v", worldObject, err)
	}

	if action.PlayerName != world.Player1.Name &&
		action.PlayerName != world.Player2.Name {
		return fmt.Errorf("invalid action player name %q: not in {%q, %q}",
			action.PlayerName, world.Player1.Name, world.Player2.Name)
	}

	updatedWorld := Update(action, world)
	updatedWorldObject := simulator.Object{
		Key:   worldKey,
		Value: updatedWorld,
	}
	if err := c.gimulator.Set(updatedWorldObject); err != nil {
		return fmt.Errorf("can not set %v object: %v", updatedWorldObject, err)
	}

	return nil
}
