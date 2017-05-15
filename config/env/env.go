package env

import (
	"fmt"
	"github.com/jinzhu/configor"
	log "gopkg.in/Sirupsen/logrus.v0"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Site     string
}

type DBConfig struct {
	Name     string
	Adapter  string
	Host     string
	Port     string
	User     string
	Password string
}

type AppConfig struct {
	// web port
	Port uint `default:"7000" env:"PORT"`
}

type EnvConfig struct {
	App  AppConfig
	DB   DBConfig
	SMTP SMTPConfig
}

var Config = EnvConfig{}

func Init() {
	log.Debug("init env config")
	if err := configor.Load(&Config, "resources/config/database.yml", "resources/config/smtp.yml", "resources/config.app.yml"); err != nil {
		panic(err)
	}
	log.Infof("load EnvConfig %v", Config)
}

func (this AppConfig) String() string {
	return fmt.Sprintf("app:[port:%d]", this.Port)
}

func (this SMTPConfig) String() string {
	return fmt.Sprintf("smtp:[host:%s,port:%s,user:%s,site:%s]", this.Host, this.Port, this.User, this.Site)
}

func (this DBConfig) String() string {
	return fmt.Sprintf("db:[name:%s,adapter:%s,host:%s,port:%s,user:%s]", this.Name, this.Adapter, this.Host, this.Port, this.User)
}

func (this EnvConfig) String() string {
	return fmt.Sprintf("env:[%v,%v,%v]", this.App, this.DB, this.SMTP)
}

func (s SMTPConfig) HostWithPort() string {
	return s.Host + ":" + s.Port
}
