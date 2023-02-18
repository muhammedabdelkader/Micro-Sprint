package ocrExtracter

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

/*
This bot defines an OCR struct that contains a reference to the OCR client, and methods for initializing and closing the client, as well as recognizing text from an image. The struct is then instantiated and used in the main function to recognize text from an image.
*/
type OCR struct {
	client *gosseract.Client
}

func (o *OCR) Init() {
	o.client = gosseract.NewClient()

}

func (o *OCR) Close() {
	o.client.Close()

}

func (o *OCR) Recognize(imagePath string) (string, error) {
	o.client.SetImage(imagePath)
	text, err := o.client.Text()
	return text, err

}

func main() {
	// create new OCR object
	ocr := OCR{}
	ocr.Init()
	defer ocr.Close()
	// recognize text
	text, err := ocr.Recognize("path/to/image.png")
	if err != nil {
		fmt.Println(err)
		return

	}
	// print the recognized text
	fmt.Println(text)

}
