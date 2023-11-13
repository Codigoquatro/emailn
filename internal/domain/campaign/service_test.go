package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}
func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}
func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaingn{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"alves@gmail.com"},
	}
	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaingn{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))

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
	service.Repository = repositoryMock

	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaing)

	assert.True(errors.Is(internalerrors.ErrInternal, err))

}
func Test_GetById_ReturnCampig(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something wrong"))
	service.Repository = repositoryMock

	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
