package main

import "fmt"

type position struct {
	x int
	y int
}

func (p *position) move(x, y int) {
	p.x += x
	p.y += y
}

func (p *position) teleport(x, y int) {
	p.x = x
	p.y = y
}

type player struct {
	position
}

func newPlayer() *player {
	return &player{}
}

func (p *player) String() string {
	return fmt.Sprintf("(x, y) = (%d, %d)", p.x, p.y)
}

func main() {
	p := newPlayer()
	p.teleport(100, 200)
	p.move(10, 20)
	fmt.Println(p)
}
