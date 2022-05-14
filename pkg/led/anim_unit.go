package led

import "github.com/maxthom/spectrum-go/pkg/utils"

type AnimUnit struct {
	CancelToken chan struct{}     `json:"-"`
	Segment     StripSegment      `json:"segment"`
	Animer      Animer            `json:"-"`
	Animation   string            `json:"animation"`
	Options     map[string]string `json:"options"`
	IsRunning   bool              `json:"isRunning"`
}

func (s *AnimUnit) StartAnimation() {
	s.IsRunning = true
	utils.InvokeAsync(s.Animer, s.Animation, s.CancelToken, s.Segment, s.Options)
}

func (s *AnimUnit) StopAnimation() {
	if s.IsRunning && s.CancelToken != nil {
		close(s.CancelToken)
	}
	s.IsRunning = false
}
