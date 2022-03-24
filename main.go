package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

var localFileFlag = flag.String("file", "", "local file name")
var localPathFlag = flag.String("path", "", "local file path")
var showHelpFlag = flag.Bool("help", false, "help")
var languageFlag = flag.String("lang", "", "go,ts")
var outputFlag = flag.String("output", "", "save path")

var languages []string
var ext = ".variant"

func main() {
	flag.Parse()
	if *showHelpFlag || (*localFileFlag == "" && *localPathFlag == "") || *languageFlag == "" || *outputFlag == "" {
		flag.Usage()
	} else {
		languages = strings.Split(*languageFlag, ",")
		fmt.Println("convert language", languages)
		if *localFileFlag != "" {
			ConvertFile(*localFileFlag)
		}
	}
}

func ConvertFile(filePath string) {
	if ext != path.Ext(filePath) {
		return
	}
	var lines []string
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err == nil {
		buf := bufio.NewReader(file)
		for {
			line := ""
			line, err = buf.ReadString('\n')
			if (err != nil && err != io.EOF) || (err == io.EOF && line == "") {
				if err == io.EOF {
					err = nil
				}
				break
			}
			lines = append(lines, strings.TrimSpace(line))
		}
		if err == nil {
			fileName := strings.TrimRight(path.Base(filePath), ext)
			for _, language := range languages {
				switch strings.ToUpper(language) {
				case "GO":
					convertGO := new(Convert)
					convertGO.ConvertIDL(lines, fileName, *outputFlag)
				}
			}

		}
	}
	if err != nil {
		fmt.Println("convert faild " + err.Error())
	}
}
