package spacegame

import (
	"log"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

type SFX struct {
	streamer beep.StreamSeekCloser
	format   beep.Format
}

func loadSound(path string) (*SFX, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return &SFX{
		streamer: streamer,
		format:   format,
	}, nil
}
