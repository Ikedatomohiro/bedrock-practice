package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func badSample() {
	// pdfcpu
	pdfFilePath := "norin_suisan_sho.pdf"
	// pdfFilePath := "sample.pdf"
	file, err := os.Open(pdfFilePath)
	if err != nil {
		log.Fatalf("Error opening PDF file: %v\n", err)
	}
	defer file.Close()

	conf := &model.Configuration{
		Path:             pdfFilePath,
		CheckFileNameExt: true,
	}
	// PDFコンテキストを読み込む
	ctx, err := api.ReadContext(file, conf)
	if err != nil {
		log.Fatalf("Error reading PDF context: %v\n", err)
	}

	// ページ数を取得
	numPages := ctx.PageCount
	fmt.Println(numPages, ctx)
	// 全ページのテキストを抽出する
	for i := 1; i <= numPages; i++ {
		text, err := extractPageText(ctx, i)
		if err != nil {
			log.Printf("Error extracting text from page %d: %v\n", i, err)
			continue
		}
		fmt.Printf("Text on page %d:\n%s\n", i, text)
	}
	return
}

func extractPageText(ctx *model.Context, pageNr int) (string, error) {
	r, err := pdfcpu.ExtractPageContent(ctx, pageNr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String(), nil
}
