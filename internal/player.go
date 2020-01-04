package spacegame

import (
	"github.com/faiface/pixel"
)

type Player struct {
	direction Direction
	world     *World
	sprite    *pixel.Sprite
	life      int
	pos       *pixel.Vec
}

const playerVel = 50

func NewPlayer(path string, life int, world *World) (*Player, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}
	spr := pixel.NewSprite(pic, pic.Bounds())
	initialPos := pixel.V(world.Bounds().W()/2, spr.Frame().H())

	return &Player{life: life, sprite: spr, world: world, pos: &initialPos}, nil
}

func (p Player) Frame() pixel.Rect {
	return p.sprite.Frame()
}

func (p Player) Draw(t pixel.Target) {
	p.sprite.Draw(t, pixel.IM.Moved(*p.pos))
}

func (p Player) Update(direction Direction) {
	p.direction = direction
	p.move(direction)
}

func (p *Player) move(direction Direction) {
	switch direction {
	case Left:

		newX := p.pos.X - playerVel
		if newX > 0 {
			p.pos.X = newX
		}

	case Right:
		newX := p.pos.X + playerVel
		if newX < p.world.Bounds().W() {
			p.pos.X = newX
		}
	}
}
