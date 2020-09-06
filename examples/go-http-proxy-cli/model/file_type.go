package model

type FileType string

const (
	INI FileType = "INI"
	BPP FileType = "BPP"
	CLI FileType = "CLI"
)

type FileSource string

const (
	local  FileSource = "local"
	remote FileSource = "remote"
)
