package spectrum

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/maxthom/spectrum-go/pkg/led"
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
		Segment:   led.NewStripSegment(0, 36),
		Animer:    anim1d,
		Animation: "Maze",
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
		Segment:   led.NewStripSegment(36, 72),
		Animer:    anim1d,
		Animation: "Rainbow",
		Options: map[string]string{
			"wait": "5",
		},
		IsRunning: false,
	})
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(72, 108),
		Animer:    anim1d,
		Animation: "Wipe",
		Options: map[string]string{
			"wait":  "30",
			"color": "0x00ff0077",
		},
		IsRunning: false,
	})
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(108, 144),
		Animer:    anim1d,
		Animation: "Rainbow",
		Options: map[string]string{
			"wait": "5",
		},
		IsRunning: false,
	})

	// Start animations...
	for i, animUnit := range animations {
		//m.log.V(0).Info("Starting default animation", "index", i, "name", utils.GetFunctionName(animUnit.Anim), "segment", animUnit.Segment, "options", animUnit.Options)
		log.V(0).Info("Starting default animation", "index", i, "details", animUnit)
		animUnit.StartAnimation()
	}
	log.V(0).Info("All animation started 🙂.")
}

func SetAnimation(anim AnimUnitDO) {
	if anim.Index == -1 {
		stopAllAnimations(true)
		anim.Index = 0
	}

	play := &led.AnimUnit{
		Segment:   led.NewStripSegment(anim.Segment.Start, anim.Segment.End),
		Animer:    anim1d,
		Animation: anim.Animation,
		Options:   anim.Options,
		IsRunning: false,
	}

	if len(animations) > anim.Index {
		animations[anim.Index] = play
		animations[anim.Index].StopAnimation()
	} else {
		animations = append(animations, play)
		anim.Index = len(animations) - 1
	}

	log.V(0).Info("Starting animation", "index", anim.Index, "details", animations[anim.Index])
	animations[anim.Index].StartAnimation()
}

func StopAnimation(anim AnimStopDO) {
	if anim.Index == -1 {
		stopAllAnimations(anim.ShouldClear)
	} else if len(animations) > anim.Index {
		log.V(0).Info("Stopping animation", "index", anim.Index, "details", animations[anim.Index])
		animations[anim.Index].StopAnimation()
		if anim.ShouldClear {
			animations[anim.Index].Animation = "Clear"
			animations[anim.Index].StartAnimation()
		}
	}

}

func stopAllAnimations(shouldClear bool) {
	for _, animUnit := range animations {
		animUnit.StopAnimation()
		if shouldClear {
			animUnit.Animation = "Clear"
			animUnit.StartAnimation()
		}
	}
	log.V(0).Info("All animation stopped 😐.")
	animations = []*led.AnimUnit{}
}
