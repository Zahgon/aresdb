package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uber/aresdb/common"
)

// AddFlags adds flags to command
func AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// ReadConfig populate AresServerConfig
func ReadConfig(defaultCfg map[string]interface{}, flags *pflag.FlagSet) (common.AresServerConfig, error) {
	_ = "STUB: not implemented"
	return *new(common.AresServerConfig), nil
}

// bind command flags

// set defaults

// merge in config file
