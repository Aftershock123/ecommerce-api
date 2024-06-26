package models

import (
	"gorm.io/gorm"
    "time"
)

type Order struct {
    gorm.Model
    UserID     uint
    OrderDate  time.Time
    PaidDate   *time.Time
    Status     string
    OrderItems []OrderItem
    ShippingAddress ShippingAddress
}
