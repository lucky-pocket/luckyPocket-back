package config_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tmp, _ := os.Create(t.TempDir() + "conf.yml")
	defer tmp.Close()

	testcases := []struct {
		desc     string
		filePath string
		env      map[string]any
		config   []byte
		assert   func(t *testing.T, err error)
	}{
		{
			desc:     "success",
			filePath: tmp.Name(),
			env: map[string]any{
				"HTTP_PORT": 8080,
			},
			config: []byte(`
            web:
              http:
                port: ${HTTP_PORT}
            `),
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc:     "config not found",
			filePath: tmp.Name() + "hi?",
			env:      nil,
			config:   nil,
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "load config")
			},
		},
		{
			desc:     "environment variable not found",
			filePath: tmp.Name(),
			env:      nil,
			config: []byte(`
            web:
              http:
                port: ${HTTP_PORT}
            `),
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "env var")
			},
		},
		{
			desc:     "invalid value",
			filePath: tmp.Name(),
			env: map[string]any{
				"HTTP_PORT": "hi?",
			},
			config: []byte(`
            web:
              http:
                port: ${HTTP_PORT}
            `),
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "unmarshaling")
			},
		},
		{
			desc:     "invalid format",
			filePath: tmp.Name(),
			env: map[string]any{
				"HTTP_PORT": 8080,
			},
			config: []byte(`
            web:
              http:
                port: ${HTTP_PORT}
                hi: true
            `),
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "unmarshaling")
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			defer viper.Reset()

			_ = tmp.Truncate(0)
			_, _ = tmp.Seek(0, 0)
			_, _ = tmp.Write(
				bytes.Trim(
					bytes.TrimSpace(tc.config),
					strings.Repeat(" ", 12),
				),
			)

			for k, v := range tc.env {
				t.Setenv(k, fmt.Sprint(v))
			}

			err := config.Load(tc.filePath)

			tc.assert(t, err)
		})
	}
}
