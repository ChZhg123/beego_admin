package controllers

import (
	"blog/global"
	beego "github.com/beego/beego/v2/adapter"
	"time"
)

type BaseController struct {
	beego.Controller
}

type ResponseData struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Code int `json:"code"`
	Time string `json:"time"`
}

func (b *BaseController) Response(code int,message string,data interface{}){
	var jsonData *ResponseData = &ResponseData{message,data,code,time.Now().Format("2006-01-02 15:04:05")}
	b.Data["json"] = jsonData
	b.ServeJSON()
	b.StopRun()
}

func (b *BaseController) Success(message string,data interface{})  {
	b.Response(2,message,data)
}

func (b *BaseController) Error(message string)  {
	b.Response(1,message,global.Data)
}

type LayTableResponseData struct {
	Code int `json:"code"`
	Msg string 	`json:"msg"`
	Count int `json:"count"`
	Data interface{} `json:"data"`
}

func (b *BaseController) LayTableApi(count int64,data interface{})  {
	var jsonData *LayTableResponseData = &LayTableResponseData{0,"请求成功",int(count),data}
	b.Data["json"] = jsonData
	b.ServeJSON()
	b.StopRun()
}
