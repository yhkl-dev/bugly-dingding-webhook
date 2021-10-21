package bugly

import (
	"fmt"
	"testing"
)

func TestSendDingDingMessage(t *testing.T) {
	config := GetDingdingConfig("../config.yml")
	fmt.Println(config.Dingding.DingdingAccessToken)
	type args struct {
		title       string
		content     string
		dingdingURL string
		secretKey   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"x",
			args{"Test A", `**Version:1.2.3**
			- 联网用户数: 12972
			- crash次数: 21
			- crash影响用户数：20`, config.Dingding.DingdingURL, config.Dingding.DingdingAccessToken},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendDingDingMessage(tt.args.title, tt.args.content, tt.args.dingdingURL, tt.args.secretKey)
		})
	}
}
