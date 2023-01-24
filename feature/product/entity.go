package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Name        string `validate:"min=3"`
	Description string `validate:"min=5"`
	SellerName  string
	City        string
	Price       int `validate:"gte=10000"`
	Stock       int
	Image       string
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ProductService interface {
	Add(token interface{}, newProduct Core, fileHeader multipart.File) error
	GetAll(page uint) ([]Core, error)
	GetByID(productID uint) (Core, error)
	Update(token interface{}, productID uint, updateProduct Core, fileHeader *multipart.FileHeader) error
	Delete(token interface{}, productID uint) error
}

type ProductData interface {
	Add(userID uint, newProduct Core) error
	GetAll(limit int, offset int) ([]Core, error)
	GetByID(productID uint) (Core, error)
	Update(userID uint, productID uint, updateProduct Core) error
	Delete(userID uint, productID uint) error
}
