package models

import "time"

type Note struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
