package campanha

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internalErros"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	novaCampanha = contract.NovaCampanha{
		Nome:     "X",
		Conteudo: "Body",
		Emails:   []string{"email2@g.com", "email3@g.com"},
	}

	service = Service{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campanha *Campanha) error {
	args := r.Called(campanha)
	return args.Error(0)
}

func Test_Criar_Campanha(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(novaCampanha)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Criar_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	novaCampanha.Nome = ""
	_, err := service.Create(novaCampanha)

	assert.NotNil(err)
	assert.Equal("nome is required", err.Error())
}

func Test_Criar_SaveCampanha(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campanha *Campanha) bool {
		if campanha.Nome != novaCampanha.Nome ||
			campanha.Conteudo != novaCampanha.Conteudo ||
			len(campanha.Contatos) != len(novaCampanha.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(novaCampanha)

	repositoryMock.AssertExpectations(t)
}

func Test_Criar_ValidateDataBaseSave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(novaCampanha)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}
