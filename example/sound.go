// This example downloads (if necessary) a sound file to the current directory,
// loads it, and plays it. Once it's done playing the program quits.
package main

import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/acodec"
	"github.com/dradtke/go-allegro/allegro/audio"
	"io"
	"net/http"
	"os"
	"path"
)

const SOUND_FILE = "http://www.kozco.com/tech/piano2.wav"

func download(url string) (string, error) {
	filename := path.Base(url)
	fmt.Printf("downloading %s...", SOUND_FILE)

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Println("not necessary")
		return filename, nil
	}

	defer fmt.Println("done")

	output, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func main() {
	allegro.Run(func() {
		var (
			filename string
			sample   *audio.Sample
			instance *audio.SampleInstance
			err      error
		)

		if err = audio.Install(); err != nil {
			panic(err)
		}

		if err = acodec.Install(); err != nil {
			panic(err)
		}

		audio.ReserveSamples(1)
		defer audio.Uninstall()

		if filename, err = download(SOUND_FILE); err != nil {
			panic(err)
		}

		fmt.Printf("loading %s...", filename)
		sample, err = audio.LoadSample(filename)
		fmt.Println("done")
		if err != nil {
			panic(err)
		}

		instance = audio.CreateSampleInstance(sample)
		err = instance.AttachToMixer(audio.DefaultMixer())
		if err != nil {
			panic(err)
		}

		fmt.Print("playing...")
		err = instance.Play()
		if err != nil {
			panic(err)
		}

		for instance.Playing() {
			allegro.Rest(0.5)
		}

		fmt.Println("done")
	})
}
