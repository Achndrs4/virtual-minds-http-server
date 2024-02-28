package models

type IPBlacklist struct {
	IP string `gorm:"primaryKey;varchar(255);NOT NULL"`
}
