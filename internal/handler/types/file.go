package types

import "io"

type File struct {
	FileType string
	FileData io.Reader
}
