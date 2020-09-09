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

type SrcFlags struct {
	DirectOnly bool
}

func SetupSrcCommand() *cobra.Command {
	srcFlags := &SrcFlags{}
	command := &cobra.Command{
		Use:   "src PATH_TO_DOT_FILE NODE_NAME",
		Short: "the set of all 'sources' (direct and transitive) of the specified nodes",
		Long:  "the set of all 'sources' (direct and transitive) of the specified nodes",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			utils.DoOrDie(RunDepsCommand(args[0], args[1], srcFlags))
		},
	}

	command.Flags().BoolVarP(&srcFlags.DirectOnly, "direct-only", "d", false, "if enabled, only print the direct dependencies")
	return command
}

func RunDepsCommand(filePath string, nodeName string, srcFlags *SrcFlags) error {
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

	if srcFlags.DirectOnly {
		log.Infof("calculating only direct 'sources' of: '%s'", nodeName)
		fmt.Println(GetOnlyDirectSrcNodes(graph, nodeName))
	} else {
		log.Infof("calculating all 'sources' of: %s", nodeName)
		fmt.Println(GetAllSrcNodes(graph, nodeName))
	}
	return nil
}

func GetAllSrcNodes(graph *gographviz.Graph, nodeName string) []string {
	var allDeps []string
	directDependencies := GetOnlyDirectSrcNodes(graph, nodeName)
	if 0 == len(directDependencies) {
		return allDeps
	} else {
		allDeps = append(allDeps, directDependencies...)
		for _, directDependency := range directDependencies {
			allDeps = append(allDeps, GetAllSrcNodes(graph, directDependency)...)
		}
	}
	return allDeps
}

func GetOnlyDirectSrcNodes(graph *gographviz.Graph, nodeName string) []string {
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
