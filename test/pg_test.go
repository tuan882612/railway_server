package test

import (
	"main/tools/config"
	"main/tools/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectionToPostgres(t *testing.T) {
	config.LoadEnv("../.env")
	assert.Equal(t, true, validators.ValidatePostgres())
}