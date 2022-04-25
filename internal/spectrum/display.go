package display

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-logr/logr"
	//"github.com/go-logr/logr"
)

var (
	log   logr.Logger
	strip ledstripControl
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

func Run(log logr.Logger) {
	log = log.WithName("display")

	options := ledstripOptions{
		brightness:     brightness,
		ledCount:       ledCounts,
		gpioPin:        gpioPin,
		renderWaitTime: renderWaitTime,
		frequency:      frequency,
		dmaNum:         dmaNum,
		stripType:      stripType,
	}
	log.V(0).Info("Initiating led strip 💡", "options", fmt.Sprintf("%+v", options))
	strip = &ws2811Control{nil, options}
	strip.Init()

	args := os.Args[1:]

	log.V(0).Info("Lauching animation 🎈", "args", args)
	if len(args) > 0 {
		if args[0] == "clear" {
			log.V(0).Info("Clear 🧹")
			clear_strip()
		} else if args[0] == "rainbow" {
			log.V(0).Info("Rainbow 🌈")
			clear_strip()
			rainbown()
		} else if args[0] == "wipe" {
			log.V(0).Info("Wipe 🎢")
			wipe()
		}
	} else {
		clear_strip()
	}
}

func clear_strip() {
	for i := 0; i < strip.Count(0); i++ {
		strip.SetLed(0, i, 0x00000000)
	}
	checkError(strip.Render())
}

func wipe() {
	for {
		clear_strip()
		for i := 0; i < strip.Count(0); i++ {
			strip.SetLed(0, i, 0xff000000)
			checkError(strip.Render())
			strip.Sync()
		}
	}
}

func rainbown() {
	for {
		for i := 0; i < 256; i++ {
			for j := 0; j < strip.Count(0); j++ {
				strip.SetLed(0, j, wheel(((j*256/strip.Count(0))+i)%256))
			}
			//t1 := time.Now()
			checkError(strip.Render())
			//dev.Wait()
			//t2 := time.Now()
			//diff := t2.Sub(t1)
			//fmt.Println(diff)
		}
	}

}

func wheel(pos int) uint32 {
	var r, g, b int
	if pos < 85 {
		r = pos * 3
		g = 255 - pos*3
		b = 0
	} else if pos < 170 {
		pos -= 85
		r = 255 - pos*3
		g = 0
		b = pos * 3
	} else {
		pos -= 170
		r = 0
		g = pos * 3
		b = 255 - pos*3
	}

	value, err := strconv.ParseUint(fmt.Sprintf("%02x%02x%02x", r, g, b), 16, 32)
	checkError(err)
	return uint32(value)
}
