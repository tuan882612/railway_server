package test

import (
	"main/test"
	"main/pkg/config"
	"main/pkg/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectionToPostgres(t *testing.T) {
	config.LoadEnv(test.EnvPath)
	assert.Equal(t, true, validators.ValidatePostgres())
}