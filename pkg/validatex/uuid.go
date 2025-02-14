package validatex

import "github.com/google/uuid"

func ValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
