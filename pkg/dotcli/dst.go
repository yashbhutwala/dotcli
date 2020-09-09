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

type DstFlags struct {
	DirectOnly bool
}

func SetupDstCommand() *cobra.Command {
	dstFlags := &DstFlags{}
	command := &cobra.Command{
		Use:   "dst PATH_TO_DOT_FILE NODE_NAME",
		Short: "the set of all 'destinations' (direct and transitive) of the specified nodes",
		Long:  "the set of all 'destinations' (direct and transitive) of the specified nodes",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			utils.DoOrDie(RunDstCommand(args[0], args[1], dstFlags))
		},
	}

	command.Flags().BoolVarP(&dstFlags.DirectOnly, "direct-only", "d", false, "if enabled, only print the direct dependencies")
	return command
}

func RunDstCommand(filePath string, nodeName string, dstFlags *DstFlags) error {
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

	if dstFlags.DirectOnly {
		log.Infof("calculating only direct 'destinations' of: '%s'", nodeName)
		fmt.Println(GetOnlyDirectDstNodes(graph, nodeName))
	} else {
		log.Infof("calculating all 'destinations' of: %s", nodeName)
		fmt.Println(GetAllDstNodes(graph, nodeName))
	}
	return nil
}

func GetAllDstNodes(graph *gographviz.Graph, nodeName string) []string {
	var allDstNodes []string
	allDirectDstNodes := GetOnlyDirectDstNodes(graph, nodeName)
	if 0 == len(allDirectDstNodes) {
		return allDstNodes
	} else {
		allDstNodes = append(allDstNodes, allDirectDstNodes...)
		for _, directDstNode := range allDirectDstNodes {
			allDstNodes = append(allDstNodes, GetAllDstNodes(graph, directDstNode)...)
		}
	}
	return allDstNodes
}

func GetOnlyDirectDstNodes(graph *gographviz.Graph, nodeName string) []string {
	srcToDstForNode := graph.Edges.SrcToDsts[nodeName]
	if 0 == len(srcToDstForNode) {
		return []string{}
	} else {
		dstList := make([]string, 0, len(srcToDstForNode))
		for dst := range srcToDstForNode {
			dstList = append(dstList, dst)
		}
		return dstList
	}
}
