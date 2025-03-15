package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type SearchServiceInterface interface {
	SearchFiles(pattern, directory string) ([]string, error)
	CountMatches(pattern, filename string) (string, error)
	FindFiles(glob, directory string, days int) error
}

var _ SearchServiceInterface = &SearchService{}

type SearchService struct {
	CaseSensitive bool
}

func NewSearch(caseSensitive bool) *SearchService {
	return &SearchService{CaseSensitive: caseSensitive}
}

func (s *SearchService) SearchFiles(pattern, directory string) ([]string, error) {
	reFlags := regexp.MustCompile("")
	if !s.CaseSensitive {
		reFlags = regexp.MustCompile(`(?i)`)
	}
	re := regexp.MustCompile(reFlags.String() + regexp.QuoteMeta(pattern))
	result := []string{}
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, _ := ioutil.ReadFile(path)
			if re.Match(content) {
				result = append(result, fmt.Sprintf("Match found in: %s\n", path))
			}
		}
		return nil
	})
	return result, nil
}

func (s *SearchService) CountMatches(pattern, filename string) (string, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	re := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(pattern))
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			count++
		}
	}
	return fmt.Sprintf("Found %d matches in %s\n", count, filename), nil
}

func (s *SearchService) FindFiles(glob, directory string, days int) error {
	cutoffTime := time.Now().AddDate(0, 0, -days)
	pattern := filepath.Join(directory, glob)

	matches, _ := filepath.Glob(pattern)
	for _, file := range matches {
		info, _ := os.Stat(file)
		if info.ModTime().Before(cutoffTime) {
			fmt.Printf("Found old file: %s (%s)\n", file, info.ModTime().Format("2006-01-02"))
		}
	}
	return nil
}
