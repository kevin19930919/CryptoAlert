package model

type Alert struct {
	ID        int64   `gorm:"primaryKey;autoIncrement:true;"`
	Crypto    string  `gorm:"unique;not null"`
	Direction bool    `gorm:"unique;not null"`
	Price     float64 `gorm:"unique;not null"`
}
