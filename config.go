package main

import "github.com/owen-gxz/wechat-subscribe/db"

type Config struct {
	Port   string      `json:"port"`
	DB     db.Mysql `json:"db"`
	Wechat Wechat   `json:"wechat"`
	EveryTime string `json:"every_time"`
}

type Wechat struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Temps     []Temp `json:"temps"`
}

type Temp struct {
	ID     string   `json:"id"`
	Page string `json:"page"`
	Values map[string]string `json:"values"`
}
