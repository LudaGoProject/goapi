package auth

import (
	// "gopkg.in/authboss.v1"
	log "gopkg.in/Sirupsen/logrus.v0"
	_ "gopkg.in/authboss.v1/auth"
)

func Init() {
	log.Debug("init auth config")
}
