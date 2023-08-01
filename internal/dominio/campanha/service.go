package campanha

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internalErros"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(novaCampanha contract.NovaCampanha) (string, error) {

	campanha, err := NovaCampanha(novaCampanha.Nome, novaCampanha.Conteudo, novaCampanha.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campanha)
	if err != nil {
		return "", internalerros.ErrInternal
	}

	return campanha.ID, nil
}
