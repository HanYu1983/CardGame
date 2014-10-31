package tool

import (
	"bytes"
	"image"
	"os"
	"encoding/base64"
	"image/png"
	"io"
)

func PngToBase64(img image.Image) []byte{
	buf := &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

func EncodeFileToBase64(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	
	fi, _ := file.Stat()
	size := fi.Size()
	
	data := make([]byte, size)
	file.Read(data)
	return base64.StdEncoding.EncodeToString(data)
}
 
func DecodeBase64ToImage(b64 string) image.Image{
	data, _ := base64.StdEncoding.DecodeString(b64) 
	image, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err.Error())
	}
	return image
}

func WritePng(writer io.Writer, img image.Image){
	err := png.Encode(writer, img)
	if err != nil {
		panic(err.Error())
	}
}