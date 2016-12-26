package errs

import "log"

func ErrCheck(err error, content string) error {
	if err != nil {
		log.Printf("%s: %s", err, content)
		return err
	} else {
		return nil
	}
}
