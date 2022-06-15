package led

import "github.com/maxthom/spectrum-go/pkg/utils"

type AnimUnit struct {
	cancelToken chan struct{}     `json:"-"`
	Segment     StripSegment      `json:"segment"`
	Animer      Animer            `json:"-"`
	Animation   string            `json:"animation"`
	Options     map[string]string `json:"options"`
	IsRunning   bool              `json:"isRunning"`
}

func (s *AnimUnit) StartAnimation() {
	if s.Animation != "" {
		s.cancelToken = make(chan struct{})
		s.IsRunning = true
		utils.InvokeAsync(s.Animer, s.Animation, s.cancelToken, s.Segment, s.Options)
	}
}

func (s *AnimUnit) StopAnimation() {
	if s.IsRunning && s.cancelToken != nil {
		close(s.cancelToken)
	}
	s.IsRunning = false
}
