package led

type AnimFunc func(chan struct{}, StripSegment, map[string]string)

type AnimUnit struct {
	CancelToken chan struct{}     `json:"-"`
	Segment     StripSegment      `json:"segment"`
	Anim        AnimFunc          `json:"-"`
	Animation   string            `json:"animation"`
	Options     map[string]string `json:"options"`
	IsRunning   bool              `json:"isRunning"`
}

func (s *AnimUnit) StartAnimation() {
	s.IsRunning = true
	go s.Anim(s.CancelToken, s.Segment, s.Options)
}

func (s *AnimUnit) StopAnimation() {
	if s.IsRunning && s.CancelToken != nil {
		close(s.CancelToken)
	}
	s.IsRunning = false
}
