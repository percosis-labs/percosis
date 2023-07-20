package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/percosis-labs/percosis/v16/app"
)

const (
	EnvVariable = "PERCOSISD_ENVIRONMENT"
	EnvMainnet  = "mainnet"
	EnvLocalnet = "localpercosis"
)

// ExportAirdropSnapshotCmd generates a snapshot.json from a provided exported genesis.json.
func ChangeEnvironmentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-env [new env]",
		Short: "Set home environment variables for commands",
		Long: `Set home environment variables for commands
Example:
	percosisd set-env mainnet
	percosisd set-env localpercosis
	percosisd set-env $HOME/.custom-dir
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			newEnv := args[0]

			currentEnvironment := getHomeEnvironment()
			fmt.Println("Current environment: ", currentEnvironment)

			if _, err := environmentNameToPath(newEnv); err != nil {
				return err
			}

			fmt.Println("New environment: ", newEnv)

			envMap := make(map[string]string)
			envMap[EnvVariable] = newEnv
			err := godotenv.Write(envMap, filepath.Join(app.DefaultNodeHome, ".env"))
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

// PrintEnvironmentCmd prints the current environment.
func PrintEnvironmentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-env",
		Short: "Prints the current environment",
		Long: `Prints the current environment
Example:
	percosisd get-env'

	Returns one of:
	- mainnet implying $HOME/.percosisd
	- localpercosis implying $HOME/.percosisd-local
	- localpercosis
	- custom path`,
		RunE: func(cmd *cobra.Command, args []string) error {
			environment := getHomeEnvironment()
			path, err := environmentNameToPath(environment)
			if err != nil {
				return err
			}

			fmt.Println("Environment name: ", environment)
			fmt.Println("Environment path: ", path)
			return nil
		},
	}
	return cmd
}

func environmentNameToPath(environmentName string) (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch environmentName {
	case EnvMainnet:
		return app.DefaultNodeHome, nil
	case EnvLocalnet:
		return filepath.Join(userHomeDir, ".percosisd-local/"), nil
	default:
		percosisdPath := filepath.Join(userHomeDir, environmentName)
		_, err := os.Stat(percosisdPath)
		if os.IsNotExist(err) {
			// Creating new environment directory
			if err := os.Mkdir(percosisdPath, os.ModePerm); err != nil {
				return "", err
			}
		}
		return percosisdPath, nil
	}
}
