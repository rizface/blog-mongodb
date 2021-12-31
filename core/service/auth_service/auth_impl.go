package auth_service

import (
	"blog-mongo/app/exception"
	"blog-mongo/app/model/request"
	"blog-mongo/core/repository/auth_repository"
	"blog-mongo/helper"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type auth struct {
	db    *mongo.Database
	repo  auth_repository.Auth
	valid *validator.Validate
}

func NewAuth(db *mongo.Database, repo auth_repository.Auth, valid *validator.Validate) Auth {
	return &auth{
		db:    db,
		repo:  repo,
		valid: valid,
	}
}

func (a *auth) Register(request request.Register) string {
	var err error
	err = a.valid.Struct(request)
	helper.PanicCustomException(exception.BadRequest{Err: err},err != nil)
	password, err := helper.GeneratePassword(request.Password)
	helper.Panic(err)
	request.Password = password
	ctx, cancel := helper.CreateCtx(10)
	defer cancel()
	success := a.repo.Register(a.db, ctx, request)
	if success {
		return "registrasi berhasil"
	}
	return "registrasi gagal"
}

func (a *auth) Login(request request.Login) map[string]string {
	ctx,cancel := helper.CreateCtx(5)
	defer cancel()
	user,err := a.repo.Login(a.db,ctx,request.Email)
	helper.PanicCustomException(exception.NotFound{Err: errors.New("akun tidak terdaftar")},errors.Is(err,mongo.ErrNoDocuments))
	err = helper.ComparePassword(request.Password,user["password"].(string))
	helper.PanicCustomException(exception.BadRequest{Err:errors.New("username / password salah")},err != nil)
	token,err := helper.GenerateToken(user)
	helper.Panic(err)
	return map[string]string {
		"token": token,
	}
}
