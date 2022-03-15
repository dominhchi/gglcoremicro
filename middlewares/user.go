package middlewares

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID              string      `json:"user_id"`
	IsSuperUser     bool        `json:"is_superuser"`
	IsStaff         bool        `json:"is_staff"`
	StaffPermission []string    `json:"staff_permissions"`
	Buildings       []uuid.UUID `json:"buildings"`
	Apartments      []uuid.UUID `json:"apartments"`
	Service         string      `json:"service"`
	Company         string      `json:"service"`
	IssuedAt        int64       `json:"iat"`
	ExpiredAt       int64       `json:"exp"`
}

func (payload *User) Valid() error {
	if time.Now().Unix() > payload.ExpiredAt {
		return ErrExpiredToken
	}
	return nil
}
