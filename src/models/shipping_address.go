package models

import "gorm.io/gorm"

type ShippingAddress struct {
    gorm.Model
    OrderID       uint
    RecipientName string
    Address       string
    City          string
    PostalCode    string
    Country       string
}