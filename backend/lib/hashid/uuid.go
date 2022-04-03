package hashid

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	t := uuid.New()
	return t.String()
}
