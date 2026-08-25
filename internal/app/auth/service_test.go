package authorization

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthService_SignUp_Success(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockAuthRepo(ctrl)
	hashRepo := NewMockPasswordHasher(ctrl)
	svc := NewAuthService(repo, hashRepo)

	testDTO := SignUpRequestDTO{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "123123",
	}

	hashRepo.EXPECT().Hash(testDTO.Password).Return("hashed:123123", nil)

	repo.EXPECT().CreateUser(SignUpRequestDTO{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "hashed:123123",
	}).Return(nil)

	err := svc.SignUp(testDTO)
	assert.NoError(t, err)
}

func TestAuthService_SignUp_Validation(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		CaseName   string
		RequestDTO SignUpRequestDTO
		wantErr    error
	}{
		{
			"NoErr",
			SignUpRequestDTO{
				Name:     "validName",
				Email:    "validmail@mail.com",
				Password: "123123",
			},
			nil,
		},
		{
			"NilBody",
			SignUpRequestDTO{},
			ErrBodyIsNil,
		},
		{
			"ErrNoName",
			SignUpRequestDTO{
				Name:     "",
				Email:    "validmail@mail.com",
				Password: "123123",
			},
			ErrNameIsEmpty,
		},
		{
			"ErrNoEmail",
			SignUpRequestDTO{
				Name:     "validName",
				Email:    "",
				Password: "123123",
			},
			ErrEmailIsEmpty,
		},
		{
			"ErrNoPassword",
			SignUpRequestDTO{
				Name:     "validName",
				Email:    "validmail@mail.com",
				Password: "",
			},
			ErrPasswordIsEmpty,
		},
		{
			"ErrWrongEmailFormat",
			SignUpRequestDTO{
				Name:     "validName",
				Email:    "bad%&email.com",
				Password: "123123",
			},
			ErrEmailName,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.CaseName, func(t *testing.T) {
			t.Parallel()
			req := SignUpRequestDTO{
				Name:     tt.RequestDTO.Name,
				Email:    tt.RequestDTO.Email,
				Password: tt.RequestDTO.Password,
			}
			err := validateSignUpRequest(&req)
			if tt.wantErr == nil {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
		})
	}
}
