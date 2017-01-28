package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/koron/go-waveout"
)

const (
	rate = 8000
	freq = 400
)

func stage2(h syscall.Handle, hdr *waveout.WaveHdr) {
	r := waveout.Write(h, hdr, uint32(unsafe.Sizeof(*hdr)))
	if r != 0 {
		log.Printf("Write() failed: %s", r.Error())
		return
	}
	log.Printf("Write() done")
	time.Sleep(2 * time.Second)
	r = waveout.Reset(h)
	if r != 0 {
		log.Printf("Reset() failed: %s", r.Error())
		return
	}
}

func stage1(h syscall.Handle) {
	d := make([]byte, rate*2)
	l := rate / freq
	for i := range d {
		if i%l < l/2 {
			d[i] = 128 + 64
		} else {
			d[i] = 128 - 64
		}
	}
	hdr := waveout.WaveHdr{
		Data:         &d[0],
		BufferLength: uint32(len(d)),
		Flags:        uint32(waveout.WHDR_BEGINLOOP | waveout.WHDR_ENDLOOP),
		Loops:        1,
	}
	r := waveout.PrepareHeader(h, &hdr, uint32(unsafe.Sizeof(hdr)))
	if r != 0 {
		log.Printf("PrepareHeader() failed: %s", r.Error())
		return
	}
	defer func() {
		if r := waveout.UnprepareHeader(h, &hdr, uint32(unsafe.Sizeof(hdr))); r != 0 {
			log.Printf("UnprepareHeader() failed: %s", r.Error())
			return
		}
	}()
	stage2(h, &hdr)
}

func main() {
	p := waveout.WaveFormatEx{
		FormatTag:      waveout.WAVE_FORMAT_PCM,
		Channels:       1,
		SamplesPerSec:  uint32(rate),
		AvgBytesPerSec: uint32(rate * 1),
		BlockAlign:     1,
		BitsPerSample:  8,
	}
	var h syscall.Handle
	r := waveout.Open(&h, waveout.WAVE_MAPPER, &p, 0, 0, waveout.CALLBACK_NULL)
	if r != 0 {
		log.Printf("Open() failed: %s", r.Error())
		return
	}
	defer func() {
		r := waveout.Close(h)
		if r != 0 {
			fmt.Printf("Close() failed: %s", r.Error())
		}
	}()
	stage1(h)
}
