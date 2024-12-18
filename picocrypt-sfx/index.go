package main

import (
	_ "embed"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//go:embed index.html
var template string

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Exactly one .pcv must be specified!")
		os.Exit(1)
	}
	path := flag.Arg(0)
	if len(path) == 0 {
		fmt.Println("Filepath to the .pcv must not be empty!")
		os.Exit(1)
	}
	if !strings.HasSuffix(path, ".pcv") {
		fmt.Println("Filepath must end in .pcv!")
		os.Exit(1)
	}
	if stat, err := os.Stat(path); err != nil || stat.IsDir() {
		fmt.Println("Invalid .pcv file!")
		os.Exit(1)
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to open .pcv!")
		os.Exit(1)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Unable to read file!")
		os.Exit(1)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	filename := filepath.Base(path)

	template = strings.ReplaceAll(template, "<VOLUME_FILENAME>", filename)
	template = strings.ReplaceAll(template, "<VOLUME_AS_BASE64>", encoded)

	fout, err := os.Create(path + ".html")
	defer fout.Close()
	if err != nil {
		fmt.Println("Failed to create output .html!")
		os.Exit(1)
	}
	if _, err := io.WriteString(fout, template); err != nil {
		fmt.Println("Failed writing to output .html!")
		os.Exit(1)
	}

	fmt.Println(filename + ".html" + " created successfully! Please test it to ensure it works.")
}
