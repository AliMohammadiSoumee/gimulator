package game

import (
	"fmt"
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
		watcher:   make(chan simulator.Reconcile, 32),
	}
}

func (c *Controller) Run() {
	watchingList := []simulator.Object{
		simulator.Object{
			Key: simulator.Key{
				Type:      types.PlayerIntroType,
				Namespace: c.Namespace,
			},
		},
		simulator.Object{
			Key: simulator.Key{
				Type:      types.ActionType,
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
	log.Println("Start reconciling...")
	switch {
	case r.Object.Type == types.PlayerIntroType:
		return c.playerJoined()
	case r.Object.Type == types.ActionType:
		r.Object.Struct(&types.Action{})
		return c.playerActed(r.Object)
	}
	return nil
}

func (c *Controller) playerJoined() error {
	log.Println("Register Players")
	worldKey := simulator.Key{
		Type:      types.WorldType,
		Name:      c.Name,
		Namespace: c.Namespace,
	}
	_, err := c.gimulator.Get(worldKey)
	if err == nil {
		// World already created
		return nil
	}

	players, err := c.gimulator.Find(simulator.Object{
		Key: simulator.Key{
			Namespace: c.Namespace,
			Type:      types.PlayerIntroType,
		},
	})
	if err != nil {
		log.Printf("Error Registering Players:Cannot find %v\n", err)
		return err
	}

	if len(players) < 2 {
		log.Println("Number of Players are less than 2")
		return nil
	}

	playerName1 := players[0].Name
	playerName2 := players[1].Name

	world := types.NewWorld(playerName1, playerName2)
	worldObject := simulator.Object{
		Key: simulator.Key{
			Type:      types.WorldType,
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Value: world,
	}
	log.Printf("Two players are reconciled with names: %s, %s", playerName1, playerName2)

	err = c.gimulator.Set(worldObject)
	if err != nil {
		log.Println("Can not set object world")
		return err
	}

	return nil
}

func (c *Controller) playerActed(actionObject simulator.Object) error {
	action, ok := actionObject.Value.(types.Action)
	if !ok {
		return fmt.Errorf("type of %v is not Action", actionObject)
	}

	worldKey := simulator.Key{
		Name:      c.Name,
		Namespace: c.Namespace,
		Type:      types.WorldType,
	}
	worldObject, err := c.gimulator.Get(worldKey)
	if err != nil {
		return fmt.Errorf("can not get world: %v", err)
	}
	var world types.World
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
