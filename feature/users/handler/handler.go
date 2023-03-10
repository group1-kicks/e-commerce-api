package handler

import (
	"e-commerce-api/feature/users"
	"e-commerce-api/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userControl struct {
	srv users.UserService
}

func New(srv users.UserService) users.UserHandler {
	return &userControl{
		srv: srv,
	}
}

func (uc *userControl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Username, input.Password)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "berhasil login", ToResponse(res), token))
	}
}
func (uc *userControl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "berhasil mendaftar", ToResponse(res)))
	}
}
func (uc *userControl) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "berhasil lihat profil", MyProfile(res.(users.Core))))

	}
}

func (uc *userControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user")
		input := UpdateRequest{}

		//cek input json dengan format yang benar
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		_, err := uc.srv.Update(token, *ReqToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.SuccessResponse(http.StatusOK, "berhasil update profil"))
	}
}

func (uc *userControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := uc.srv.Delete(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "berhasil delete profil", err))
	}
}
