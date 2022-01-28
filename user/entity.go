package user

import "time"

type User struct {
	ID             int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name           string    `gorm:"column:name;type:varchar;size:255"`
	Occupation     string    `gorm:"column:occupation;type:varchar;size:255"`
	Email          string    `gorm:"column:email;type:varchar;size:255"`
	PasswordHash   string    `gorm:"column:password_hash;type:varchar;size:255"`
	AvatarFileName string    `gorm:"column:avatar_file_name;type:varchar;size:255"`
	Role           string    `gorm:"column:role;type:varchar;size:255"`
	Token          string    `gorm:"column:token;type:varchar;size:255"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp"`
}
