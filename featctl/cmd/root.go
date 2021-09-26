package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/onestore-ai/onestore/featctl/pkg/database"
	database2 "github.com/onestore-ai/onestore/pkg/database"
	"github.com/onestore-ai/onestore/version"
)

const (
	envPrefix = "FEATCTL"
)

var cfgFile string
var defaultCfgFile = filepath.Join(xdg.ConfigHome, "featctl", "config.yaml")
var dbOption database.Option
var sqlxDbOption database2.Option

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "featctl",
	Short:   "a cli tool lets you control the feature store.",
	Version: version.String(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		sqlxDbOption = buildSqlxDBOption(dbOption)
	}

	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVar(&cfgFile, "config", defaultCfgFile, "config file")

	pFlags.StringVarP(&dbOption.Host, "host", "H", "127.0.0.1", "feature store database host")
	pFlags.StringVarP(&dbOption.Port, "port", "P", "4000", "feature store database port")
	pFlags.StringVarP(&dbOption.User, "user", "u", "root", "feature store database user")
	pFlags.StringVarP(&dbOption.Pass, "pass", "p", "", "feature store database pass")
	pFlags.StringVarP(&dbOption.DbName, "database", "d", "onestore", "feature store database name")

	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
	bindViperFlags(rootCmd)
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindViperFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			envName := fmt.Sprintf("%s_%s", envPrefix, envVarSuffix)
			if err := viper.BindEnv(f.Name, envName); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			if err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	})
}

func buildSqlxDBOption(option database.Option) database2.Option {
	return database2.Option{
		Host:   option.Host,
		Port:   option.Port,
		User:   option.User,
		Pass:   option.Pass,
		DbName: option.DbName,
	}
}
