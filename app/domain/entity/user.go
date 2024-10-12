package entity

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(255);not null"`
}
