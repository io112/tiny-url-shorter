package urlutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/io112/tiny-url-shorter/configs"
	"github.com/io112/tiny-url-shorter/internal/db"
	"github.com/io112/tiny-url-shorter/internal/encoder"
	"net/http"
	u "net/url"
	"strings"
)

// ShortUrl shorts given long url and returns shorten url.
func ShortUrl(c *gin.Context) {
	url := c.PostForm("url")
	if url == "" {
		sendError(c, http.StatusBadRequest, "url param is required")
		return
	}
	id, err := db.CreateURL(url)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	val, err := encoder.Encode(*id)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	resultUrl := fmt.Sprintf("%s:%s/%s", configs.HOST, configs.PORT, *val)
	c.JSON(http.StatusOK, gin.H{
		"url": resultUrl,
	})
}

// LongUrl finds long url by given short url.
func LongUrl(c *gin.Context) {
	formUrl := c.PostForm("url")
	if formUrl == "" {
		sendError(c, http.StatusBadRequest, "url param is required")
		return
	}
	url, err := u.Parse(formUrl)
	if err != nil {
		sendError(c, http.StatusBadRequest, "url parse error")
		return
	}

	shortCode, ok := getShortCode(url.Path)
	if !ok {
		sendError(c, http.StatusBadRequest, "wrong url")
		return
	}
	val, err := encoder.Decode(*shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.QueryURL(*val)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			sendError(c, http.StatusNotFound, err.Error())
		default:
			sendError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": res.Long,
	})
}

// sendError sends error to client
func sendError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"error": msg,
	})
}

// getShortCode checks url validity and returns shortCode
func getShortCode(path string) (*string, bool) {
	shortCodes := make([]string, 0)
	for _, val := range strings.Split(path, "/") {
		if val != "" {
			shortCodes = append(shortCodes, val)
		}
	}

	if len(shortCodes) != 1 {
		return nil, false
	}

	return &shortCodes[0], true
}
