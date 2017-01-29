package main

import (
	"log"

	"github.com/koron/go-waveout"
)

const (
	rate  = 8000
	freq1 = 400
	freq2 = 440
	freq3 = 360
)

func wave(size, freq int) []byte {
	d := make([]byte, size)
	l := size / freq
	for i := range d {
		if i%l < l/2 {
			d[i] = 128 + 64
		} else {
			d[i] = 128 - 64
		}
	}
	return d
}

func main() {
	p, err := waveout.New(1, rate, 8)
	if err != nil {
		log.Printf("New() failed; %s", err)
		return
	}
	defer func() {
		err := p.Close()
		if err != nil {
			log.Printf("Close() failed: %s", err)
		}
	}()
	err = p.AppendChunks(2, rate)
	if err != nil {
		log.Printf("AppendChunks() failed: %s", err)
		return
	}
	_, err = p.Write(wave(rate, freq1))
	if err != nil {
		log.Printf("Write(#1) failed: %s", err)
	}
	_, err = p.Write(wave(rate, freq2))
	if err != nil {
		log.Printf("Write(#2) failed: %s", err)
	}
	p.Wait()
	_, err = p.Write(wave(2*rate, freq3))
	if err != nil {
		log.Printf("Write(#3) failed: %s", err)
	}
	p.Wait()
}
