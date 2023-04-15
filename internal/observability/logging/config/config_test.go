package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideLogger(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		cfg := &Config{
			Provider: ProviderZerolog,
		}

		l, err := cfg.ProvideLogger()

		assert.NotNil(t, l)
		assert.NoError(t, err)
	})

	T.Run("no provider", func(t *testing.T) {
		t.Parallel()

		cfg := &Config{}
		l, err := cfg.ProvideLogger()

		assert.NotNil(t, l)
		assert.NoError(t, err)
	})
}
