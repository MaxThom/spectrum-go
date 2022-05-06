package led

type AnimFunc func(chan struct{}, StripSegment, map[string]any)

type AnimUnit struct {
	CancelToken chan struct{}
	Segment     StripSegment
	Anim        AnimFunc
	Options     map[string]any
}

func (s *AnimUnit) StartAnimation() {
	go s.Anim(s.CancelToken, s.Segment, s.Options)
}

func (s *AnimUnit) StopAnimation() {
	if s.CancelToken != nil {
		close(s.CancelToken)
	}
}
