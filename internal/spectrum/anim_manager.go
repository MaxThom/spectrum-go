package spectrum

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/maxthom/spectrum-go/pkg/led"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

var ()

const (
	brightness     = 128
	ledCounts      = 144
	gpioPin        = 18
	renderWaitTime = 0
	frequency      = 1200000
	dmaNum         = 10
	stripType      = "SK6812StripGRBW"
)

type AnimManager struct {
	log        logr.Logger
	strip      led.LedstripControl
	anim1d     *led.Animation_1d
	animations []*led.AnimUnit
}

func NewAnimManager(plog logr.Logger, args []string) *AnimManager {
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
	plog.V(0).Info("Initiating led strip 💡", "options", fmt.Sprintf("%+v", options))
	strip := &led.Ws2811Control{Strip: nil, Options: options}
	strip.Init()

	// Initialize animaters
	plog.V(0).Info("Initializing animater 🕺", "dimension", "1d")
	anim1d := &led.Animation_1d{Strip: strip}

	// Start rendering continusouly
	plog.V(0).Info("Initializing renderer 🎢")
	go strip.RenderContinuously()

	return &AnimManager{
		log:        plog.WithName("manager"),
		strip:      strip,
		anim1d:     anim1d,
		animations: []*led.AnimUnit{},
	}
}

func (m *AnimManager) GetAnimations() []*led.AnimUnit {
	return m.animations
}

func (m *AnimManager) PlayDefaultAnimations() {
	m.animations = append(m.animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(0, 36),
		Anim:        m.anim1d.Maze,
		Animation:   utils.GetFunctionName(m.anim1d.Maze),
		Options: map[string]any{
			"wait":          50 * time.Millisecond,
			"count":         3,
			"turn_chance":   2,
			"color":         uint32(0x00ff88ff),
			"contact_color": uint32(0xff00ffff),
		},
		IsRunning: false,
	})
	m.animations = append(m.animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(36, 72),
		Anim:        m.anim1d.Rainbow,
		Animation:   utils.GetFunctionName(m.anim1d.Rainbow),
		Options: map[string]any{
			"wait": 5 * time.Millisecond,
		},
		IsRunning: false,
	})
	m.animations = append(m.animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(72, 108),
		Anim:        m.anim1d.Wipe,
		Animation:   utils.GetFunctionName(m.anim1d.Wipe),
		Options: map[string]any{
			"wait":  30 * time.Millisecond,
			"color": uint32(0x00ff0077),
		},
		IsRunning: false,
	})
	m.animations = append(m.animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(108, 144),
		Anim:        m.anim1d.Rainbow,
		Animation:   utils.GetFunctionName(m.anim1d.Rainbow),
		Options: map[string]any{
			"wait": 5 * time.Millisecond,
		},
		IsRunning: false,
	})

	// Start animations...
	for i, animUnit := range m.animations {
		//m.log.V(0).Info("Starting default animation", "index", i, "name", utils.GetFunctionName(animUnit.Anim), "segment", animUnit.Segment, "options", animUnit.Options)
		m.log.V(0).Info("Starting default animation", "index", i, "details", animUnit)
		animUnit.StartAnimation()
	}
	m.log.V(0).Info("All animation started 🙂.")
}

//if len(args) > 0 {
//	if args[0] == "clear" {
//		log.V(0).Info("Clear 🧹")
//		animations = append(animations, led.AnimUnit{
//			CancelToken: nil,
//			Segment:     led.NewStripSegment(0, 144),
//			Anim:        anim1d.Clear_strip,
//			Options:     map[string]any{},
//		})
//	} else if args[0] == "rainbow" {
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(0, 36),
//			Anim:        anim1d.Maze,
//			Options: map[string]any{
//				"wait":          50 * time.Millisecond,
//				"count":         3,
//				"turn_chance":   2,
//				"color":         uint32(0x00ff88ff),
//				"contact_color": uint32(0xff00ffff),
//			},
//		})
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(36, 72),
//			Anim:        anim1d.Rainbown,
//			Options: map[string]any{
//				"wait": 5 * time.Millisecond,
//			},
//		})
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(72, 108),
//			Anim:        anim1d.Wipe,
//			Options: map[string]any{
//				"wait":  30 * time.Millisecond,
//				"color": uint32(0x00ff0077),
//			},
//		})
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(108, 144),
//			Anim:        anim1d.Rainbown,
//			Options: map[string]any{
//				"wait": 5 * time.Millisecond,
//			},
//		})
//	} else if args[0] == "wipe" {
//		log.V(0).Info("Wipe 🎢")
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(0, 144),
//			Anim:        anim1d.Wipe,
//			Options: map[string]any{
//				"wait": 10 * time.Millisecond,
//			},
//		})
//	} else if args[0] == "maze" {
//		log.V(0).Info("Maze 🌌")
//		animations = append(animations, led.AnimUnit{
//			CancelToken: make(chan struct{}),
//			Segment:     led.NewStripSegment(0, 144),
//			Anim:        anim1d.Maze,
//			Options: map[string]any{
//				"wait":          30 * time.Millisecond,
//				"count":         10,
//				"turn_chance":   2,
//				"color":         uint32(0x000000ff),
//				"contact_color": uint32(0xff000000),
//			},
//		})
//	}
//} else {
//	log.V(0).Info("Clear 🧹")
//	animations = append(animations, led.AnimUnit{
//		CancelToken: nil,
//		Segment:     led.NewStripSegment(0, 144),
//		Anim:        anim1d.Clear_strip,
//		Options:     map[string]any{},
//	})
//}

//for i, animUnit := range animations {
//	log.V(0).Info("Starting animation", "index", i, "name", GetFunctionName(animUnit.Anim), "segment", animUnit.Segment, "options", animUnit.Options)
//	animUnit.StartAnimation()
//}
//log.V(0).Info("All animation started 🙂.")
//time.Sleep(5 * time.Second)
//for i, animUnit := range animations {
//	log.V(0).Info("Stopping animation", "index", i, "name", GetFunctionName(animUnit.Anim))
//	animUnit.StopAnimation()
//}
//log.V(0).Info("All animation stopped 😐.")
