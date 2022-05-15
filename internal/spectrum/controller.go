package spectrum

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetDiscovery(c *gin.Context) {

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
