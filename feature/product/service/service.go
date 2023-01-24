package service

import (
	"e-commerce-api/feature/product"
	"e-commerce-api/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	qry product.ProductData
	vld *validator.Validate
}

func New(d product.ProductData, v *validator.Validate) product.ProductService {
	return &productService{
		qry: d,
		vld: v,
	}
}

func (ps *productService) Add(token interface{}, newProduct product.Core, file multipart.File) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("token tidak valid")
	}

	if err := ps.vld.Struct(&newProduct); err != nil {
		msg := helper.ValidationErrorHandle(err)
		return errors.New(msg)
	}

	secureURL, err := helper.UploadFile(file)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "kesalahan input") {
			msg = err.Error()
		} else {
			msg = "gagal upload gambar karena kesalahan pada sistem server"
		}
		return errors.New(msg)
	}
	newProduct.Image = secureURL

	newProduct.Stock = 1

	if err := ps.qry.Add(uint(userID), newProduct); err != nil {
		return errors.New("kesalahan pada sistem server")
	}

	return nil
}

func (ps *productService) GetAll(page uint) ([]product.Core, error) {
	return nil, nil
}

func (ps *productService) GetByID(productID uint) (product.Core, error) {
	return product.Core{}, nil
}

func (ps *productService) Update(token interface{}, productID uint, updateProduct product.Core, fileHeader *multipart.FileHeader) error {
	return nil
}

func (ps *productService) Delete(token interface{}, productID uint) error {
	return nil
}
