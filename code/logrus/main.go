package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

func main() {
	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})
	// log an event as usual with logrus
	log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from golang")
	contextLogger := log.WithFields(log.Fields{
		"common": "XXX common content XXX",
		"other":  "YYY special context YYY",
	})

	contextLogger.Info("AAAAAAAAAAAA")
	contextLogger.Info("BBBBBBBBBBBB")
}

//http://www.jianshu.com/p/a855e91c5335
//http://www.jianshu.com/p/5fac8bed4505
