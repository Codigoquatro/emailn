package campaign

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaingn{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"alves@gmail.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	newCampaing.Name = ""

	_, err := service.Create(newCampaing)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())

}

func Test_Create_SaveCampaign(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaing.Name ||
			campaign.Content != newCampaing.Content ||
			len(campaign.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}

	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaing)

	assert.Equal("error to save on database", err.Error())

}
