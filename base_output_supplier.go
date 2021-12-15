package main

import (
	"fmt"
	"io"
	"os"
)

func getOutputSupplier(shouldRemovePrefix bool) OutputSupplier {
	if shouldRemovePrefix {
		return NoPrefixOutputSupplier
	} else {
		return BaseOutputSupplier
	}
}

func NoPrefixOutputSupplier(hostname, command string) io.Writer {
	return os.Stdout
}

func BaseOutputSupplier(hostname, command string) io.Writer {
	return &PrependWriter{
		Prefix:            fmt.Sprintf("[%s](%s) - ", hostname, command),
		PassThroughWriter: os.Stdout,
	}
}
