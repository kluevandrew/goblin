package main

import (
	"bytes"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"goblin/data"
	"io"
	"os"
	"time"
)

func main()  {
	done := make(chan bool)
	reader := io.NopCloser(bytes.NewReader(data.Track))

	streamer, format, err := mp3.Decode(reader)
	if err != nil {
		os.Exit(1)
	}
	defer func(streamer beep.StreamSeekCloser) {
		err := streamer.Close()
		if err != nil {
			return
		}
	}(streamer)

	err = 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		os.Exit(2)
	}
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
