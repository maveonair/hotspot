package hotspot

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type RepositoryFile struct {
	FilePath  string `json:"file_path"`
	Revisions int    `json:"revisions"`
	Code      int    `json:"code"`
}

type RepositoryFiles []RepositoryFile

func (r RepositoryFiles) Len() int {
	return len(r)
}

func (r RepositoryFiles) Less(i, j int) bool {
	return r[i].Revisions < r[j].Revisions
}

func (r RepositoryFiles) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RepositoryFiles) ToCSV() string {
	csv := []string{}
	for _, repositoryFile := range r {
		line := fmt.Sprintf("%s,%d,%d", repositoryFile.FilePath, repositoryFile.Revisions, repositoryFile.Code)
		csv = append(csv, line)
	}

	return strings.Join(csv, "\n")
}

func (p RepositoryFiles) ToJson() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshelling repository files", err)
		return "{}"
	}

	return string(bytes)
}
