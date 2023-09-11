package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
)

type Service interface {
	Create(newCampaing contract.NewCampaingn) (string, error)
}
type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaingn) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaign.ID, nil
}
