package main

const (
	red Color = iota
	green
	blue
	gray
)

const (
	small Size = iota
	medium
	big
)

type Color int

type Size int

type Product struct {
	Name  string
	Color Color
	Size  Size
}
