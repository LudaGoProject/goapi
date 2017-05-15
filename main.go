package main

import (
	_ "github.com/LudaGoProject/goapi/config/auth"
	_ "github.com/LudaGoProject/goapi/config/log"
	log "gopkg.in/Sirupsen/logrus.v0"
)

func main() {
	log.Error("test")
}
