package spectrum

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/maxthom/spectrum-go/pkg/led"
	"github.com/maxthom/spectrum-go/pkg/utils"
)

func GetDiscovery(w http.ResponseWriter, r *http.Request) {
	result := struct {
		Options    *led.LedstripOptions `json:"options"`
		Animations []DiscoveryDO        `json:"animations"`
	}{
		Options: options,
		Animations: []DiscoveryDO{
			{
				Animation: "Clear",
				Options:   map[string]OptionDO{},
			},
			{
				Animation: "Wipe",
				Options: map[string]OptionDO{
					"wait": {
						Type:    "TimeMs",
						Default: "10",
						Min:     "0",
						Max:     "100",
					},
					"color": {
						Type:    "Color",
						Default: "0xff000000",
					},
					"reverse": {
						Type:    "Bool",
						Default: "False",
					},
				},
			},
			{
				Animation: "Rainbow",
				Options: map[string]OptionDO{
					"wait": {
						Type:    "TimeMs",
						Default: "5",
						Min:     "0",
						Max:     "100",
					},
				},
			},
			{
				Animation: "Maze",
				Options: map[string]OptionDO{
					"wait": {
						Type:    "TimeMs",
						Default: "30",
						Min:     "0",
						Max:     "100",
					},
					"color": {
						Type:    "Color",
						Default: "0x000000ff",
					},
					"contact_color": {
						Type:    "Color",
						Default: "0xff000000",
					},
					"count": {
						Type:    "Int",
						Default: "10",
						Min:     "1",
						Max:     "100",
					},
					"turn_chance": {
						Type:    "Int",
						Default: "2",
						Min:     "0",
						Max:     "100",
					},
				},
			},
		},
	}

	resultJson(w, http.StatusOK, result)
	log.V(0).Info(utils.ColoredText(utils.Green, "MuxRequest GET"), "handler", "GetDiscovery", "httpStatus", http.StatusOK)
}

func GetAnimation(w http.ResponseWriter, r *http.Request) {
	resultJson(w, http.StatusOK, animations)
	log.V(0).Info(utils.ColoredText(utils.Green, "MuxRequest GET"), "handler", "GetAnimation", "httpStatus", http.StatusOK)
}

func PostAnimation(w http.ResponseWriter, r *http.Request) {
	anim, err := getBody[AnimUnitDO](w, r)
	if err != nil {
		log.Error(err, utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostAnimation", "httpStatus", http.StatusBadRequest)
		return
	}

	if anim.Segment.Start > options.LedCount || anim.Segment.End > options.LedCount {
		resultJson(w, http.StatusOK, map[string]any{"message": "Segment start or end is higher then led count.", "segment": anim.Segment, "ledCount": options.LedCount})
		log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostAnimation", "httpStatus", http.StatusOK, "body", anim)
		return
	}

	if anim.Segment.Start >= anim.Segment.End {
		resultJson(w, http.StatusOK, map[string]any{"message": "Segment start is higer or equal to segment end.", "segment": anim.Segment, "ledCount": options.LedCount})
		log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostAnimation", "httpStatus", http.StatusOK, "body", anim)
		return
	}

	SetAnimation(*anim)
	resultJson(w, http.StatusOK, anim)
	log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostAnimation", "httpStatus", http.StatusOK, "body", anim)
}

func DeleteAnimation(w http.ResponseWriter, r *http.Request) {
	anim, err := getBody[AnimStopDO](w, r)
	if err != nil {
		log.Error(err, utils.ColoredText(utils.Yellow, "MuxRequest DELETE"), "handler", "DeleteAnimation", "httpStatus", http.StatusBadRequest)
		return
	}

	StopAnimation(*anim)

	resultJson(w, http.StatusOK, anim)
	log.V(0).Info(utils.ColoredText(utils.Yellow, "MuxRequest DELETE"), "handler", "DeleteAnimation", "httpStatus", http.StatusOK, "body", anim)
}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	resultJson(w, http.StatusOK, options)
	log.V(0).Info(utils.ColoredText(utils.Green, "MuxRequest GET"), "handler", "GetSettings", "httpStatus", http.StatusOK)
}

func PostSettings(w http.ResponseWriter, r *http.Request) {
	do, err := getBody[led.LedstripOptions](w, r)
	if err != nil {
		log.Error(err, utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostSettings", "httpStatus", http.StatusBadRequest)
		return
	}

	options = do
	options.WriteToFile(pathToConfig)

	resultJson(w, http.StatusOK, options)
	log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostSettings", "httpStatus", http.StatusOK, "body", do)
	os.Exit(0)
}

func GetBrightness(w http.ResponseWriter, r *http.Request) {
	resultJson(w, http.StatusOK, map[string]any{"brightness": options.Brightness})
	log.V(0).Info(utils.ColoredText(utils.Green, "MuxRequest GET"), "handler", "GetBrightness", "httpStatus", http.StatusOK)
}

func PostBrightness(w http.ResponseWriter, r *http.Request) {
	do, err := getBody[BrightnessDO](w, r)
	if err != nil {
		log.Error(err, utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostBrightness", "httpStatus", http.StatusBadRequest)
		return
	}

	if do.Brightness < 0 || do.Brightness > 255 {
		resultJson(w, http.StatusOK, map[string]any{"message": "Brightness must be between 0 and 255."})
		log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostBrightness", "httpStatus", http.StatusOK, "body", do)
		return
	}

	options.Brightness = do.Brightness
	strip.SetBrightness(0, do.Brightness)

	resultJson(w, http.StatusOK, do)
	log.V(0).Info(utils.ColoredText(utils.Cyan, "MuxRequest POST"), "handler", "PostBrightness", "httpStatus", http.StatusOK, "body", do)
}

func GetWifi(w http.ResponseWriter, r *http.Request) {

}

func PostWifi(w http.ResponseWriter, r *http.Request) {

}

func getBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return nil, err
	}

	var do T
	if err := json.Unmarshal(reqBody, &do); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return nil, err
	}

	return &do, nil
}

func resultJson(w http.ResponseWriter, statusCode int, r any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(r)
}
