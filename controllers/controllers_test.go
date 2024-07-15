package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gogineni1998/go-api/routes"
	"github.com/gogineni1998/go-api/utilities"
	"github.com/labstack/echo/v4"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("User Controllers", func() {

	ginkgo.Describe("GetUsers", func() {
		ginkgo.It("should return a list of users", func() {
			router := routes.Routes()
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			router.ServeHTTP(recorder, req)
			fmt.Println(recorder.Body.String())
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusAccepted))
		})
	})

	ginkgo.Describe("GetUser", func() {
		ginkgo.It("should return a single user", func() {
			req := httptest.NewRequest(http.MethodGet, "/users/1565", nil)
			router := routes.Routes()
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			fmt.Println(recorder.Body.String())
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusAccepted))
			gomega.Expect(recorder.Body.String()).To(gomega.ContainSubstring("gsn"))
		})
	})

	ginkgo.Describe("CreateUser", func() {
		ginkgo.It("should create a new user", func() {
			utilities.CreateTable()
			router := routes.Routes()
			recorder := httptest.NewRecorder()
			userJSON := `{"id":"1","username":"gogineni","email":"gogineni1998@gmail.com","summary":"a"}`
			fmt.Println(userJSON)
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			router.ServeHTTP(recorder, req)
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusCreated))
			gomega.Expect(recorder.Body.String()).To(gomega.ContainSubstring("User 1 Created Successfully"))
		})
	})

	ginkgo.Describe("UpdateUser", func() {
		ginkgo.It("should update an existing user", func() {
			router := routes.Routes()
			recorder := httptest.NewRecorder()
			userJSON := `{"id":"1","username":"gogineni","email":"gogineni1998@gmail.com","summary":"a"}`
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			router.ServeHTTP(recorder, req)
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusAccepted))
			gomega.Expect(recorder.Body.String()).To(gomega.ContainSubstring("updated successfully"))
		})
	})

	ginkgo.Describe("DeleteUser", func() {
		ginkgo.It("should delete an existing user", func() {
			router := routes.Routes()
			req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusAccepted))
			gomega.Expect(recorder.Body.String()).To(gomega.ContainSubstring("deleated successfully"))
		})
	})
})
