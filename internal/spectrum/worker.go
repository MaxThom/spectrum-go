package spectrum

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/maxthom/spectrum-go/pkg/led"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

var (
	log        logr.Logger
	strip      led.LedstripControl
	anim1d     *led.Animation_1d
	animations []*led.AnimUnit
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

func Init(plog logr.Logger, args []string) {
	log = plog.WithName("controller")

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
	anim1d = &led.Animation_1d{Strip: strip}

	// Start rendering continusouly
	log.V(0).Info("Initializing renderer 🎢")
	go strip.RenderContinuously()

}

func PlayDefaultAnimations() {
	animations = append(animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(0, 36),
		Anim:        anim1d.Maze,
		Animation:   utils.GetFunctionName(anim1d.Maze),
		Options: map[string]string{
			"wait":          "50",
			"count":         "3",
			"turn_chance":   "2",
			"color":         "0x00ff88ff",
			"contact_color": "0xff00ffff",
		},
		IsRunning: false,
	})
	animations = append(animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(36, 72),
		Anim:        anim1d.Rainbow,
		Animation:   utils.GetFunctionName(anim1d.Rainbow),
		Options: map[string]string{
			"wait": "5",
		},
		IsRunning: false,
	})
	animations = append(animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(72, 108),
		Anim:        anim1d.Wipe,
		Animation:   utils.GetFunctionName(anim1d.Wipe),
		Options: map[string]string{
			"wait":  "30",
			"color": "0x00ff0077",
		},
		IsRunning: false,
	})
	animations = append(animations, &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(108, 144),
		Anim:        anim1d.Rainbow,
		Animation:   utils.GetFunctionName(anim1d.Rainbow),
		Options: map[string]string{
			"wait": "5",
		},
		IsRunning: false,
	})

	// Start animations...
	for i, animUnit := range animations {
		//m.log.V(0).Info("Starting default animation", "index", i, "name", utils.GetFunctionName(animUnit.Anim), "segment", animUnit.Segment, "options", animUnit.Options)
		log.V(0).Info("Starting default animation", "index", i, "details", animUnit)
		animUnit.StartAnimation(anim1d)
	}
	log.V(0).Info("All animation started 🙂.")
}

func SetAnimation(anim AnimUnitDO) {
	if anim.Index == -1 {
		clearAllAnimations()
		anim.Index = 0
	}

	play := &led.AnimUnit{
		CancelToken: make(chan struct{}),
		Segment:     led.NewStripSegment(anim.Segment.Start, anim.Segment.End),
		Anim:        anim1d.Maze,
		Animation:   utils.GetFunctionName(anim1d.Maze),
		Options:     anim.Options,
		IsRunning:   false,
	}

	if len(animations) > anim.Index {
		animations[anim.Index] = play
		animations[anim.Index].StopAnimation()
	} else {
		animations = append(animations, play)
		anim.Index = len(animations) - 1
	}

	log.V(0).Info("Starting animation", "index", anim.Index, "details", animations[anim.Index])
	animations[anim.Index].StartAnimation(anim1d)
}

func clearAllAnimations() {
	for _, animUnit := range animations {
		animUnit.StopAnimation()
	}
	log.V(0).Info("All animation stopped 😐.")
	animations = []*led.AnimUnit{}
}
