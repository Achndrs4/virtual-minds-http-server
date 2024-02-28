package models

type Customer struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(255);not null"`
	Active bool   `gorm:"type:tinyint(1);default:1;not null"`
}
