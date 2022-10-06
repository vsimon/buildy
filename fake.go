package main

import (
	log "github.com/sirupsen/logrus"
)

type FakeLight struct {
}

func NewFakeLight() Light {
	return &FakeLight{}
}

func (l *FakeLight) Open() error {
	log.Info("Open light")
	return nil
}

func (l *FakeLight) Close() error {
	log.Info("Close light")
	return nil
}

func (l *FakeLight) Toggle(color Color) error {
	switch color {
	case Red:
		log.Info("Toggle red light")
	case Green:
		log.Info("Toggle green light")
	case Yellow:
		log.Info("Toggle yellow light")
	case All:
		log.Info("Toggle all lights")
	case None:
		log.Info("Toggle no lights")
	}

	return nil
}
