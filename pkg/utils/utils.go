package utils

import log "github.com/sirupsen/logrus"

func DoOrDie(err error) {
	DoOrDieWithMsg(err, "Fatal error: ")
}

func DoOrDieWithMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%s; err: %+v\n", msg, err)
	}
}