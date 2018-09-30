package commons

import "log"

func CheckError(err error, msg string, danger bool) (foundError bool) {
	if err == nil {
		return
	}

	foundError = true

	if danger {
		log.Fatal(msg, err)
	} else {
		log.Println(msg, err)
	}
	return
}
