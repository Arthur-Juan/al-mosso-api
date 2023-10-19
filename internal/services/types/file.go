package types

import "io"

type TFile struct {
	FileType  string
	FileData  io.Reader
	Extension string
	FileName  string
}
