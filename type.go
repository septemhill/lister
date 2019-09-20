package main

type FileType int

const (
	TYPE_FILE FileType = iota
	TYPE_DICT
)

const (
	FileTag      = "[F]"
	DirectoryTag = "[D]"
)
