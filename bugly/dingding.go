package bugly

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func SendDingDingMessage(title string, content string, dingdingURL string, secretKey string) {
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secretKey)
	sign := hmacSha256(stringToSign, secretKey)

	data := `{"msgtype":"markdown","markdown":{"title":"%s","text":"%s"},"at":{"atMobiles":[],"isAtAll":false}}`
	fmtDat := fmt.Sprintf(data, title, content)

	var jsonStr = []byte(fmtDat)
	buffer := bytes.NewBuffer(jsonStr)
	postURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", dingdingURL, timestamp, sign)
	request, err := http.NewRequest("POST", postURL, buffer)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
