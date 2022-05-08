package led

type AnimFunc func(chan struct{}, StripSegment, map[string]any)

type AnimUnit struct {
	CancelToken chan struct{}  `json:"-"`
	Segment     StripSegment   `json:"segment"`
	Anim        AnimFunc       `json:"-"`
	Animation   string         `json:"animation"`
	Options     map[string]any `json:"options"`
	IsRunning   bool           `json:"isRunning"`
}

func (s *AnimUnit) StartAnimation() {
	s.IsRunning = true
	go s.Anim(s.CancelToken, s.Segment, s.Options)
}

func (s *AnimUnit) StopAnimation() {
	s.IsRunning = false
	if s.CancelToken != nil {
		close(s.CancelToken)
	}
}
