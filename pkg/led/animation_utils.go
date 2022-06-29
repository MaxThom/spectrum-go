package led

import (
	"strconv"
	"strings"
	"time"

	"github.com/maxthom/spectrum-go/pkg/utils"
)

type Animer interface {
}

type Position struct {
	position  int
	direction int
}

func NewPosition(start int, end int) *Position {
	dir := utils.RandomInt(0, 2)
	if dir == 0 {
		dir = -1
	}
	return &Position{
		position:  utils.RandomInt(start, end),
		direction: dir,
	}
}

func cancellationRequest(cancelToken chan struct{}) bool {
	select {
	case <-cancelToken:
		return true
	default:
		return false
	}
}

func getColorOption(dict map[string]string, key string, def uint32) uint32 {
	if strVal, ok := dict[key]; ok {
		cleaned := strings.Replace(strVal, "0x", "", -1)
		result, _ := strconv.ParseUint(cleaned, 16, 32)
		return uint32(result)
	}
	return def
}

func getTimeMsOption(dict map[string]string, key string, def time.Duration) time.Duration {
	if strVal, ok := dict[key]; ok {
		val, err := strconv.ParseInt(strVal, 0, 32)
		utils.CheckError(err)
		return time.Duration(val) * time.Millisecond
	}
	return def
}

func getIntOption(dict map[string]string, key string, def int) int {
	if strVal, ok := dict[key]; ok {
		val, err := strconv.ParseInt(strVal, 0, 32)
		utils.CheckError(err)
		return int(val)
	}
	return def
}

func getFloatOption(dict map[string]string, key string, def float32) float32 {
	if strVal, ok := dict[key]; ok {
		val, err := strconv.ParseFloat(strVal, 32)
		utils.CheckError(err)
		return float32(val)
	}
	return def
}

func getBoolOption(dict map[string]string, key string, def bool) bool {
	if strVal, ok := dict[key]; ok {
		val, err := strconv.ParseBool(strVal)
		utils.CheckError(err)
		return bool(val)
	}
	return def
}
