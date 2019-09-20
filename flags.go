package main

import (
	"math"

	"github.com/urfave/cli"
)

const (
	maxListLimit = math.MaxInt64
)

var DirFlag = cli.StringFlag{
	Name:     "dir",
	Usage:    "start with specified directory",
	Required: true,
}

var FindFlag = cli.StringFlag{
	Name:     "find",
	Usage:    "filter folder/file name with specified characteries",
	Required: false,
}

var caseSensitive = true
var CaseSensitiveFlag = cli.BoolFlag{
	Name:        "case",
	Usage:       "case sensitive or not",
	Required:    false,
	Destination: &caseSensitive,
}

var listLimit int64
var ListLimitFlag = cli.Int64Flag{
	Name:        "count",
	Usage:       "list limit count",
	Required:    false,
	Value:       maxListLimit,
	Destination: &listLimit,
}

var showType = true
var TypeDisplayFlag = cli.BoolFlag{
	Name:        "showtype",
	Usage:       "display type (file/dictionary)",
	Required:    false,
	Destination: &showType,
}
