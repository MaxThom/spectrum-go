package spectrum

import (
	"net/http"

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
	SetAnimation(anim)
	c.JSON(http.StatusOK, anim)
}

func DeleteAnimation(c *gin.Context) {

}

func GetSettings(c *gin.Context) {

}

func PostSettings(c *gin.Context) {

}

func GetBrightness(c *gin.Context) {

}

func PostBrightness(c *gin.Context) {

}

func GetWifi(c *gin.Context) {

}

func PostWifi(c *gin.Context) {

}
