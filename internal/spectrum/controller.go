package spectrum

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maxthom/spectrum-go/pkg/led"
)

func GetDiscovery(c *gin.Context) {
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
						Default: "1",
					},
					"color": {
						Type:    "Color",
						Default: "0xff000000",
					},
				},
			},
			{
				Animation: "Rainbow",
				Options: map[string]OptionDO{
					"wait": {
						Type:    "TimeMs",
						Default: "1",
					},
				},
			},
			{
				Animation: "Maze",
				Options: map[string]OptionDO{
					"wait": {
						Type:    "TimeMs",
						Default: "1",
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
					},
					"turn_chance": {
						Type:    "Int",
						Default: "2",
					},
				},
			},
		},
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

func GetAnimation(c *gin.Context) {
	c.JSON(http.StatusOK, animations)
}

func PostAnimation(c *gin.Context) {
	var anim AnimUnitDO
	if err := c.ShouldBindJSON(&anim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if anim.Segment.Start > options.LedCount || anim.Segment.End > options.LedCount {
		c.JSON(http.StatusOK, gin.H{"message": "Segment start or end is higher then led count.", "segment": anim.Segment, "ledCount": options.LedCount})
		return
	}

	if anim.Segment.Start >= anim.Segment.End {
		c.JSON(http.StatusOK, gin.H{"message": "Segment start is higer or equal to segment end.", "segment": anim.Segment, "ledCount": options.LedCount})
		return
	}

	SetAnimation(anim)
	c.JSON(http.StatusOK, anim)
}

func DeleteAnimation(c *gin.Context) {
	var anim AnimStopDO
	if err := c.ShouldBindJSON(&anim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	StopAnimation(anim)
	c.JSON(http.StatusOK, anim)
}

func GetSettings(c *gin.Context) {
	c.JSON(http.StatusOK, options)
}

func PostSettings(c *gin.Context) {
	if err := c.ShouldBindJSON(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	options.WriteToFile(pathToConfig)
	c.JSON(http.StatusOK, options)
	// Restart to load new config
	os.Exit(0)
}

func GetBrightness(c *gin.Context) {
	c.JSON(http.StatusOK, options.Brightness)
}

func PostBrightness(c *gin.Context) {
	var do BrightnessDO
	if err := c.ShouldBindJSON(&do); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	strip.SetBrightness(0, do.Brightness)
	c.JSON(http.StatusOK, do)
}

func GetWifi(c *gin.Context) {

}

func PostWifi(c *gin.Context) {

}
