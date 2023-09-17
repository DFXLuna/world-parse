package worldparse

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Lorebook struct {
	Entries map[string]Entry `json:"entries"`
}

type Entry struct {
	UID              int      `json:"uid"`
	Key              []string `json:"key"`
	Keysecondary     []string `json:"keysecondary"`
	Comment          string   `json:"comment"`
	Content          string   `json:"content"`
	Constant         bool     `json:"constant"`
	Selective        bool     `json:"selective"`
	SelectiveLogic   int      `json:"selectiveLogic"`
	AddMemo          bool     `json:"addMemo"`
	Order            int      `json:"order"`
	Position         int      `json:"position"`
	Disable          bool     `json:"disable"`
	ExcludeRecursion bool     `json:"excludeRecursion"`
	Probability      int      `json:"probability"`
	UseProbability   bool     `json:"useProbability"`
	DisplayIndex     int      `json:"displayIndex"`
}

func NewEntry(c *Content) *Entry {
	return &Entry{
		Key:          []string{c.Title},
		Content:      c.Content,
		Probability:  100,
		Keysecondary: []string{},
		AddMemo:      true,
		Position:     0,
	}
}

func NewLorebook(cs []*Content) *Lorebook {
	l := Lorebook{
		Entries: map[string]Entry{},
	}
	for i, c := range cs {
		entry := NewEntry(c)
		entry.UID = i
		entry.DisplayIndex = i
		entry.Order = 100
		l.Entries[fmt.Sprint(i)] = *entry
	}
	return &l
}

func SerializeLorebook(l *Lorebook, path string) error {
	bs, err := json.Marshal(l)
	if err != nil {
		return err
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bs, os.ModePerm)
}
