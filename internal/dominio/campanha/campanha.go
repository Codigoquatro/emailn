package campanha

import (
	internalerros "emailn/internal/internalErros"
	"time"

	"github.com/rs/xid"
)

type Contato struct {
	Email string `validate:"email"`
}

type Campanha struct {
	ID       string    `validate:"required"`
	Nome     string    `validate:"min=5,max=24"`
	CriadoEm time.Time `validate:"required"`
	Conteudo string    `validate:"min=5,max=1024"`
	Contatos []Contato `validate:"min=1,dive"`
}

func NovaCampanha(nome string, conteudo string, emails []string) (*Campanha, error) {

	contatos := make([]Contato, len(emails))
	for index, email := range emails {
		contatos[index].Email = email
	}

	campanha := &Campanha{
		ID:       xid.New().String(),
		Nome:     nome,
		CriadoEm: time.Now(),
		Conteudo: conteudo,
		Contatos: contatos,
	}

	err := internalerros.ValidateStruct(campanha)
	if err == nil {
		return campanha, nil
	}
	return nil, err
}
