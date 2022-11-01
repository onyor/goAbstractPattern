package main

type IProduct interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Product struct {
	logo string
	size int
}

func (s *Product) setLogo(logo string) {
	s.logo = logo
}

func (s *Product) getLogo() string {
	return s.logo
}

func (s *Product) setSize(size int) {
	s.size = size
}

func (s *Product) getSize() int {
	return s.size
}
