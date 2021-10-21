package bugly

import (
	"testing"
)

func TestParseJsonString(t *testing.T) {
	var x string = `{
		"eventType": "bugly_crash_trend",
		"timestamp": 1462780713515,
		"isEncrypt": 0,
		"eventContent": {
		  "datas": [
			{
			  "accessUser": 12972,
			  "crashCount": 21,
			  "crashUser": 20,
			  "version": "1.2.3",
			  "url": "http://bugly.qq.com/realtime?app=1104512706&pid=1&ptag=1005-10003&vers=0.0.0.12.12&time=last_7_day&tab=crash"
			},
			{
			  "accessUser": 15019,
			  "crashCount": 66,
			  "crashUser": 64,
			  "version": "1.2.4",
			  "url": "http://bugly.qq.com/realtime?app=1104512706&pid=1&ptag=1005-10003&vers=0.0.0.12.12&time=last_7_day&tab=crash"
			},
			{
			  "accessUser": 15120,
			  "crashCount": 1430,
			  "crashUser": 1423,
			  "version": "1.2.4",
			  "url": "http://bugly.qq.com/realtime?app=1104512706&pid=1&ptag=1005-10003&vers=0.0.0.12.12&time=last_7_day&tab=crash"
			}
		  ],
		  "appId": "1104512706", 
		  "platformId": 1,  
	  	"appName": "AF", 
		  "date": "20160508",
	  "appUrl":"http://bugly.qq.com/issueIndex?app=1104512706&pid=1&ptag=1005-10000" 
		},
		"signature": "ACE346A4AE13A23A52A0D0D19350B466AF51728A"
	  }
	  `
	type args struct {
		data string
	}
	want := "res"
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test A", args{x}, want},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseJsonString(tt.args.data); got != tt.want {
				t.Errorf("ParseJsonString() = %v, want %v", got, tt.want)
			}
		})
	}
}
