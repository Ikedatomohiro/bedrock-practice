package main

import (
	"bytes"
	"fmt"

	"github.com/dslipak/pdf"
)

func main() {
	content, err := readPdf("sample.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
