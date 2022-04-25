package display

import (
	"fmt"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type ledstripControl interface {
	Init()
	SetLed(channel int, i int, c uint32)
	GetLed(i int, channel int) uint32
	Count(channel int) int
	Render() error
	Sync() error
	SetBrightness(channel int, brightness int)
	SetCustomGammaFactor(gammaFactor float64)
	SetLedSync(channel int, leds []uint32)
	Info()
	Dispose()
}

type ledstripOptions struct {
	brightness     int
	ledCount       int
	gpioPin        int
	renderWaitTime int
	frequency      int
	dmaNum         int
	stripType      string
}

type ws2811Control struct {
	strip   *ws2811.WS2811
	options ledstripOptions
}

func (s *ws2811Control) Init() {
	opt := ws2811.DefaultOptions
	opt.RenderWaitTime = s.options.renderWaitTime
	opt.Frequency = s.options.frequency
	opt.DmaNum = s.options.dmaNum
	opt.Channels[0].GpioPin = s.options.gpioPin
	opt.Channels[0].Brightness = s.options.brightness
	opt.Channels[0].LedCount = s.options.ledCount

	if s.options.stripType == "SK6812StripGRBW" {
		opt.Channels[0].StripeType = ws2811.SK6812StripGRBW
	}

	strip, err := ws2811.MakeWS2811(&opt)
	checkError(err)
	s.strip = strip

	checkError(s.strip.Init())
}

func (s *ws2811Control) Info() {
	hw := ws2811.HwDetect()
	fmt.Println("----------------")
	fmt.Println(" Hardware Check ")
	fmt.Println("----------------")
	fmt.Printf("  Hardware Type    : %d\n", hw.Type)
	fmt.Printf("  Hardware Version : 0x%08X\n", hw.Version)
	fmt.Printf("  Periph base      : 0x%08X\n", hw.PeriphBase)
	fmt.Printf("  Video core base  : 0x%08X\n", hw.VideocoreBase)
	fmt.Printf("  Description      : %v\n", hw.Desc)

	fmt.Println("-----------")
	fmt.Println(" Led Strip ")
	fmt.Println("-----------")
	fmt.Printf("  RenderWaitTime : %d\n", s.options.renderWaitTime)
	fmt.Printf("  Frequency      : %v\n", s.options.frequency)
	fmt.Printf("  DmaNum         : %v\n", s.options.dmaNum)
	fmt.Printf("  GpioPin        : %v\n", s.options.gpioPin)
	fmt.Printf("  Brightness     : %v\n", s.options.brightness)
	fmt.Printf("  LedCount       : %v\n", s.options.ledCount)
	fmt.Printf("  StripeType     : %v\n", s.options.stripType)
}

func (s *ws2811Control) SetLed(channel int, i int, c uint32) {
	s.strip.Leds(channel)[i] = c
}

func (s *ws2811Control) GetLed(i int, channel int) uint32 {
	return s.strip.Leds(channel)[i]
}

func (s *ws2811Control) Count(channel int) int {
	return len(s.strip.Leds(channel))
}

func (s *ws2811Control) Render() error {
	return s.strip.Render()
}

func (s *ws2811Control) Sync() error {
	return s.strip.Wait()
}

func (s *ws2811Control) SetBrightness(channel int, brightness int) {
	s.strip.SetBrightness(channel, brightness)
}

func (s *ws2811Control) SetCustomGammaFactor(gammaFactor float64) {
	s.strip.SetCustomGammaFactor(gammaFactor)
}

func (s *ws2811Control) SetLedSync(channel int, leds []uint32) {
	s.strip.SetLedsSync(channel, leds)
}

func (s *ws2811Control) Dispose() {
	s.strip.Fini()
}
