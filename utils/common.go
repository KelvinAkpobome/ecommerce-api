package utils

import "github.com/go-playground/validator/v10"


type OrderStatus int

const (
	Pending OrderStatus = iota // 0
	Processed                   // 1
	Shipped                     // 2
	Delivered                   // 3
	Cancelled                   // 4
)

func (os OrderStatus) String() string {
	switch os {
	case Pending:
		return "Pending"
	case Processed:
		return "Processed"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}


type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Response interface{} `json:"response,omitempty"`
	Error   string      `json:"error,omitempty"`
}

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}