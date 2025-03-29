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

type ActionHandlerTestSuite struct {
	suite.Suite
	db      *repository.FakeDB
	service *usecases.ActionService
	handler *server.ActionHandler
	srv     *httptest.Server
	router  *http.ServeMux
}

func (s *ActionHandlerTestSuite) SetupSuite() {
	// set up db
	db, err := repository.NewFakeDB("../../../db/users.json", "../../../db/actions.json")
	s.NoError(err, "could no create DB")

	// set up usecase and handler
	s.service = usecases.NewActionService(db)
	s.handler = server.NewActionHandler(s.service)

	// create router, testserver and register routes
	s.router = http.NewServeMux()
	s.handler.RegisterRoutes(s.router)
	s.srv = httptest.NewServer(s.router)
}

func TestActionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ActionHandlerTestSuite))
}

func (s *ActionHandlerTestSuite) TestGetUserByID() {
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
