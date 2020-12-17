package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/owen-gxz/wechat-func"
	"github.com/owen-gxz/wechat-func/service"
	"github.com/owen-gxz/wechat-subscribe/modal"
	"github.com/jasonlvhit/gocron"
	"time"

	"io/ioutil"
	"net/http"
)

var (
	config    = &Config{}
	wxService = &service.Server{}
	gormDB    = &gorm.DB{}
)

func init() {
	configData, err := ioutil.ReadFile("./app.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	wxService = service.New(&service.WxConfig{
		AppId:  config.Wechat.AppID,
		Secret: config.Wechat.AppSecret,
	}, nil)

	if config.DB.Address == "" {
		gormDB, err = gorm.Open("gorm.db")
		if err != nil {
			panic(err)
		}
	} else {
		gormDB, err = config.DB.New()
	}

	go Tacker()
}

func Tacker() {
	gocron.Every(1).Days().At(config.EveryTime).Do(task)
	<-gocron.Start()
}

func task() {
	nt := time.Now()
	start := nt.AddDate(0, 0, -1).Format("2006-01-02")
	end := nt.Format("2006-01-02")
	ss := []modal.Subscribe{}
	err := gormDB.Model(&modal.Subscribe{}).
		Where("created_at >=? and created_at<=?", start, end).Find(&ss).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, wechatItem := range ss {
		for _, item := range config.Wechat.Temps {
			datas := make(map[string]wechat.SubscribeSendDate)
			for k, v := range item.Values {
				datas[k] = wechat.SubscribeSendDate{
					Value: v,
				}
			}
			ssr := wechat.SubscribeSendRequest{
				Touser:     wechatItem.OpenID,
				Page:       item.Page,
				TemplateID: item.ID,
				Data:       datas,
			}
			err := wxService.SubscribeSend(ssr)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/code2session", code2session)
	r.GET("/subscribe", subscribe)
	//r.GET("/send", subscribe)
	r.Run(config.Port)
}

func subscribe(c *gin.Context) {
	openID := c.Query("open_id")
	//nt:=time.Now()
	err := gormDB.Model(&modal.Subscribe{}).Create(&modal.Subscribe{
		OpenID: openID,
		//Time: &nt,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_code":    1,
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
	})
}

func code2session(c *gin.Context) {
	code := c.Query("code")
	wxSession, err := wxService.Jscode2Session(code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_code":    1,
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"data":       wxSession, // return all
	})
}
