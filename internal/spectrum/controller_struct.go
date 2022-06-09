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

type DiscoveryDO struct {
	Animation string              `json:"animation"`
	Options   map[string]OptionDO `json:"options"`
}

type OptionDO struct {
	Type    string `json:"type"`
	Default string `json:"default"`
}
