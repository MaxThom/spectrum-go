package led

import (
	"fmt"

	"github.com/maxthom/spectrum-go/pkg/utils"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type LedstripControl interface {
	Init()
	SetLed(channel int, i int, c uint32)
	GetLed(i int, channel int) uint32
	Count(channel int) int
	Render() error
	RenderContinuously()
	Sync() error
	SetBrightness(channel int, brightness int)
	SetCustomGammaFactor(gammaFactor float64)
	SetLedSync(channel int, leds []uint32)
	Info()
	Dispose()
}

type StripSegment struct {
	Start int `json:"start"` // Include
	End   int `json:"end"`   // Exclude
	len   int // End - Start
}

func NewStripSegment(start int, end int) StripSegment {
	return StripSegment{
		Start: start,
		End:   end,
		len:   end - start,
	}
}

type Ws2811Control struct {
	Strip   *ws2811.WS2811
	Options LedstripOptions
}

func (s *Ws2811Control) Init() {
	opt := ws2811.DefaultOptions
	opt.RenderWaitTime = s.Options.RenderWaitTime
	opt.Frequency = s.Options.Frequency
	opt.DmaNum = s.Options.DmaNum
	opt.Channels[0].GpioPin = s.Options.GpioPin
	opt.Channels[0].Brightness = s.Options.Brightness
	opt.Channels[0].LedCount = s.Options.LedCount

	if s.Options.StripType == "SK6812StripGRBW" {
		opt.Channels[0].StripeType = ws2811.SK6812StripGRBW
	}

	strip, err := ws2811.MakeWS2811(&opt)
	utils.CheckError(err)
	s.Strip = strip

	utils.CheckError(s.Strip.Init())
}

func (s *Ws2811Control) Info() {
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
	fmt.Printf("  RenderWaitTime : %d\n", s.Options.RenderWaitTime)
	fmt.Printf("  Frequency      : %v\n", s.Options.Frequency)
	fmt.Printf("  DmaNum         : %v\n", s.Options.DmaNum)
	fmt.Printf("  GpioPin        : %v\n", s.Options.GpioPin)
	fmt.Printf("  Brightness     : %v\n", s.Options.Brightness)
	fmt.Printf("  LedCount       : %v\n", s.Options.LedCount)
	fmt.Printf("  StripeType     : %v\n", s.Options.StripType)
}

func (s *Ws2811Control) SetLed(channel int, i int, c uint32) {
	s.Strip.Leds(channel)[i] = c
}

func (s *Ws2811Control) GetLed(i int, channel int) uint32 {
	return s.Strip.Leds(channel)[i]
}

func (s *Ws2811Control) Count(channel int) int {
	return len(s.Strip.Leds(channel))
}

func (s *Ws2811Control) Render() error {
	return s.Strip.Render()
}

func (s *Ws2811Control) Sync() error {
	return s.Strip.Wait()
}

func (s *Ws2811Control) SetBrightness(channel int, brightness int) {
	s.Strip.SetBrightness(channel, brightness)
}

func (s *Ws2811Control) SetCustomGammaFactor(gammaFactor float64) {
	s.Strip.SetCustomGammaFactor(gammaFactor)
}

func (s *Ws2811Control) SetLedSync(channel int, leds []uint32) {
	s.Strip.SetLedsSync(channel, leds)
}

func (s *Ws2811Control) Dispose() {
	s.Strip.Fini()
}

func (s *Ws2811Control) RenderContinuously() {
	for {
		//t1 := time.Now()
		utils.CheckError(s.Render())
		utils.CheckError(s.Sync())
		//t2 := time.Now()
		//diff := t2.Sub(t1)
		//log.V(0).Info("Render time", "ms", diff)
	}
}
