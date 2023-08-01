package campanha

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	nome     = "X"
	conteudo = "body Hi"
	contatos = []string{"email1@g.com", "email2@g.com"}
	fake     = faker.New()
)

func Test_NovaCampanha_CriarCampanha(t *testing.T) {
	assert := assert.New(t)

	campanha, _ := NovaCampanha(nome, conteudo, contatos)

	assert.Equal(campanha.Nome, nome)
	assert.Equal(campanha.Conteudo, conteudo)
	assert.Equal(len(campanha.Contatos), len(contatos))

}

func Test_NovaCampanha_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campanha, _ := NovaCampanha(nome, conteudo, contatos)

	assert.NotNil(campanha.ID)
}

func Test_NovaCampanha_CriadoEmMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campanha, _ := NovaCampanha(nome, conteudo, contatos)

	assert.Greater(campanha.CriadoEm, now)
}

func Test_NovaCampanha_MustValidateNomeMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NovaCampanha("", conteudo, contatos)

	assert.Equal("nome is required with min 5", err.Error())
}
func Test_NovaCampanha_MustValidateNomeMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NovaCampanha(fake.Lorem().Text(30), conteudo, contatos)

	assert.Equal("nome is required with max 24", err.Error())
}

func Test_NovaCampanha_MustValidateConteudoMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NovaCampanha(nome, "", contatos)

	assert.Equal("conteudo is required with min 5", err.Error())
}
func Test_NovaCampanha_MustValidateConteudoMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NovaCampanha(nome, "", contatos)

	assert.Equal("conteudo is required with min 5", err.Error())
}

func Test_NovaCampanha_MustValidateContatosMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NovaCampanha(nome, "conteudo", nil)

	assert.Equal("contatos is required with min 1", err.Error())
}
