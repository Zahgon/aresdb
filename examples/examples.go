package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	schemaDir  = "schema"
	queriesDir = "queries"
	dataDir    = "data"

	dataSetKeyName = "dataset"
	hostKeyName    = "host"
	portKeyName    = "port"
)

func queryDataSet() { _ = "STUB: not implemented"; return }

func createTablesForDataSet() { _ = "STUB: not implemented"; return }

func ingestDataForDataSet() { _ = "STUB: not implemented"; return }

func main() {
	rootCmd := &cobra.Command{
		Use:     "examples",
		Short:   "AresDB Examples",
		Long:    `AresDB Examples Contains examples for interact with aresdb`,
		Example: `./examples help tables`,
	}
	rootCmd.PersistentFlags().String(dataSetKeyName, "1k_trips", "name for data set")
	rootCmd.PersistentFlags().String(hostKeyName, "localhost", "host of aresdb server")
	rootCmd.PersistentFlags().Int(portKeyName, 9374, "port of aresdb server")
	viper.SetDefault(dataSetKeyName, "1k_trips")
	viper.SetDefault(hostKeyName, "localhost")
	viper.SetDefault(portKeyName, 9374)
	viper.BindPFlags(rootCmd.PersistentFlags())

	dataCmd := &cobra.Command{
		Use:     "data",
		Short:   "Ingest data for example dataset",
		Long:    `Ingest data for example dataset`,
		Example: `./examples data --dataset 1k_trips`,
		Run: func(cmd *cobra.Command, args []string) {
			ingestDataForDataSet()
		},
	}

	tableCmd := &cobra.Command{
		Use:     "tables",
		Short:   `Create tables for example dataset`,
		Long:    `Create tables for example dataset`,
		Example: `./examples tables --dataset 1k_trips`,
		Run: func(cmd *cobra.Command, args []string) {
			createTablesForDataSet()
		},
	}

	queryCmd := &cobra.Command{
		Use:     "query",
		Short:   `Run sample queries against example dataset`,
		Long:    `Run sample queries against example dataset`,
		Example: `./examples query --dataset 1k_trips`,
		Run: func(cmd *cobra.Command, args []string) {
			queryDataSet()
		},
	}

	rootCmd.AddCommand(tableCmd, dataCmd, queryCmd)
	rootCmd.Execute()
}
