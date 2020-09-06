package dotcli

import (
	"fmt"
	"io/ioutil"

	"github.com/awalterschulze/gographviz"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/yashbhutwala/dotcli/pkg/utils"
)

type DepFlags struct {
	DirectOnly bool
}

func SetupDepsCommand() *cobra.Command {
	depFlags := &DepFlags{}
	command := &cobra.Command{
		Use:   "deps",
		Short: "the set of all dependencies (direct and transitive) of the specified nodes",
		Long:  "the set of all dependencies (direct and transitive) of the specified nodes",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			utils.DoOrDie(RunDepsCommand(args[0], args[1], depFlags))
		},
	}

	command.Flags().BoolVarP(&depFlags.DirectOnly, "direct-only", "d", false, "if enabled, only print the direct dependencies")
	return command
}

func RunDepsCommand(filePath string, nodeName string, depFlags *DepFlags) error {
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

	// check if nodeName even exists
	// TODO: quoted nodeName
	nodeName = fmt.Sprintf("\"%s\"", nodeName)
	if false == graph.IsNode(nodeName) {
		return errors.Errorf("given node name: %s does not exist in the given graph", nodeName)
	}
	log.Debugf("nodeName: %s found in the graph...continuing", nodeName)

	if depFlags.DirectOnly {
		log.Infof("calculating only direct dependencies of: '%s'", nodeName)
		fmt.Println(GetDirectDependencies(graph, nodeName))
	} else {
		log.Infof("calculating all dependencies of: %s", nodeName)
		fmt.Println(GetAllDependencies(graph, nodeName))
	}
	return nil
}

func GetAllDependencies(graph *gographviz.Graph, nodeName string) []string {
	var allDeps []string
	directDependencies := GetDirectDependencies(graph, nodeName)
	if 0 == len(directDependencies) {
		return allDeps
	} else {
		allDeps = append(allDeps, directDependencies...)
		for _, directDependency := range directDependencies {
			allDeps = append(allDeps, GetAllDependencies(graph, directDependency)...)
		}
	}
	return allDeps
}

func GetDirectDependencies(graph *gographviz.Graph, nodeName string) []string {
	dstToSrcForNode := graph.Edges.DstToSrcs[nodeName]
	if 0 == len(dstToSrcForNode) {
		return []string{}
	} else {
		srcList := make([]string, 0, len(dstToSrcForNode))
		for src := range dstToSrcForNode {
			srcList = append(srcList, src)
		}
		return srcList
	}
}
