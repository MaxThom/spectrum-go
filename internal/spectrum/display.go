package display

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-logr/logr"
	"github.com/maxthom/spectrum-go/pkg/led"
)

var (
	log    logr.Logger
	strip  led.LedstripControl
	anim1d led.Animation_1d
)

const (
	brightness     = 128
	ledCounts      = 144
	gpioPin        = 18
	renderWaitTime = 0
	frequency      = 1200000
	dmaNum         = 10
	stripType      = "SK6812StripGRBW"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run(plog logr.Logger) {
	log = plog.WithName("display")

	// Initialize LED Strip
	options := led.LedstripOptions{
		Brightness:     brightness,
		LedCount:       ledCounts,
		GpioPin:        gpioPin,
		RenderWaitTime: renderWaitTime,
		Frequency:      frequency,
		DmaNum:         dmaNum,
		StripType:      stripType,
	}
	log.V(0).Info("Initiating led strip 💡", "options", fmt.Sprintf("%+v", options))
	strip = &led.Ws2811Control{Strip: nil, Options: options}
	strip.Init()

	// Initialize animaters
	log.V(0).Info("Initializing animater 🕺", "dimension", "1d")
	anim1d = led.Animation_1d{Strip: strip}

	// Start rendering continusouly
	log.V(0).Info("Initializing renderer 🎢")
	go RenderContinuously()

	args := os.Args[1:]

	// Start animations...
	log.V(0).Info("Lauching animations 🎈", "args", args)
	if len(args) > 0 {
		if args[0] == "clear" {
			log.V(0).Info("Clear 🧹")
			go anim1d.Clear_strip(led.NewStripSegment(0, 144))
		} else if args[0] == "rainbow" {
			log.V(0).Info("Rainbow 🌈")
			anim1d.Clear_strip(led.NewStripSegment(0, 144))
			//done := make(chan struct{})
			//done2 := make(chan struct{})
			ctx, cancel := context.WithCancel(context.Background())
			go anim1d.Wipe(ctx, led.NewStripSegment(0, 36), 30*time.Millisecond)
			go anim1d.Rainbown(led.NewStripSegment(36, 72), 5*time.Millisecond)
			go anim1d.Wipe(ctx, led.NewStripSegment(72, 108), 30*time.Millisecond)
			go anim1d.Rainbown(led.NewStripSegment(108, 144), 5*time.Millisecond)

			time.Sleep(5 * time.Second)
			//done <- struct{}{}
			//done2 <- struct{}{}
			cancel()
			log.V(0).Info("Done Display")
			time.Sleep(3 * time.Second)
			go anim1d.Wipe(ctx, led.NewStripSegment(0, 36), 30*time.Millisecond)
			go anim1d.Wipe(ctx, led.NewStripSegment(72, 108), 30*time.Millisecond)

			time.Sleep(5 * time.Second)
			//done <- struct{}{}
			//done2 <- struct{}{}
			cancel()
			log.V(0).Info("Done Display")

		} else if args[0] == "wipe" {
			log.V(0).Info("Wipe 🎢")
			//anim1d.Wipe(led.NewStripSegment(0, 144), 30*time.Millisecond)
		}
	} else {
		go anim1d.Clear_strip(led.NewStripSegment(0, 144))
	}
}

func RenderContinuously() {
	for {
		//t1 := time.Now()
		checkError(strip.Render())
		checkError(strip.Sync())
		//t2 := time.Now()
		//diff := t2.Sub(t1)
		//log.V(0).Info("Render time", "ms", diff)
	}
}
