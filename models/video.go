package models

import (
	"encoding/json"
	"time"
)

type Video struct {
	ID          string    `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Duration    int       `json:"duration" db:"duration"`
	Key         string    `json:"key" db:"key"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Owner       string    `json:"owner" db:"owner"`
}

type VideoNested struct {
	ID          string    `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Duration    int       `json:"duration" db:"duration"`
	Key         string    `json:"key" db:"key"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Owner       *User     `json:"owner"`
}

// String is not required by pop and may be deleted
func (v Video) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Videos is not required by pop and may be deleted
type Videos []*Video

// String is not required by pop and may be deleted
func (v Videos) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}
