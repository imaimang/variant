package main

import (
	"os"
	"path"
	"strings"
)

type ConvertGO struct {
}

func (c *ConvertGO) ConvertIDL(lines []string, fileName string, output string) error {
	var err error
	savePath := path.Join(output, "go")
	if _, err = os.Stat(savePath); os.IsNotExist(err) {
		err = os.MkdirAll(savePath, os.ModePerm)
	}
	if err == nil {
		fileSaveName := path.Join(savePath, fileName+".go")
		var file *os.File
		file, err = os.OpenFile(fileSaveName, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err == nil {
			for _, line := range lines {
				if strings.Index(line, "namespace go ") == 0 {
					nameSpace := strings.TrimPrefix(line, "namespace go ")
					file.WriteString("package ")
					file.WriteString(nameSpace)
					file.WriteString("\r\n")
				} else if strings.Index(line, "class") == 0 {
					class := strings.TrimPrefix(line, "class ")
					class = strings.TrimSuffix(class, "{")
					class = strings.TrimSpace(class)
					file.WriteString("type ")
					file.WriteString(class)
					file.WriteString("{")
					file.WriteString("\r\n")
				}
			}
		}
	}
	return err
}
