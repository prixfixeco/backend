package fakes

import (
	"encoding/base32"
	"fmt"
	"log"
	"time"

	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/pquerna/otp/totp"
	"github.com/segmentio/ksuid"

	"github.com/prixfixeco/api_server/internal/authorization"
	"github.com/prixfixeco/api_server/pkg/types"
)

// BuildFakeUser builds a faked User.
func BuildFakeUser() *types.User {
	return &types.User{
		ID:                        ksuid.New().String(),
		EmailAddress:              fake.Email(),
		Username:                  fake.Password(true, true, true, false, false, 32),
		BirthDay:                  fake.Uint8(),
		BirthMonth:                fake.Uint8(),
		ServiceHouseholdStatus:    types.GoodStandingHouseholdStatus,
		TwoFactorSecret:           base32.StdEncoding.EncodeToString([]byte(fake.Password(false, true, true, false, false, 32))),
		TwoFactorSecretVerifiedOn: func(i uint64) *uint64 { return &i }(uint64(uint32(fake.Date().Unix()))),
		ServiceRoles:              []string{authorization.ServiceUserRole.String()},
		CreatedOn:                 uint64(uint32(fake.Date().Unix())),
	}
}

// BuildUserCreationResponseFromUser builds a faked UserCreationResponse.
func BuildUserCreationResponseFromUser(user *types.User) *types.UserCreationResponse {
	return &types.UserCreationResponse{
		CreatedUserID: user.ID,
		Username:      user.Username,
		CreatedOn:     user.CreatedOn,
	}
}

// BuildFakeUserList builds a faked UserList.
func BuildFakeUserList() *types.UserList {
	var examples []*types.User
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeUser())
	}

	return &types.UserList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		Users: examples,
	}
}

// BuildFakeUserCreationInput builds a faked UserRegistrationInput.
func BuildFakeUserCreationInput() *types.UserRegistrationInput {
	exampleUser := BuildFakeUser()

	return &types.UserRegistrationInput{
		Username:     exampleUser.Username,
		EmailAddress: fake.Email(),
		Password:     fake.Password(true, true, true, true, false, 32),
		BirthDay:     exampleUser.BirthDay,
		BirthMonth:   exampleUser.BirthMonth,
	}
}

// BuildFakeUserRegistrationInputFromUser builds a faked UserRegistrationInput.
func BuildFakeUserRegistrationInputFromUser(user *types.User) *types.UserRegistrationInput {
	return &types.UserRegistrationInput{
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Password:     fake.Password(true, true, true, true, false, 32),
		BirthDay:     user.BirthDay,
		BirthMonth:   user.BirthMonth,
	}
}

// BuildFakeUserDataStoreCreationInputFromUser builds a faked UserDataStoreCreationInput.
func BuildFakeUserDataStoreCreationInputFromUser(user *types.User) *types.UserDataStoreCreationInput {
	return &types.UserDataStoreCreationInput{
		ID:              user.ID,
		EmailAddress:    user.EmailAddress,
		Username:        user.Username,
		HashedPassword:  user.HashedPassword,
		TwoFactorSecret: user.TwoFactorSecret,
		BirthDay:        user.BirthDay,
		BirthMonth:      user.BirthMonth,
	}
}

// BuildFakeUserReputationUpdateInputFromUser builds a faked UserReputationUpdateInput.
func BuildFakeUserReputationUpdateInputFromUser(user *types.User) *types.UserReputationUpdateInput {
	return &types.UserReputationUpdateInput{
		TargetUserID:  ksuid.New().String(),
		NewReputation: user.ServiceHouseholdStatus,
		Reason:        fake.Sentence(10),
	}
}

// BuildFakeUserRegistrationInput builds a faked UserLoginInput.
func BuildFakeUserRegistrationInput() *types.UserRegistrationInput {
	return &types.UserRegistrationInput{
		Username:     fake.Username(),
		Password:     fake.Password(true, true, true, true, false, 32),
		EmailAddress: fake.Email(),
		BirthDay:     fake.Uint8(),
		BirthMonth:   fake.Uint8(),
	}
}

// BuildFakeUserLoginInputFromUser builds a faked UserLoginInput.
func BuildFakeUserLoginInputFromUser(user *types.User) *types.UserLoginInput {
	return &types.UserLoginInput{
		Username:  user.Username,
		Password:  fake.Password(true, true, true, true, false, 32),
		TOTPToken: fmt.Sprintf("0%s", fake.Zip()),
	}
}

// BuildFakePasswordUpdateInput builds a faked PasswordUpdateInput.
func BuildFakePasswordUpdateInput() *types.PasswordUpdateInput {
	return &types.PasswordUpdateInput{
		NewPassword:     fake.Password(true, true, true, true, false, 32),
		CurrentPassword: fake.Password(true, true, true, true, false, 32),
		TOTPToken:       fmt.Sprintf("0%s", fake.Zip()),
	}
}

// BuildFakeTOTPSecretRefreshInput builds a faked TOTPSecretRefreshInput.
func BuildFakeTOTPSecretRefreshInput() *types.TOTPSecretRefreshInput {
	return &types.TOTPSecretRefreshInput{
		CurrentPassword: fake.Password(true, true, true, true, false, 32),
		TOTPToken:       fmt.Sprintf("0%s", fake.Zip()),
	}
}

// BuildFakeTOTPSecretVerificationInput builds a faked TOTPSecretVerificationInput.
func BuildFakeTOTPSecretVerificationInput() *types.TOTPSecretVerificationInput {
	user := BuildFakeUser()

	token, err := totp.GenerateCode(user.TwoFactorSecret, time.Now().UTC())
	if err != nil {
		log.Panicf("error generating TOTP token for fakes user: %v", err)
	}

	return &types.TOTPSecretVerificationInput{
		UserID:    user.ID,
		TOTPToken: token,
	}
}

// BuildFakeTOTPSecretVerificationInputForUser builds a faked TOTPSecretVerificationInput for a given user.
func BuildFakeTOTPSecretVerificationInputForUser(user *types.User) *types.TOTPSecretVerificationInput {
	token, err := totp.GenerateCode(user.TwoFactorSecret, time.Now().UTC())
	if err != nil {
		log.Panicf("error generating TOTP token for fakes user: %v", err)
	}

	return &types.TOTPSecretVerificationInput{
		UserID:    user.ID,
		TOTPToken: token,
	}
}
