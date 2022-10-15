package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting()(*Setting,error){
     vp:= viper.New()
     vp.SetConfigName("config")
     //相对路径，可以设置多个配置路径，可以尽可能的尝试解决路径查找问题
     vp.AddConfigPath("configs/")
     vp.SetConfigType("yaml")
     err:=vp.ReadInConfig()
	if err != nil {
		return nil,err
	}
	return &Setting{vp:vp},nil

}