// JOSEPH S MARTINEZ VER 0.0.1 07-28-2019
// TODOs
// + PATH to access file on other directories
// + JSON Un/Marshaling function
// + Delete file(s)

package datahandler

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultBufSize = 4096
)

// DataHandler provides a convenient interface for reading and writting
// to files.
type DataHandler struct {
	buf int
}

// NewDataHandler returns a new DataHandler read to write stream data
func NewDataHandler() *DataHandler {
	return &DataHandler{
		buf: defaultBufSize,
	}
}

// PrintToFile returns the number of bytes wirtten to a file and
// a bool values upon completion.
func (dh DataHandler) PrintToFile(data *string, filename string) (bool, int) {
	var isDataWritten bool
	f, err := os.Create(filename)
	if err != nil {
		isDataWritten = false
		log.Fatal(err)
	}
	buffSize := (defaultBufSize)
	w := bufio.NewWriterSize(f, buffSize)
	numBytesWritten, err := w.WriteString(*data)
	if err != nil {
		isDataWritten = false
		log.Fatal(err)
	}
	isDataWritten = true
	w.Flush()
	return isDataWritten, numBytesWritten
}

// ReadFile returns a pointer of bytes
func (dh DataHandler) ReadFile(fileName string) []byte {

	pFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pFile.Close()
	byteValue, _ := ioutil.ReadAll(pFile)

	return byteValue
}

// ReadJSONfile reads a .json file and returns a map
func (dh DataHandler) ReadJSONFile(fileName string) []byte {

	pFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pFile.Close()
	byteValue, _ := ioutil.ReadAll(pFile)

	return byteValue
}

// func jsonMarshal([]byte) ([]byte, error)  {

// }

// func jsonUnmarshal([]byte) error {

// }
