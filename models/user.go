package models

import "time"

type User struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name"`
	Password        string    `json:"password,omitempty"`
	ConfirmPassword string    `json:"confirm_password,omitempty"`
	Role            string    `json:"role,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
