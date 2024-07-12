package auth

import	(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	email := "xyz@mail.com"
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBtYWlsLmNvbSJ9.Ee7O9cwe7wBFizHY1hvAN0wJBj9PH2m6MIGx6trsncQ"
	actual, _ := GenerateJWT(email)

	assert.Equal(t, expected, actual)
}