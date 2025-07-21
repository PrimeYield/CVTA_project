package setting

import (
	"log"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/spf13/viper"
)

type ServerSetting struct {
	Port string `json:"httpport"`
}

type DatabaseSetting struct {
	MongodbHost string
	MongodbPort string
	Mongodb_db  string
}

type JWTSetting struct {
	Algorithm jwa.SignatureAlgorithm
	Secret string
	Issuer string
	Expire time.Duration
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./config")
	err := vp.ReadInConfig()
	if err != nil {
		log.Printf("package setting NewSetting error: %v",err)
		return nil,err
	}
	return &Setting{vp:vp},nil
}

func (s *Setting) ReadSection(key string,value interface{}) error {
	err := s.vp.UnmarshalKey(key,value)
	if err != nil {
		log.Printf("package setting ReadSection error: %v",err)
		return err
	}
	return nil
}