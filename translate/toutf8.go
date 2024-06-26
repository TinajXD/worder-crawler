package translate

import(
	"bytes"
	"io"
	"golang.org/x/net/html/charset"
)

func ConvertToUTF8(str string, origEncoding string) string {
	strBytes := []byte(str)
	byteReader := bytes.NewReader(strBytes)
	reader, _ := charset.NewReaderLabel(origEncoding, byteReader)
	strBytes, _ = io.ReadAll(reader)
	return string(strBytes)
}
