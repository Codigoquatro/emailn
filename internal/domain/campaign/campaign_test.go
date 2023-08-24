package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	//Arrange (Preparação): Configura o cenário de teste.
	name := "Campaign X"
	content := "Body"
	contacts := []string{"alves@gmail.com", "teste@gmail.com"}

	//Act (Execução): Chama o método ou função sob teste.
	campaign := NewCampaign(name, content, contacts)

	//Assert (Verificação): Verifica se o resultado é o esperado.
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
