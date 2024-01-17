package lib_test

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
	"goth.stack/lib"
)

func TestExtractFrontMatter(t *testing.T) {
	// Create a temporary file with some front matter
	tmpfile, err := os.CreateTemp("../content", "example.*.md")
	println(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	text := `---
name: "Test Title"
description: "Test Description"
---

# Test Content

`
	if _, err := tmpfile.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// Get the directory entry for the temporary file
	dirEntry, err := os.ReadDir(filepath.Dir(tmpfile.Name()))
	if err != nil {
		log.Fatal(err)
	}

	var tmpFileEntry fs.DirEntry
	for _, entry := range dirEntry {
		if entry.Name() == filepath.Base(tmpfile.Name()) {
			tmpFileEntry = entry
			break
		}
	}

	// Now we can test ExtractFrontMatter
	frontMatter, err := lib.ExtractFrontMatter(tmpFileEntry, "../content/")
	assert.NoError(t, err)
	assert.Equal(t, "Test Title", frontMatter.Name)
	assert.Equal(t, "Test Description", frontMatter.Description)
}
