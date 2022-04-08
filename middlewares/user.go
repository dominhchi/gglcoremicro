package middlewares

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          string `json:"user_id"`
	IsSuperUser bool   `json:"is_superuser"`

	IsStaff         bool     `json:"is_staff"`
	StaffPermission []string `json:"staff_permissions"`

	IsAdmin     bool     `json:"is_admin"`
	Permissions []string `json:"permissions"`

	Buildings  []uuid.UUID `json:"buildings"`
	Blocks     []uuid.UUID `json:"blocks"`
	Apartments []uuid.UUID `json:"apartments"`
	Service    string      `json:"service"`
	Company    string      `json:"company"`
	IssuedAt   int64       `json:"iat"`
	ExpiredAt  int64       `json:"exp"`
}

func (payload *User) Valid() error {
	if time.Now().Unix() > payload.ExpiredAt {
		return ErrExpiredToken
	}
	return nil
}
