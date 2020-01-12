package main

import (
	"fmt"
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	spacegame "github.com/friendsofgo/spacegame/internal"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Friends of Go: Space Game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}
	world := spacegame.NewWorld(windowWidth, windowHeight)
	if err := world.AddBackground("resources/background.png"); err != nil {
		log.Fatal(err)
	}

	player, err := spacegame.NewPlayer("resources/player.png", 5, world)
	if err != nil {
		log.Fatal(err)
	}

	direction := spacegame.Idle
	action := spacegame.NoneAction
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Black)
		world.Draw(win)

		if win.Pressed(pixelgl.KeyLeft) {
			direction = spacegame.LeftDirection
		}

		if win.Pressed(pixelgl.KeyRight) {
			direction = spacegame.RightDirection
		}

		if win.Pressed(pixelgl.KeySpace) {
			action = spacegame.ShootAction
		}

		player.Update(direction, action, dt)
		player.Draw(win)
		direction = spacegame.Idle
		action = spacegame.NoneAction

		fps := 1 / dt
		fmt.Println("FPS: ", int(fps))

		win.Update()
	}
}
