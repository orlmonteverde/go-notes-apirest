package commons

import (
	"encoding/json"
)

func JsonParser(i interface{}) (j []byte, err error) {
	j, err = json.Marshal(i)
	CheckError(err, "Error al parsear usuarios", false)
	return
}
