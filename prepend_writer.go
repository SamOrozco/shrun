package main

import (
	"fmt"
	"io"
	"strings"
)

type PrependWriter struct {
	Prefix            string
	PassThroughWriter io.Writer
}

func (p2 PrependWriter) Write(p []byte) (n int, err error) {
	writtenString := string(p)
	lines := strings.Split(writtenString, "\n")

	// convert back to string
	var bldr strings.Builder

	lineLength := len(lines)
	for i := range lines {
		currentLine := lines[i]
		if strings.TrimSpace(currentLine) == "" {
			continue
		}

		bldr.WriteString(fmt.Sprintf("%s", p2.Prefix))
		bldr.WriteString(currentLine)
		if i != lineLength-1 {
			bldr.WriteString("\n")
		}
	}

	// write to passthrough
	bldrString := bldr.String()
	if n, err = p2.PassThroughWriter.Write([]byte(bldrString)); err != nil {
		return 0, err
	}
	return len(p), nil
}
