package config

import (
	"github.com/LudaGoProject/goapi/config/auth"
	"github.com/LudaGoProject/goapi/config/env"
	"github.com/LudaGoProject/goapi/config/log"
	"github.com/LudaGoProject/goapi/config/service"
)

func init() {
	log.Init()
	env.Init()
	auth.Init()
	service.Init()
}
