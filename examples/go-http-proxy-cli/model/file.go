package model

type File struct {
	fileId     string     ""
	fileType   FileType   "INI | BPP | CLI"
	fileSource FileSource "remote"
	url        string
	fileHash   string
}

func NewFile(fileId string, fileType FileType, fileSource FileSource, url string, fileHash string) *File {
	file := &File{fileId, fileType, fileSource, url, fileHash}

	return file
}
