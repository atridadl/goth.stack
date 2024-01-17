package lib

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

func ExtractFrontMatter(file os.DirEntry, dir string) (CardLink, error) {
	f, err := os.Open(dir + file.Name())
	if err != nil {
		return CardLink{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return CardLink{}, fmt.Errorf("failed to read file: %w", err)
	}

	content := strings.Join(lines, "\n")
	splitContent := strings.SplitN(content, "---", 3)
	if len(splitContent) < 3 {
		return CardLink{}, fmt.Errorf("invalid file format: %s", file.Name())
	}

	frontMatter := CardLink{}
	if err := yaml.Unmarshal([]byte(splitContent[1]), &frontMatter); err != nil {
		return CardLink{}, fmt.Errorf("failed to unmarshal frontmatter: %w", err)
	}

	md := goldmark.New(goldmark.WithExtensions())
	var buf bytes.Buffer
	if err := md.Convert([]byte(splitContent[2]), &buf); err != nil {
		return CardLink{}, fmt.Errorf("failed to convert markdown: %w", err)
	}

	return frontMatter, nil
}

func SplitFrontmatter(md []byte) (frontmatter []byte, content []byte, err error) {
	parts := bytes.SplitN(md, []byte("---"), 3)

	if len(parts) < 3 {
		return nil, nil, errors.New("invalid or missing frontmatter")
	}

	return parts[1], parts[2], nil
}
