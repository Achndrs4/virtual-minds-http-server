package models

import "time"

type HourlyStat struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerID   uint      `gorm:"type:int(11) unsigned;index"`
	Customer     Customer  `gorm:"foreignKey:CustomerID;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	Time         time.Time `gorm:"type:timestamp;not null"`
	RequestCount uint64    `gorm:"type:bigint(20) unsigned;default:0;not null"`
	InvalidCount uint64    `gorm:"type:bigint(20) unsigned;default:0;not null"`
}
