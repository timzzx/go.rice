// Package embedded defines embedded data types that are shared between the go.rice package and generated code.
package embedded

import (
	"fmt"
	"time"
)

// EmbeddedBox defines an embedded box
type EmbeddedBox struct {
	Name  string                   // box name
	Time  time.Time                // embed time
	Files map[string]*EmbeddedFile // ALL embedded files by full path
	Dirs  map[string]*EmbeddedDir  // ALL embedded dirs by full path
}

// EmbeddedSingle defines an embedded single
type EmbeddedSingle struct {
	Name string        // single name
	Time time.Time     // embed time
	File *EmbeddedFile // embedded file
}

// EmbeddedDir is instanced in the code generated by the rice tool and contains all necicary information about an embedded file
type EmbeddedDir struct {
	Filename   string
	DirModTime time.Time
	ChildDirs  []*EmbeddedDir  // direct childs, as returned by virtualDir.Readdir()
	ChildFiles []*EmbeddedFile // direct childs, as returned by virtualDir.Readdir()
}

// EmbeddedFile is instanced in the code generated by the rice tool and contains all necicary information about an embedded file
type EmbeddedFile struct {
	Filename    string // filename
	FileModTime time.Time
	Content     string
}

// EmbeddedBoxes is a public register of embedded boxes
var EmbeddedBoxes = make(map[string]*EmbeddedBox)

// RegisterEmbeddedBox registers an EmbeddedBox
func RegisterEmbeddedBox(name string, box *EmbeddedBox) {
	if _, exists := EmbeddedBoxes[name]; exists {
		panic(fmt.Sprintf("EmbeddedBox with name `%s` exists already", name))
	}
	EmbeddedBoxes[name] = box
}
