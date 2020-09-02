package dotcli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func SetupNodesCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "nodes",
		Short: "the set of all nodes",
		Long:  "the set of all nodes",
		Run: func(cmd *cobra.Command, args []string) {
			RunNodesCommand()
		},
	}
	return command
}

func RunNodesCommand () {
	log.Tracef("inside of run nodes command")
}