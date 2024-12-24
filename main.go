package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8"/>
			<title>Rodix Markdown Preview Tool</title>
		</head>
		<body>
	`
	footer = `
		</body>
	</html>
	`
)

func main() {
	filename := flag.String("file", "", "Markdown file to Preview")
	flag.Usage()

	os.Exit()

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(filename string) error {
	input, err := os.ReadFile(filename)

	if err != nil {
		return err
	}
	htmlData := parserContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)
	return saveHTML(outName, htmlData)
}

func parserContent(input []byte) []byte {
	output := blackfriday.Run(input)
	// valid html block
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// combine the body with the header and the footer.
	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outname string, htmlData []byte) error {
	return os.WriteFile(outname, htmlData, 0644)
}
