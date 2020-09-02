package dotcli

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func InitAndExecute() {
	rootCmd := SetupRootCommand()
	if err := errors.Wrapf(rootCmd.Execute(), "run dotcli root command"); err != nil {
		log.Fatalf("unable to run root command: %+v", err)
		os.Exit(1)
	}
}

func SetupRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dotcli",
		Short: "TODO",
		Long:  `TODO`,
		Args:  cobra.MaximumNArgs(0),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return rootCmd
}
