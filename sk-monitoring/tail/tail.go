package tail

import (
	"os"
)

//Get will return a end of line of a file
func Get(filePath string) (string, error) {

	f, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	//Read a character
	b := make([]byte, 1)
	f.Read(b)
	offset := int64(-2)

	for string(b) != "\n" {
		f.Seek(offset, os.SEEK_END)
		f.Read(b)
		offset--
	}

	bOfEndLine := make([]byte, offset*(-1))
	f.Read(bOfEndLine)

	return string(bOfEndLine), nil
}
