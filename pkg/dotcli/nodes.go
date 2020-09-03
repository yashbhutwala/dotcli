package dotcli

import (
	"fmt"
	"io/ioutil"

	"github.com/awalterschulze/gographviz"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/yashbhutwala/dotcli/pkg/utils"
)

func SetupNodesCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "nodes",
		Short: "the set of all nodes",
		Long:  "the set of all nodes",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			utils.DoOrDie(RunNodesCommand(args[0]))
		},
	}
	return command
}

func RunNodesCommand(filePath string) error {
	log.Tracef("reading filePath: %s", filePath)
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	log.Infof("parsing file contents into a graph")
	graph, err := gographviz.Read(fileContents)
	if err != nil {
		return err
	}

	for _, node := range graph.Nodes.Nodes {
		fmt.Println(node.Name)
	}

	return nil
}
