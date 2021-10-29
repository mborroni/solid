package main

type Square struct {
	Rectangle
}

func (s *Square) setWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) setHeight(height int) {
	s.Height = height
	s.Width = height
}
