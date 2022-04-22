package main

import (
	"fmt"
	"os"
	"strconv"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness     = 128
	ledCounts      = 144
	gpioPin        = 18
	renderWaitTime = 0
	frequency      = 1200000
	dmaNum         = 10
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("----------------")
	fmt.Println(" Hardware Check ")
	fmt.Println("----------------")
	hw := ws2811.HwDetect()
	fmt.Printf("  Hardware Type    : %d\n", hw.Type)
	fmt.Printf("  Hardware Version : 0x%08X\n", hw.Version)
	fmt.Printf("  Periph base      : 0x%08X\n", hw.PeriphBase)
	fmt.Printf("  Video core base  : 0x%08X\n", hw.VideocoreBase)
	fmt.Printf("  Description      : %v\n", hw.Desc)

	fmt.Println("-----------")
	fmt.Println(" Led Strip ")
	fmt.Println("-----------")

	opt := ws2811.DefaultOptions
	opt.RenderWaitTime = renderWaitTime
	opt.Frequency = frequency
	opt.DmaNum = dmaNum
	opt.Channels[0].GpioPin = gpioPin
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCounts
	opt.Channels[0].StripeType = ws2811.SK6812StripRGBW
	fmt.Printf("  RenderWaitTime : %d\n", opt.RenderWaitTime)
	fmt.Printf("  Frequency      : %v\n", opt.Frequency)
	fmt.Printf("  DmaNum         : %v\n", opt.DmaNum)
	fmt.Printf("  GpioPin        : %v\n", opt.Channels[0].GpioPin)
	fmt.Printf("  Brightness     : %v\n", opt.Channels[0].Brightness)
	fmt.Printf("  LedCount       : %v\n", opt.Channels[0].LedCount)
	fmt.Printf("  StripeType     : %v\n", opt.Channels[0].StripeType)

	fmt.Println("------")
	fmt.Println(" Args ")
	fmt.Println("------")

	args := os.Args[1:]
	fmt.Println(" ", args)

	fmt.Println("----------")
	fmt.Println(" Spectrum ")
	fmt.Println("----------")

	dev, err := ws2811.MakeWS2811(&opt)
	checkError(err)

	checkError(dev.Init())
	defer dev.Fini()

	if len(args) > 0 {
		if args[0] == "clear" {
			clear_strip(dev)
		} else if args[0] == "rainbow" {
			clear_strip(dev)
			rainbown(dev)
		}
	} else {
		clear_strip(dev)
	}
}

func clear_strip(dev *ws2811.WS2811) {
	fmt.Println(" Clear 🧹")
	for i := 0; i < len(dev.Leds(0)); i++ {
		dev.Leds(0)[i] = 0x000000
	}
	checkError(dev.Render())
}

func rainbown(dev *ws2811.WS2811) {
	fmt.Println(" Rainbow 🌈")
	for {
		for i := 0; i < 256; i++ {
			for j := 0; j < len(dev.Leds(0)); j++ {
				dev.Leds(0)[j] = wheel(((j * 256 / len(dev.Leds(0))) + i) % 256)
			}
			//t1 := time.Now()
			checkError(dev.Render())
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
