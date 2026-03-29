package code

import (
	"code/internal/diff"
	"code/internal/formatter"
	"code/internal/parser"
)

func Gendiff(format, filepath1, filepath2 string) (string, error) {
	fileData1, err := parser.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	fileData2, err := parser.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	diffTree := diff.BuildDiffTree(fileData1, fileData2)
	diff, err := formatter.Format(format, diffTree)
	if err != nil {
		return "", err
	}

	return diff, nil
}
