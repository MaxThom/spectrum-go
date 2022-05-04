package led

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Animation_1d struct {
	Strip LedstripControl
}

func (s *Animation_1d) Clear_strip(segment StripSegment) {
	for i := segment.Start; i < segment.End; i++ {
		s.Strip.SetLed(0, i, 0x00000000)
	}
	//checkError(strip.Render())
}

func (s *Animation_1d) Wipe(ctx context.Context, segment StripSegment, wait time.Duration) {
	for {
		s.Clear_strip(segment)
		for i := segment.Start; i < segment.End; i++ {
			s.Strip.SetLed(0, i, 0xff000000)
			//checkError(strip.Render())
			//strip.Sync()
			time.Sleep(5*time.Millisecond + wait)
		}
		select {
		case <-ctx.Done():
			fmt.Println("DONE ANIM")
			return
		default:
		}
	}
}

func (s *Animation_1d) Rainbown(segment StripSegment, wait time.Duration) {
	for {
		for i := 0; i < 256; i++ {
			for j := segment.Start; j < segment.End; j++ {
				s.Strip.SetLed(0, j, wheel(((j*256/segment.End)+i)%256))
			}
			//t1 := time.Now()
			//checkError(strip.Render())
			time.Sleep(1*time.Millisecond + wait)
			//t2 := time.Now()
			//diff := t2.Sub(t1)
			//fmt.Println(diff)
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
	checkError(err)
	return uint32(value)
}
