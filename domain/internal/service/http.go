package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const address = "http://127.0.0.1:23333/video"

func (s *Service) HttpGetVideo(ctx *gin.Context) {
	id := ctx.Query("id")
	resp, err := http.Get(address + fmt.Sprintf("?id=%s", id))
	if err != nil {
		logrus.Errorf("HTTP Get Video Error: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}
		ctx.String(http.StatusOK, string(body))
	} else {
		ctx.String(http.StatusInternalServerError, "Failed")
	}
}

func (s *Service) HttpCreateVideo(ctx *gin.Context) {
	title := ctx.PostForm("title")
	note := ctx.PostForm("note")
	pic := ctx.PostForm("pic")
	video := ctx.PostForm("video")

	resp, err := http.PostForm(address, url.Values{
		"title": {title},
		"note":  {note},
		"pic":   {pic},
		"video": {video},
	})

	if err != nil {
		logrus.Errorf("Post Form Error: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		ctx.String(http.StatusOK, "OK")
	} else {
		ctx.String(http.StatusInternalServerError, "Failed")
	}
}
