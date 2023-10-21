package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	errs "github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Load loads the configuration in filepath.
// You can use this config with functions with config.Web(), etc.
func Load(filePath string) error {
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return errs.Wrap(err, "failed to load config")
	}

	if err := applyEnv(); err != nil {
		return errs.Wrap(err, "apply env failed")
	}

	var config RuntimeConfig
	if err := viper.UnmarshalExact(&config); err != nil {
		return errs.Wrap(err, "config unmarshaling failed")
	}

	conf = config

	return nil
}

// applyEnv applies environment variables to config's value.
// If matching variable is not found or is "", it returns error.
func applyEnv() (err error) {
	for _, key := range viper.AllKeys() {
		val := viper.GetString(key)

		if strings.HasPrefix(val, "${") && strings.HasSuffix(val, "}") {
			k := val
			val = os.ExpandEnv(val)
			if val == "" {
				err = errors.Join(err, fmt.Errorf("%s: env var does not exist", k))
			}

			viper.Set(key, val)
		}
	}
	return
}
