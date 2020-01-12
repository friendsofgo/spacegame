package spacegame

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
)

type Laser struct {
	pic       pixel.Picture
	sfxPath   string
	pos       *pixel.Vec
	vel       float64
	sprite    *pixel.Sprite
	isVisible bool
	world     *World
}

func NewBaseLaser(path, sfxPath string, origPos *pixel.Vec, vel float64, world *World) (*Laser, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	return &Laser{
		pic:     pic,
		pos:     origPos,
		vel:     vel,
		world:   world,
		sfxPath: sfxPath,
	}, nil
}

func (l *Laser) NewLaser(pos pixel.Vec) *Laser {
	spr := pixel.NewSprite(l.pic, l.pic.Bounds())

	return &Laser{
		pos:       &pos,
		vel:       l.vel,
		sprite:    spr,
		isVisible: true,
		world:     l.world,
		sfxPath:   l.sfxPath,
	}
}

func (l Laser) Draw(t pixel.Target) {
	if l.isVisible == true {
		l.sprite.Draw(t, pixel.IM.Moved(*l.pos))
	}
}

func (l *Laser) Update() {
	l.pos.Y += l.vel
	if l.pos.Y > l.world.height {
		l.isVisible = false
	}
}

func (l Laser) Shoot() {
	sfx, err := loadSound(l.sfxPath)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(sfx.format.SampleRate, sfx.format.SampleRate.N(time.Second/10))
	defer sfx.streamer.Close()

	done := make(chan bool)
	speaker.Play(beep.Seq(sfx.streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
