package fakes

import (
	"fmt"

	"github.com/dinnerdonebetter/backend/pkg/types"

	fake "github.com/brianvoe/gofakeit/v5"
)

// BuildFakeOAuth2Client builds a faked OAuth2Client.
func BuildFakeOAuth2Client() *types.OAuth2Client {
	return &types.OAuth2Client{
		ID:           BuildFakeID(),
		Name:         fake.Password(true, true, true, false, false, 32),
		ClientID:     BuildFakeID(),
		ClientSecret: BuildFakePassword(),
		CreatedAt:    BuildFakeTime(),
	}
}

// BuildFakeOAuth2ClientCreationResponseFromOAuth2Client builds a faked OAuth2ClientCreationResponse.
func BuildFakeOAuth2ClientCreationResponseFromOAuth2Client(client *types.OAuth2Client) *types.OAuth2ClientCreationResponse {
	return &types.OAuth2ClientCreationResponse{
		ID:           client.ID,
		ClientID:     client.ClientID,
		ClientSecret: client.ClientSecret,
	}
}

// BuildFakeOAuth2ClientList builds a faked OAuth2ClientList.
func BuildFakeOAuth2ClientList() *types.QueryFilteredResult[types.OAuth2Client] {
	var examples []*types.OAuth2Client
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeOAuth2Client())
	}

	return &types.QueryFilteredResult[types.OAuth2Client]{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		Data: examples,
	}
}

// BuildFakeOAuth2ClientCreationInput builds a faked OAuth2ClientCreationRequestInput.
func BuildFakeOAuth2ClientCreationInput() *types.OAuth2ClientCreationRequestInput {
	client := BuildFakeOAuth2Client()

	return &types.OAuth2ClientCreationRequestInput{
		UserLoginInput: types.UserLoginInput{
			Username:  fake.Username(),
			Password:  BuildFakePassword(),
			TOTPToken: fmt.Sprintf("0%s", fake.Zip()),
		},
		Name: client.Name,
	}
}
