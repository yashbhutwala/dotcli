package dotcli

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func InitAndExecute() {
	rootFlags := &RootFlags{}
	rootCmd := rootFlags.SetupRootCommand()
	if err := errors.Wrapf(rootCmd.Execute(), "run dotcli root command"); err != nil {
		log.Fatalf("unable to run root command: %+v", err)
		os.Exit(1)
	}
}

type RootFlags struct {
	LogLevel string
}

func (r *RootFlags) SetupRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dotcli",
		Short: "TODO",
		Long:  `TODO`,
		Args:  cobra.MaximumNArgs(0),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return r.setUpLogger()
		},
	}

	rootCmd.PersistentFlags().StringVarP(&r.LogLevel, "verbosity", "v", "info", "log level; one of [info, debug, trace, warn, error, fatal, panic]")

	rootCmd.AddCommand(SetupVersionCommand())
	rootCmd.AddCommand(SetupNodesCommand())
	rootCmd.AddCommand(SetupDepsCommand())

	return rootCmd
}

func (r *RootFlags) setUpLogger() error {
	logLevel, err := log.ParseLevel(r.LogLevel)
	if err != nil {
		return errors.Wrapf(err, "unable to parse the specified log level: '%s'", logLevel)
	}
	log.SetLevel(logLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Infof("log level set to '%s'", log.GetLevel())
	return nil
}
