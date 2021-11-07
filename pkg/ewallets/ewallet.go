package ewallets

import "gorm.io/gorm"

type EwalletType uint8

const (
	TYPE_NOT_IDENTITY EwalletType = iota
	TYPE_IDENTITY
)

type Ewallet struct {
	gorm.Model
	UserId int64 `gorm:"not null"`
	Balance int64
	Type EwalletType `gorm:"not null;default:0"`
}
