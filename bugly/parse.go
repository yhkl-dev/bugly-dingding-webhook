package bugly

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type CrashTrendData struct {
	AccessUser int    `json:"accessUser"`
	CrashCount int    `json:"crashCount"`
	CrashUser  int    `json:"crashUser"`
	Version    string `json:"version"`
}

type EventContent struct {
	Datas      []CrashTrendData `json:"datas"`
	AppId      string           `json:"appId"`
	PlatformId int              `json:"platformId"`
	AppName    string           `json:"appName"`
	Date       string           `json:"date"`
	AppUrl     string           `json:"appUrl"`
}

type CrashTrend struct {
	EventType    string       `json:"eventType"`
	Timestamp    int          `json:"timestamp"`
	IsEncrypt    int          `json:"isEncrypt"`
	EventContent EventContent `json:"eventContent"`
	Signature    string       `json:"signature"`
}

func ParseJsonString(data map[string]interface{}) string {

	crashData := CrashTrend{}
	err := mapstructure.Decode(data, &crashData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(crashData.EventContent.AppName)

	TextData := `**APP Name: %s** \n`
	res := fmt.Sprintf(TextData, crashData.EventContent.AppName)
	crashInfo := `\n *Version:%s*
	- 联网用户数: %d
	- crash次数: %d
	- crash影响用户数：%d \n`

	for _, data := range crashData.EventContent.Datas {
		res = res + fmt.Sprintf(crashInfo, data.Version, data.AccessUser, data.CrashCount, data.CrashUser)
	}
	fmt.Println(res)
	return res
}
