package led

import (
	"fmt"
	"strconv"
	"time"

	"github.com/maxthom/spectrum-go/pkg/utils"
)

// https://www.springtree.net/audio-visual-blog/rgb-led-color-mixing/

type Animation_1d struct {
	Strip LedstripControl
}

func (s *Animation_1d) Clear_strip(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	for i := segment.Start; i < segment.End; i++ {
		s.Strip.SetLed(0, i, 0x00000000)
	}
}

func (s *Animation_1d) Wipe(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	wait := utils.If_key_exist_else(options, "wait", 1*time.Millisecond)
	color := utils.If_key_exist_else(options, "color", uint32(0xff000000))
	for {
		s.Clear_strip(cancelToken, segment, options)
		for i := segment.Start; i < segment.End; i++ {
			s.Strip.SetLed(0, i, color)
			time.Sleep(5*time.Millisecond + wait)
			if cancellationRequest(cancelToken) {
				return
			}
		}
	}
}

func (s *Animation_1d) Rainbown(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	wait := utils.If_key_exist_else(options, "wait", 1*time.Millisecond)
	for {
		for i := 0; i < 256; i++ {
			for j := segment.Start; j < segment.End; j++ {
				s.Strip.SetLed(0, j, wheel(((j*256/segment.End)+i)%256))
			}
			time.Sleep(1*time.Millisecond + wait)
			if cancellationRequest(cancelToken) {
				return
			}
		}
	}
}

func (s *Animation_1d) Maze(cancelToken chan struct{}, segment StripSegment, options map[string]any) {
	wait := utils.If_key_exist_else(options, "wait", 1*time.Millisecond)
	count := utils.If_key_exist_else(options, "count", 10)
	turn_chance := utils.If_key_exist_else(options, "turn_chance", 2)
	color := utils.If_key_exist_else(options, "color", uint32(0x000000ff))
	contact_color := utils.If_key_exist_else(options, "contact_color", uint32(0xff000000))

	// Generate initial positions for all dots
	dots := []*Position{}
	location := make([]int, segment.len)
	for i := 0; i < count; i++ {
		dots = append(dots, NewPosition(0, segment.len))
		location[dots[i].position] += 1
	}

	for {
		// Set color according to number of dots on an index
		for i, nb := range location {
			switch nb {
			case 0:
				s.Strip.SetLed(0, i+segment.Start, 0x00000000)
			case 1:
				// Check if dots on neigbours position, if so we also want contact
				if i > 0 && location[i-1] >= 1 || i < segment.len-1 && location[i+1] >= 1 {
					s.Strip.SetLed(0, i+segment.Start, contact_color)
				} else {
					s.Strip.SetLed(0, i+segment.Start, color)
				}
			default:
				s.Strip.SetLed(0, i+segment.Start, contact_color)
			}
		}

		// Calculate next dot positions
		for i, dot := range dots {
			location[dots[i].position] -= 1
			if utils.RandomInt(0, 100) <= turn_chance {
				dot.direction *= -1
			}

			// Bounds
			dot.position += dot.direction
			if dot.position < 0 {
				dot.position = segment.len - 1
			} else if dot.position >= segment.len {
				dot.position = 0
			}
			location[dots[i].position] += 1
		}

		time.Sleep(5*time.Millisecond + wait)
		if cancellationRequest(cancelToken) {
			return
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
