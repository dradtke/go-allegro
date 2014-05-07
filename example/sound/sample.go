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

func download(url, filename string) error {
    output, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    _, err = io.Copy(output, response.Body)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    if err := allegro.Install(); err != nil {
        panic(err)
    }
    defer allegro.Uninstall()

    if err := acodec.Install(); err != nil {
        panic(err)
    }

    if err := audio.Install(); err != nil {
        panic(err)
    }
    audio.ReserveSamples(1)
    defer audio.Uninstall()

    filename := path.Base(SOUND_FILE)
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Printf("downloading %s...", SOUND_FILE)
        err = download(SOUND_FILE, filename)
        fmt.Println("done")
        if err != nil {
            panic(err)
        }
    }

    fmt.Printf("loading %s...", filename)
    sample, err := audio.LoadSample(filename)
    fmt.Println("done")
    if err != nil {
        panic(err)
    }

    instance := audio.CreateSampleInstance(sample)
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
}
