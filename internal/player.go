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
	vel       float64
	laser     *Laser

	lasers map[string]*Laser
}

const (
	playerVel    = 250.0
	laserImg     = "resources/laser.png"
	laserSfx     = "resources/sfx/pew.wav"
	laserVel     = 270.0
	rechargeTime = 35
)

var (
	laserDelay = rechargeTime
)

func NewPlayer(path string, life int, world *World) (*Player, error) {

	// Initialize sprite to use with the player
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}
	spr := pixel.NewSprite(pic, pic.Bounds())
	initialPos := pixel.V(world.Bounds().W()/2, spr.Frame().H())

	// Initialize the laser for the player
	l, err := NewBaseLaser(laserImg, laserSfx, laserVel, world)
	if err != nil {
		return nil, err
	}

	return &Player{
		life:   life,
		sprite: spr,
		world:  world,
		pos:    &initialPos,
		vel:    playerVel,
		laser:  l,
		lasers: make(map[string]*Laser),
	}, nil
}

func (p Player) Frame() pixel.Rect {
	return p.sprite.Frame()
}

func (p Player) Draw(t pixel.Target) {
	p.sprite.Draw(t, pixel.IM.Moved(*p.pos))
	for _, l := range p.lasers {
		l.Draw(t)
	}
}

func (p *Player) Update(direction Direction, action Action, dt float64) {
	p.direction = direction
	p.move(direction, dt)
	p.shoot(action, dt)

	for k, l := range p.lasers {
		l.Update()

		// remove unused lasers
		if !l.isVisible {
			delete(p.lasers, k)
		}
	}
}

func (p *Player) move(direction Direction, dt float64) {
	switch direction {
	case LeftDirection:

		newX := p.pos.X - (p.vel * dt)
		if newX > 0 {
			p.pos.X = newX
		}

	case RightDirection:
		newX := p.pos.X + (p.vel * dt)
		if newX < p.world.Bounds().W() {
			p.pos.X = newX
		}
	}
}

func (p *Player) shoot(action Action, dt float64) {

	if laserDelay >= 0 {
		laserDelay--
	}

	if action == ShootAction && laserDelay <= 0 {
		l := p.laser.NewLaser(*p.pos)
		go l.Shoot()
		l.vel *= dt

		p.lasers[NewULID()] = l
		laserDelay = rechargeTime
	}
}
