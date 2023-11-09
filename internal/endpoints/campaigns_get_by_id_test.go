package endpoints

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_Should_Return_Campaign_Campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "343",
		Name:    "teste",
		Content: "Hi!",
		Status:  "Pending",
	}
	service := new(internalmock.CampaingServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)

}
func Test_CampaignsGetById_Should_Return_Error_When_Something_Wrong(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaingServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())

}
