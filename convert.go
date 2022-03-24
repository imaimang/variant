package main

import (
	"strings"
)

type IDL struct {
	PropertyType string
	Name         string
	Language     string
}

func ConvertIDL(lines []string, fileName string, output string) error {
	var idls []*IDL
	var err error
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "namespace ") {
			items := strings.Split(line, "")
			if len(items) == 3 {
				idl := new(IDL)
				idl.PropertyType = items[0]
				idl.Language = items[1]
				idl.Name = items[2]
				idls = append(idls, idl)
			}
		} else if strings.HasPrefix(line, "class ") {
			items := strings.Split(line, "")
			if len(items) == 3 {
				idl := new(IDL)
				idl.PropertyType = items[0]
				idl.Language = items[1]
				idl.Name = items[2]
				idls = append(idls, idl)
			}
		}
	}
	return err
}
