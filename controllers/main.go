package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SONOR_URL = "https://apigmining.xyz/api/sonar"
	DIG_URL   = "https://apigmining.xyz/api/dig"
	COIN_URL  = "http://150.95.134.56/coin"
	GET       = "GET"
	PUT       = "PUT"
	MAX_SONAR = 1
)

type Sonor struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}

type Message struct {
	Message string `json:"message"`
}

var API_1 string
var DIG Sonor

// Operations about object
type MainController struct {
	beego.Controller
}

// @Title Get
// @Description return coin position.
// @Success 200 {string} models.Sonor
// @Failure 403 body is empty
// @router /coins [get]
func (this *MainController) Get() {
	this.Data["json"] = &DIG
	this.ServeJSON()
}

func init() {
	API_1 = beego.AppConfig.String("api::key")
	go func() {
		t := time.NewTicker(60 * time.Second) // 指定時間置きに実行
		do()
		for {
			select {
			case <-t.C:
				do()
			}
		}

		t.Stop() // タイマを止める。
	}()
}

func do() {
	fmt.Println("start")
	apis := []string{API_1}
	// ソナー部
	s1 := Sonor{}
	s2 := Sonor{}
	s3 := Sonor{}
	s4 := Sonor{}
	sonors := []Sonor{s1, s2, s3, s4}
	for i := 0; i < MAX_SONAR; i++ {
		byteArray, err := Request(GET, SONOR_URL, apis[i])
		if err != nil {
			fmt.Println("err:", err)
		}
		if err := json.Unmarshal(byteArray, &sonors[i]); err != nil {
			fmt.Println("err:", err)
		}
	}
	//byteArray, err := Request(GET, COIN_URL, "")
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//if err := json.Unmarshal(byteArray, &sonors[3]); err != nil {
	//	fmt.Println("err:", err)
	//}
	fmt.Println(sonors)
	// dig部
	DIG.X = 0
	DIG.Y = 0
	DIG.Z = 0
	for i := 0; i < MAX_SONAR; i++ {
		d := sonors[i]
		if d.X > 0 {
			DIG.X = d.X
		}
		if d.Y > 0 {
			DIG.Y = d.Y
		}
		if d.Z > 0 {
			DIG.Z = d.Z
		}
	}
	fmt.Println(DIG)
	//if DIG.X > 0 && DIG.Y > 0 && DIG.Z > 0 {
	//	ret := doPut(DIG_URL, API_1, DIG)
	//	if ret == true {
	//		time.Sleep(1 * time.Second)
	//		do()
	//	}
	//}
}

func Request(method string, uri string, api string) ([]byte, error) {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	if api != "" {
		req.Header.Set("Authorization", "Bearer "+api)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	return byteArray, err
}
func doPut(url string, api string, dig Sonor) bool {
	ret := false
	client := &http.Client{}
	b, err := json.Marshal(dig)
	if err != nil {
		fmt.Println(err)
		return ret
	}
	fmt.Println(string(b))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return ret
	}
	req.Header.Set("Authorization", "Bearer "+api)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ret
	} else {
		defer response.Body.Close()
		ret = true
		return ret
	}
}
