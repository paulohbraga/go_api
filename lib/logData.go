package lib

import "github.com/gin-gonic/gin"

func getuserIpAdress(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

func LogService(text string, c *gin.Context) {
	ip := getuserIpAdress(c)
	WriteFile(text, "api.log", ip)
}
