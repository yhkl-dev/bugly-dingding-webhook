package bugly

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Dingding Dingding
}

type Dingding struct {
	DingdingURL         string `yaml:"dingdingurl"`
	DingdingAccessToken string `yaml:"accessToken"`
}

//GetConfig 获取配置数据
func GetDingdingConfig(filePath string) Config {
	config := Config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("解析config.yml读取错误: %v", err)
	}

	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}
