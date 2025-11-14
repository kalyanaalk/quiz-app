package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
}

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}
