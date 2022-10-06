package main

type Color int

const (
	Red Color = iota
	Green
	Yellow
	All
	None
)

type Light interface {
	Open() error
	Close() error
	Toggle(color Color) error
}
