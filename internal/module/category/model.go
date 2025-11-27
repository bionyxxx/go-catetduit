package category

import "time"

type Category struct {
	ID        uint      `json:"id" db:"id"`
	UserID    uint      `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
