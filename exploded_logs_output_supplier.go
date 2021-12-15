package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var dirLocation = "." + string(filepath.Separator) + "ssh_logs"

func ExplodedLogOutputSupplier(hostname string, command string) io.Writer {
	currentLocation := fmt.Sprintf("%s_%s", dirLocation, time.Now().Format(time.Kitchen))
	if err := os.MkdirAll(currentLocation, os.ModePerm); err != nil {
		panic(err)
	}

	// create file for hostname and command to write to
	fileLocation := currentLocation + string(filepath.Separator) + fmt.Sprintf("%s.log", hostname)
	file, err := os.Create(fileLocation)
	if err != nil {
		panic(err)
	}
	return file
}
