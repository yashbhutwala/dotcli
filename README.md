<!-- omit in toc -->
# dotcli

CLI to query simple graph data written in [**Graphviz DOT**](https://en.wikipedia.org/wiki/Graphviz) language.  [Graphviz](https://gitlab.com/graphviz/graphviz) DOT [language](https://graphviz.org/) is a popular layman way of quickly describing graph data in textual form and visualizing it.  However, once you have such a DOT file, how do you query this graph data using classic graph algorithms?  This is the problem this tool intends to solve.  One could store the data in a graph database like [neo4j](https://github.com/neo4j/neo4j) or [dgraph](https://github.com/dgraph-io/dgraph), but often dot language is a good least common denominator between devs, qa, devops and non-technical people.  Hence, it makes sense to use the dot file as a "database" and use it to drive multiple downstream consumers based on the ability to query it.

<!-- omit in toc -->
## Table of Contents

- [Installation](#installation)
  - [From source](#from-source)
    - [Build](#build)
    - [Run](#run)
- [Usage](#usage)
  - [`dotcli nodes <PATH_TO_DOT_FILE>`: print the set of all nodes](#dotcli-nodes-path_to_dot_file-print-the-set-of-all-nodes)
  - [`dotcli deps <PATH_TO_DOT_FILE> <NODE_NAME>`: print set of all dependencies of the specified node](#dotcli-deps-path_to_dot_file-node_name-print-set-of-all-dependencies-of-the-specified-node)
    - [Both Direct and Transitive](#both-direct-and-transitive)
    - [Direct only](#direct-only)
- [Links](#links)
  - [Visualization](#visualization)
  - [Relevant Golang libraries](#relevant-golang-libraries)
  - [Other interesting projects](#other-interesting-projects)
- [Future](#future)

## Installation

### From source

#### Build

```bash
go build ./cmd/dotcli/main.go

./dotcli --help
```

#### Run

```bash
go run ./cmd/dotcli/main.go --help
```

## Usage

Use the `--help` flag with the root and/or any of the subcommands to find out more information

```bash
dotcli --help
```

### `dotcli nodes <PATH_TO_DOT_FILE>`: print the set of all nodes

```bash
dotcli nodes --help

PATH_TO_DOT_FILE="TODO"
dotcli nodes $PATH_TO_DOT_FILE
```

### `dotcli deps <PATH_TO_DOT_FILE> <NODE_NAME>`: print set of all dependencies of the specified node

#### Both Direct and Transitive

```bash
dotcli deps --help

PATH_TO_DOT_FILE="TODO"
NODE_NAME="TODO"
dotcli deps $PATH_TO_DOT_FILE $NODE_NAME
```

#### Direct only

```bash
dotcli deps --help

PATH_TO_DOT_FILE="TODO"
NODE_NAME="TODO"
dotcli deps $PATH_TO_DOT_FILE $NODE_NAME --direct-only
```

## Links

### Visualization

[dreampuf GraphvizOnline](https://dreampuf.github.io/GraphvizOnline)
[WebGraphviz](http://www.webgraphviz.com/)
[Viz.js](http://viz-js.com/)

### Relevant Golang libraries

[awalterschulze/gographviz](https://github.com/awalterschulze/gographviz) <- this repo uses this library
[goccy/go-graphviz](https://github.com/goccy/go-graphviz) <- good library, but [requires cgo](https://github.com/goccy/go-graphviz/issues/28)
[emicklei/dot](https://github.com/emicklei/dot) <- for writing dot
[tmc/dot](https://github.com/tmc/dot) <- for writing dot
[gonum/gonum](https://github.com/gonum/gonum/tree/master/graph/topo) <- graph algo library
[golang/tools/cmd/digraph](https://github.com/golang/tools/blob/gopls/v0.4.4/cmd/digraph/digraph.go) <- only supports text; does not support dot language
[yashbhutwala/go-directed-acyclic-graph](https://github.com/yashbhutwala/go-directed-acyclic-graph)
[yashbhutwala/go-scheduler](https://github.com/yashbhutwala/go-scheduler)

### Other interesting projects

[ofabry/go-callvis](https://github.com/ofabry/go-callvis)
[cycloidio/inframap](https://github.com/cycloidio/inframap)
[jpreese/kustomize-graph](https://github.com/jpreese/kustomize-graph)
[neo4j/neo4j](https://github.com/neo4j/neo4j)
[dgraph-io/dgraph](https://github.com/dgraph-io/dgraph)

## Future

- parse dot language into [neo4j database](https://neo4j.com/developer/go/) compatible [cypher query language](https://neo4j.com/developer/cypher/) and/or [GraphQL query syntax](https://dgraph.io/docs/query-language/graphql-fundamentals/)
