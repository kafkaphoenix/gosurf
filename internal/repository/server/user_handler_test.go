//go:build unit

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kafkaphoenix/gosurf/internal/repository/db"
	"github.com/kafkaphoenix/gosurf/internal/repository/server"
	"github.com/kafkaphoenix/gosurf/internal/usecases"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	db      *db.FakeDB
	service *usecases.UserService
	handler *server.UserHandler
	srv     *httptest.Server
	router  *http.ServeMux
}

func (s *UserHandlerTestSuite) SetupSuite() {
	// set up db
	db, err := db.NewFakeDB("../../../db/users.json", "../../../db/actions.json")
	s.NoError(err, "could not create DB")

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
	req := httptest.NewRequest(http.MethodGet, "/v1/users/1", nil)
	rec := httptest.NewRecorder()

	// WHEN
	s.srv.Config.Handler.ServeHTTP(rec, req)

	// THEN
	s.Equal(http.StatusOK, rec.Code)
	expectedResponse := `{"id":1,"name":"Ferdinande","createdAt":"2020-07-14T05:48:54.798Z"}`
	s.Equal(expectedResponse, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestGetTotalActionsByID() {
	// GIVEN
	req := httptest.NewRequest(http.MethodGet, "/v1/users/1/actions/total", nil)
	rec := httptest.NewRecorder()

	// WHEN
	s.srv.Config.Handler.ServeHTTP(rec, req)

	// THEN
	s.Equal(http.StatusOK, rec.Code)
	expectedResponse := `{"count":49}`
	s.Equal(expectedResponse, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestGetRefferalsIndex() {
	// GIVEN
	req := httptest.NewRequest(http.MethodGet, "/v1/referral-index", nil)
	rec := httptest.NewRecorder()

	// WHEN
	s.srv.Config.Handler.ServeHTTP(rec, req)

	// THEN
	s.Equal(http.StatusOK, rec.Code)
}
