package spectrum

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/maxthom/spectrum-go/pkg/led"
)

var (
	log        logr.Logger
	strip      led.LedstripControl
	options    *led.LedstripOptions
	anim1d     *led.Animation_1d
	animations []*led.AnimUnit
)

const (
	pathToConfig = "configs/spectrum/ledstrip.json"
)

func Run(plog logr.Logger) {
	log = plog.WithName("controller")

	Init(led.NewFromFileLedstripOptions(pathToConfig))
	PlayDefaultAnimations()
}

func Init(p_options *led.LedstripOptions) {
	// Initialize LED Strip
	options = p_options

	log.V(0).Info("Initiating led strip 💡", "options", fmt.Sprintf("%+v", options))
	strip = &led.Ws2811Control{Strip: nil, Options: *options}
	strip.Init()

	// Initialize animaters
	log.V(0).Info("Initializing animater 🕺", "dimension", "1d")
	anim1d = &led.Animation_1d{Strip: strip}

	// Start rendering continusouly
	log.V(0).Info("Initializing renderer 🎢")
	go strip.RenderContinuously()

}

func PlayDefaultAnimations() {
	segmentCount := options.LedCount / 4
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(segmentCount*0, segmentCount*1),
		Animer:    anim1d,
		Animation: "Maze",
		Options: map[string]string{
			"wait":          "50",
			"count":         "3",
			"turn_chance":   "2",
			"color":         "0x00ff88ff",
			"contact_color": "0xff00ffff",
		},
	})
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(segmentCount*1, segmentCount*2),
		Animer:    anim1d,
		Animation: "Rainbow",
		Options: map[string]string{
			"wait": "5",
		},
	})
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(segmentCount*2, segmentCount*3),
		Animer:    anim1d,
		Animation: "Wipe",
		Options: map[string]string{
			"wait":  "30",
			"color": "0x00ff0077",
		},
	})
	animations = append(animations, &led.AnimUnit{
		Segment:   led.NewStripSegment(segmentCount*3, segmentCount*4),
		Animer:    anim1d,
		Animation: "Rainbow",
		Options: map[string]string{
			"wait": "5",
		},
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
		animations = []*led.AnimUnit{}
		anim.Index = 0
	}

	play := &led.AnimUnit{
		Segment:   led.NewStripSegment(anim.Segment.Start, anim.Segment.End),
		Animer:    anim1d,
		Animation: anim.Animation,
		Options:   anim.Options,
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
		animations = []*led.AnimUnit{}
	} else if len(animations) > anim.Index {
		log.V(0).Info("Stopping animation", "index", anim.Index, "details", animations[anim.Index])
		animations[anim.Index].StopAnimation()
		if anim.ShouldClear {
			animations[anim.Index].Animation = "Clear"
			animations[anim.Index].StartAnimation()
		}
		animations = append(animations[:anim.Index], animations[anim.Index+1:]...)
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
}

func startAllAnimations() {
	for _, animUnit := range animations {
		animUnit.StartAnimation()
	}
	log.V(0).Info("All animation started 🙂.")
}
