package handlers

import (
	"crypto/sha256"
	"fmt"
)

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
