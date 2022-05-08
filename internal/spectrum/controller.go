package spectrum

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
)

var (
	mngr *AnimManager
)

func Run(log logr.Logger, args []string) {
	mngr = NewAnimManager(log, args)
	mngr.PlayDefaultAnimations()
}

func GetDiscovery(c *gin.Context) {

}

func GetAnimation(c *gin.Context) {
	c.JSON(http.StatusOK, mngr.GetAnimations())
}

func PostAnimation(c *gin.Context) {

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
