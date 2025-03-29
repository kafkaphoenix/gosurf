//go:build unit

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kafkaphoenix/gosurf/internal/repository"
	"github.com/kafkaphoenix/gosurf/internal/repository/server"
	"github.com/kafkaphoenix/gosurf/internal/usecases"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	db      *repository.FakeDB
	service *usecases.UserService
	handler *server.UserHandler
	srv     *httptest.Server
	router  *http.ServeMux
}

func (s *UserHandlerTestSuite) SetupSuite() {
	// set up db
	db, err := repository.NewFakeDB("../../../db/users.json", "../../../db/actions.json")
	s.NoError(err, "could no create DB")

	// set up usecase and handler
	s.service = usecases.NewUserService(db)
	s.handler = server.NewUserHandler(s.service)

	// create router, testserver and register routes
	s.router = http.NewServeMux()
	s.handler.RegisterRoutes(s.router)
	s.srv = httptest.NewServer(s.router)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) TestGetUserByID() {
	// GIVEN
	req := httptest.NewRequest("GET", s.srv.URL+"/v1/user/1/", nil)
	rec := httptest.NewRecorder()

	// WHEN
	s.srv.Config.Handler.ServeHTTP(rec, req)

	// THEN
	s.Equal(http.StatusOK, rec.Code, "Expected status 200")
	expectedResponse := `{"id":1,"name":"Ferdinande","createdAt":"2020-07-14T05:48:54.798Z"}`
	s.Equal(expectedResponse, rec.Body.String())
}
