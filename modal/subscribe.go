package modal

import (
	"gorm.io/gorm"
)

type Subscribe struct {
	gorm.Model

	OpenID string     `json:"open_id"`
}
