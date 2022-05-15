package led

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/maxthom/spectrum-go/pkg/utils"
)

type LedstripOptions struct {
	Brightness     int    `json:"brightness,omitempty"`
	LedCount       int    `json:"ledCount,omitempty"`
	GpioPin        int    `json:"gpioPin,omitempty"`
	RenderWaitTime int    `json:"renderWaitTime,omitempty"`
	Frequency      int    `json:"frequency,omitempty"`
	DmaNum         int    `json:"dmaNum,omitempty"`
	StripType      string `json:"stripType,omitempty"`
}

func NewLedstripOptions() *LedstripOptions {
	return &LedstripOptions{
		Brightness:     128,
		LedCount:       144,
		GpioPin:        18,
		RenderWaitTime: 0,
		Frequency:      1200000,
		DmaNum:         10,
		StripType:      "SK6812StripGRBW",
	}
}

func NewFromFileLedstripOptions(filePath string) *LedstripOptions {
	result := NewLedstripOptions()
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsExist(err) {
			utils.CheckError(err)
		} else if os.IsNotExist(err) {
			utils.CheckError(result.WriteToFile(filePath))
			return result
		}
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	utils.CheckError(err)

	err = json.Unmarshal(data, &result)
	utils.CheckError(err)

	return result
}

func (s *LedstripOptions) WriteToFile(filePath string) error {
	file, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, file, 0777)
	if err != nil {
		return err
	}

	return nil
}
