package commons

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"github.com/orlmonteverde/go-notes-apirest/models"
)

func PreparePassword(u *models.User, compare bool) error {
	if compare && u.Password != u.ConfirmPassword {
		log.Println("Error: las contraseñas no coinciden")
		return errors.New("Las contraseñas no coinciden")
	}

	pwd := sha256.Sum256([]byte(u.Password))
	u.Password = fmt.Sprintf("%x", pwd)
	return nil
}
