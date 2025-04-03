package usecases_test

import (
	"testing"
	"time"

	"github.com/kafkaphoenix/gosurf/internal/domain"
	"github.com/kafkaphoenix/gosurf/internal/mocks"
	"github.com/kafkaphoenix/gosurf/internal/usecases"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	mockDBRepo  *mocks.DBRepo
	userService *usecases.UserService
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
func (s *UserServiceTestSuite) SetupSuite() {
	s.mockDBRepo = mocks.NewDBRepo(s.T())
	s.userService = usecases.NewUserService(s.mockDBRepo)
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.mockDBRepo.AssertExpectations(s.T())
}

func (s *UserServiceTestSuite) TestGetUserByID_OK() {
	// GIVEN
	expectedUser := domain.User{
		ID: 0, Name: "Allyson", CreatedAt: time.Now(),
	}

	s.mockDBRepo.EXPECT().GetUser(0).Return(expectedUser, true).Once()

	// WHEN
	user, err := s.userService.GetUserByID(0)

	// THEN
	s.Require().NoError(err)
	s.Equal(expectedUser, *user)
}

// but usually they wouldn't use mocks or a test suite.
func (s *UserServiceTestSuite) TestGetUserByID() {
	tests := map[string]struct {
		user     domain.User
		expected domain.User
		noError  bool
	}{
		"ok": {
			user: domain.User{
				ID: 0, Name: "Allyson", CreatedAt: time.Now(),
			},
			expected: domain.User{
				ID: 0, Name: "Allyson", CreatedAt: time.Now(),
			},
			noError: true,
		},
		"ko": {
			user: domain.User{
				ID: 0, Name: "Allyson", CreatedAt: time.Now(),
			},
			expected: domain.User{
				ID: 0, Name: "Allyson", CreatedAt: time.Now(),
			},
			noError: false,
		},
	}

	for name, tc := range tests {
		s.T().Run(name, func(_ *testing.T) {
			s.mockDBRepo.EXPECT().GetUser(tc.user.ID).Return(tc.expected, tc.noError).Once()
			user, err := s.userService.GetUserByID(tc.user.ID)

			if tc.noError {
				s.Require().NoError(err)
				s.Equal(tc.expected, *user)
			} else {
				s.Require().Error(err)
				s.Nil(user)
			}
		})
	}
}

func (s *UserServiceTestSuite) TestGetUserByID_KO() {
	// GIVEN
	expectedErr := &usecases.ServiceError{Message: "user with id 0 not found"}

	s.mockDBRepo.EXPECT().GetUser(0).Return(domain.User{}, false).Once()

	// WHEN
	user, err := s.userService.GetUserByID(0)

	// THEN
	s.Require().Error(err, &expectedErr)
	s.Nil(user)
}

func (s *UserServiceTestSuite) TestReferralIndex_OK() {
	// GIVEN
	users := map[int]domain.User{
		0: {ID: 0, Name: "Allyson", CreatedAt: time.Now()},
		1: {ID: 0, Name: "Ferdinande", CreatedAt: time.Now()},
		2: {ID: 0, Name: "Amelie", CreatedAt: time.Now()},
		3: {ID: 0, Name: "Aidan", CreatedAt: time.Now()},
	}
	referralGraph := map[int][]int{
		0: {1, 2}, // User 0 referred users 1 and 2
		1: {3},
		2: {},
		3: {},
	}
	expectedRefIndex := domain.ReferralIndex{
		0: domain.Referral{Count: 3},
		1: domain.Referral{Count: 1},
		2: domain.Referral{},
		3: domain.Referral{},
	}

	s.mockDBRepo.EXPECT().GetAllUsers().Return(users).Once()
	s.mockDBRepo.EXPECT().GetReferrals(0).Return(referralGraph[0], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(1).Return(referralGraph[1], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(2).Return(referralGraph[2], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(1).Return(referralGraph[1], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(2).Return(referralGraph[2], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()

	// WHEN
	refIndex := s.userService.GetReferralIndex()

	// THEN
	s.Equal(expectedRefIndex, refIndex)
}

func (s *UserServiceTestSuite) TestReferralIndex_OneTime__OK() {
	// GIVEN
	users := map[int]domain.User{
		0: {ID: 0, Name: "Allyson", CreatedAt: time.Now()},
		1: {ID: 0, Name: "Ferdinande", CreatedAt: time.Now()},
		2: {ID: 0, Name: "Amelie", CreatedAt: time.Now()},
		3: {ID: 0, Name: "Aidan", CreatedAt: time.Now()},
	}
	referralGraph := map[int][]int{
		0: {1, 2}, // User 0 referred users 1 and 2
		1: {3},
		2: {3},
		3: {},
	}
	expectedRefIndex := domain.ReferralIndex{
		0: domain.Referral{Count: 3},
		1: domain.Referral{Count: 1},
		2: domain.Referral{Count: 1},
		3: domain.Referral{},
	}

	s.mockDBRepo.EXPECT().GetAllUsers().Return(users).Once()
	s.mockDBRepo.EXPECT().GetReferrals(0).Return(referralGraph[0], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(1).Return(referralGraph[1], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(2).Return(referralGraph[2], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(1).Return(referralGraph[1], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(2).Return(referralGraph[2], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()
	s.mockDBRepo.EXPECT().GetReferrals(3).Return(referralGraph[3], true).Once()

	// WHEN
	refIndex := s.userService.GetReferralIndex()

	s.Equal(expectedRefIndex, refIndex)
}
