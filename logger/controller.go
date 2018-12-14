package main

import (
	"fmt"
	"io"

	"github.com/alidadar7676/gimulator/simulator"
)

type Controller struct {
	Writer    io.Writer
	Wait      chan struct{}
	gimulator simulator.Gimulator
	watcher   chan simulator.Reconcile
}

func NewController(writer io.Writer, gimulator simulator.Gimulator) *Controller {
	return &Controller{
		Writer:    writer,
		gimulator: gimulator,
		watcher:   make(chan simulator.Reconcile, 1024),
		Wait:      make(chan struct{}),
	}
}

func (c *Controller) Run() {
	c.gimulator.Watch(simulator.Object{}, c.watcher)

	go func() {
		defer close(c.Wait)
		for r := range c.watcher {
			fmt.Fprintln(c.Writer, r)
		}
	}()
}
