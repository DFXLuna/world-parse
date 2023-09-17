package worldparse

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	bbCodeReference = regexp.MustCompile(`@\[([^\]]+)\]\([^)]+\)`)
	TargetDirs      = []string{"World Encyclopedia", "World Atlas", "Uncategorized Articles"}
	articles        = "articles"
)

type Content struct {
	Title    string       `json:"title"`
	Type     string       `json:"entityClass"`
	Category CategoryType `json:"category"`
	Content  string       `json:"content"`
}

type CategoryType struct {
	Title string `json:"title"`
}

func ParseOneFile(path string) (*Content, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Content
	err = json.Unmarshal(bs, &c)
	if err != nil {
		return nil, err
	}
	c.Content = stripBBCode(c.Content)
	return &c, nil
}

func stripBBCode(data string) string {
	return bbCodeReference.ReplaceAllString(data, "$1")
}

func ParseWorldDirIntoLorebook(root string, outputPath string) error {
	l, err := walkWorld(root)
	if err != nil {
		return err
	}
	outputPath, err = filepath.Abs(outputPath)
	if err != nil {
		return err
	}
	return SerializeLorebook(l, outputPath)
}

func walkWorld(root string) (*Lorebook, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	root = filepath.Join(root, articles)
	cs := []*Content{}
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		isTargetDir := false
		for _, d := range TargetDirs {
			if strings.Contains(path, d) {
				isTargetDir = true
			}
		}
		if !isTargetDir {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		if strings.Contains(path, "-metadata.json") {
			return nil
		}
		fmt.Println("Parsing ", path)
		c, pErr := ParseOneFile(path)
		if pErr != nil {
			return pErr
		}
		cs = append(cs, c)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return NewLorebook(cs), nil
}
