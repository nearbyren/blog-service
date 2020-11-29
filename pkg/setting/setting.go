package setting

import (
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   //设置文件名称
	vp.AddConfigPath("configs/") //相对路径
	vp.SetConfigType("yaml")     //配置类型
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	log.Println("22222")
	return &Setting{vp}, nil
}
