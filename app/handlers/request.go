package request

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"nchan-auth/app/config"
	"nchan-auth/app/models"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	access_token := c.Query("access_token")
	uuid := c.Query("uuid")
	cfg := c.MustGet("cfg").(*config.Config)
	if access_token == "" || uuid == "" {
		c.JSON(400, gin.H{"status_code": 400, "error_msg": "Invalid Request"})
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	response, err := client.Get(cfg.Oauth2Url + "/tokeninfo?access_token=" + access_token)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	tokeninfo := new(models.TokenInfo)
	json.NewDecoder(response.Body).Decode(&tokeninfo)
	if tokeninfo.Uuid != uuid {
		c.JSON(401, gin.H{"status_code": 401, "error_msg": "Unauthorized"})
		return
	} else {
		c.JSON(200, gin.H{"status_code": 200})
	}
}

func Check(c *gin.Context) {
	c.String(200, "Hello! It's running!")
}
