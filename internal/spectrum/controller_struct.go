package spectrum

import (
	"github.com/maxthom/spectrum-go/pkg/led"
)

type AnimUnitDO struct {
	Index     int               `json:"index"`
	Segment   led.StripSegment  `json:"segment"`
	Animation string            `json:"animation"`
	Options   map[string]string `json:"options"`
}

type AnimStopDO struct {
	Index       int  `json:"index"`
	ShouldClear bool `json:"shouldClear"`
}

type BrightnessDO struct {
	Brightness int `json:"brightness"`
}
