package display

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
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

func Run(plog logr.Logger, args []string) {
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

	// Start animations...
	animations := []led.AnimUnit{}
	if len(args) > 0 {
		if args[0] == "clear" {
			log.V(0).Info("Clear 🧹")
			animations = append(animations, led.AnimUnit{
				CancelToken: nil,
				Segment:     led.NewStripSegment(0, 144),
				Anim:        anim1d.Clear_strip,
				Options:     map[string]any{},
			})
		} else if args[0] == "rainbow" {
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(0, 36),
				Anim:        anim1d.Maze,
				Options: map[string]any{
					"wait":          50 * time.Millisecond,
					"count":         3,
					"turn_chance":   2,
					"color":         uint32(0x00ff88ff),
					"contact_color": uint32(0xff00ffff),
				},
			})
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(36, 72),
				Anim:        anim1d.Rainbown,
				Options: map[string]any{
					"wait": 5 * time.Millisecond,
				},
			})
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(72, 108),
				Anim:        anim1d.Wipe,
				Options: map[string]any{
					"wait":  30 * time.Millisecond,
					"color": uint32(0x00ff0077),
				},
			})
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(108, 144),
				Anim:        anim1d.Rainbown,
				Options: map[string]any{
					"wait": 5 * time.Millisecond,
				},
			})
		} else if args[0] == "wipe" {
			log.V(0).Info("Wipe 🎢")
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(0, 144),
				Anim:        anim1d.Wipe,
				Options: map[string]any{
					"wait": 10 * time.Millisecond,
				},
			})
		} else if args[0] == "maze" {
			log.V(0).Info("Maze 🌌")
			animations = append(animations, led.AnimUnit{
				CancelToken: make(chan struct{}),
				Segment:     led.NewStripSegment(0, 144),
				Anim:        anim1d.Maze,
				Options: map[string]any{
					"wait":          30 * time.Millisecond,
					"count":         10,
					"turn_chance":   2,
					"color":         uint32(0x000000ff),
					"contact_color": uint32(0xff000000),
				},
			})
		}
	} else {
		log.V(0).Info("Clear 🧹")
		animations = append(animations, led.AnimUnit{
			CancelToken: nil,
			Segment:     led.NewStripSegment(0, 144),
			Anim:        anim1d.Clear_strip,
			Options:     map[string]any{},
		})
	}

	for i, animUnit := range animations {
		log.V(0).Info("Starting animation", "index", i, "name", GetFunctionName(animUnit.Anim), "segment", animUnit.Segment, "options", animUnit.Options)
		animUnit.StartAnimation()
	}
	log.V(0).Info("All animation started 🙂.")
	//time.Sleep(5 * time.Second)
	//for i, animUnit := range animations {
	//	log.V(0).Info("Stopping animation", "index", i, "name", GetFunctionName(animUnit.Anim))
	//	animUnit.StopAnimation()
	//}
	//log.V(0).Info("All animation stopped 😐.")
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

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
