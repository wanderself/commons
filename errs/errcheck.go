package errs

import "log"

func ErrCheck(err error, content string) {
	if err != nil {
		log.Printf("%s: %s", content, err.Error())
	}
}
