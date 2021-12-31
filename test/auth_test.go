package test

import (
	request2 "blog-mongo/app/model/request"
	"blog-mongo/app/setup"
	"blog-mongo/core/repository/auth_repository"
	"blog-mongo/core/service/auth_service"
	"blog-mongo/helper"
	"blog-mongo/route"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		password,_ := helper.GeneratePassword("rahasia")
		request := request2.Register{
			Username: "fariz",
			Email:    fmt.Sprintf("%s@gmail.com",strconv.Itoa(rand.Int())),
			Password: password,
		}
		repo := auth_repository.NewAuth()
		db,_ := helper.Connection()
		ctx,cancel  := context.WithTimeout(context.Background(),10 * time.Second)
		defer cancel()
		result := repo.Register(db,ctx,request)
		assert.True(t, result)
		t.Run("test service register", func(t *testing.T) {
			db,_ := helper.Connection()
			service := auth_service.NewAuth(db,repo,validator.New())
			t.Run("success", func(t *testing.T) {
				request.Email = strconv.Itoa(rand.Int()) + "@gmail.com"
				request.Username = strconv.Itoa(rand.Int())
				result := service.Register(request)
				assert.Equal(t, "registrasi berhasil", result)
			})
		})
		t.Run("test controller register", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				register := request2.Register{
					Username:	strconv.Itoa(rand.Int()),
					Email:    strconv.Itoa(rand.Int()) + "@gmail.com",
					Password: "rahasia",
				}
				requestJson,_ := json.Marshal(register)
				requesTest := httptest.NewRequest(http.MethodPost,route.REGISTER,bytes.NewReader(requestJson))
				recorder := httptest.NewRecorder()
				router := setup.AuthRouter()
				router.ServeHTTP(recorder,requesTest)
				assert.Equal(t, http.StatusOK,recorder.Code)
			})
			t.Run("failed", func(t *testing.T) {
				register := request2.Register{}
				requestJson,_ := json.Marshal(register)
				requesTest := httptest.NewRequest(http.MethodPost,route.REGISTER,bytes.NewReader(requestJson))
				recorder := httptest.NewRecorder()
				router := setup.AuthRouter()
				router.ServeHTTP(recorder,requesTest)
				assert.Equal(t, http.StatusBadRequest,recorder.Code)
			})
		})
	})

	t.Run("failed (duplicate)", func(t *testing.T) {
		password,_ := helper.GeneratePassword("rahasia")
		request := request2.Register{
			Username: "fariz",
			Email:    "fariz@gmail.com",
			Password: password,
		}
		repo := auth_repository.NewAuth()
		db,_ := helper.Connection()
		ctx,cancel  := context.WithTimeout(context.Background(),10 * time.Second)
		defer cancel()
		repo.Register(db,ctx,request)
		result := repo.Register(db,ctx,request)
		assert.False(t, result)
	})

}


func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		loginPayload := request2.Login{
			Email:    "fariz@gmail.com",
			Password: "rahasia",
		}
		requestJson,_ := json.Marshal(loginPayload)
		request := httptest.NewRequest(http.MethodPost,route.LOGIN,bytes.NewReader(requestJson))
		recorder := httptest.NewRecorder()
		router := setup.AuthRouter()
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("failed_wrong_password", func(t *testing.T) {
		loginPayload := request2.Login{
			Email:    "fariz@gmail.com",
			Password: "rahasias",
		}
		requestJson,_ := json.Marshal(loginPayload)
		request := httptest.NewRequest(http.MethodPost,route.LOGIN,bytes.NewReader(requestJson))
		recorder := httptest.NewRecorder()
		router := setup.AuthRouter()
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
	t.Run("failed_account_not_found", func(t *testing.T) {
		loginPayload := request2.Login{}
		requestJson,_ := json.Marshal(loginPayload)
		request := httptest.NewRequest(http.MethodPost,route.LOGIN,bytes.NewReader(requestJson))
		recorder := httptest.NewRecorder()
		router := setup.AuthRouter()
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
}