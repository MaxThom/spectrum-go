package led

import (
	"fmt"
	"strconv"
	"time"

	"github.com/maxthom/spectrum-go/pkg/utils"
)

type Animation_1d struct {
	Strip LedstripControl
}

func (s *Animation_1d) Clear_strip(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	for i := segment.Start; i < segment.End; i++ {
		s.Strip.SetLed(0, i, 0x00000000)
	}
}

func (s *Animation_1d) Wipe(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	wait := utils.If_key_exist_else[time.Duration](options, "wait", 1*time.Millisecond)
	for {
		s.Clear_strip(cancelToken, segment, options)
		for i := segment.Start; i < segment.End; i++ {
			s.Strip.SetLed(0, i, 0xff000000)
			time.Sleep(5*time.Millisecond + wait)
			select {
			case <-cancelToken:
				return
			default:
			}
		}
	}
}

func (s *Animation_1d) Rainbown(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	wait := utils.If_key_exist_else[time.Duration](options, "wait", 1*time.Millisecond)
	for {
		for i := 0; i < 256; i++ {
			for j := segment.Start; j < segment.End; j++ {
				s.Strip.SetLed(0, j, wheel(((j*256/segment.End)+i)%256))
			}
			time.Sleep(1*time.Millisecond + wait)
			select {
			case <-cancelToken:
				return
			default:
			}
		}
	}
}

func wheel(pos int) uint32 {
	var r, g, b int
	if pos < 85 {
		r = pos * 3
		g = 255 - pos*3
		b = 0
	} else if pos < 170 {
		pos -= 85
		r = 255 - pos*3
		g = 0
		b = pos * 3
	} else {
		pos -= 170
		r = 0
		g = pos * 3
		b = 255 - pos*3
	}

	value, err := strconv.ParseUint(fmt.Sprintf("%02x%02x%02x", r, g, b), 16, 32)
	utils.CheckError(err)
	return uint32(value)
}
