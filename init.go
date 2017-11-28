package cntk

import (
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "cntk")
		if !supportedSystem {
			log.Error("cntk is only available on linux/amd64. not registering cntk")
		}
	})
}
