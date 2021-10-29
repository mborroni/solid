package main

type Rectangle struct {
	Width int
	Height int
}

func (r *Rectangle) setWidth(width int) {
	r.Width = width
}

func (r *Rectangle) setHeight(height int) {
	r.Height = height
}

func (r *Rectangle) calculateArea() int {
	return r.Width * r.Height
}