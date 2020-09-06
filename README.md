<!-- omit in toc -->
# dotcli

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
