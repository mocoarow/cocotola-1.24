package domain

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
