package led

import "github.com/maxthom/spectrum-go/pkg/utils"

type AnimFunc func(chan struct{}, StripSegment, map[string]string)

type AnimUnit struct {
	CancelToken chan struct{}     `json:"-"`
	Segment     StripSegment      `json:"segment"`
	Anim        AnimFunc          `json:"-"`
	Animation   string            `json:"animation"`
	Options     map[string]string `json:"options"`
	IsRunning   bool              `json:"isRunning"`
}

func (s *AnimUnit) StartAnimation(anim1d *Animation_1d) {
	s.IsRunning = true
	utils.InvokeAsync(anim1d, "Wipe", s.CancelToken, s.Segment, s.Options)
	//go s.Anim(s.CancelToken, s.Segment, s.Options)
}

func (s *AnimUnit) StopAnimation() {
	if s.IsRunning && s.CancelToken != nil {
		close(s.CancelToken)
	}
	s.IsRunning = false
}
