package common

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// FileReader is a helper function to get an io.Reader given a filename
func FileReader(filename string) io.Reader {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}

// FileRead is a helper function to read a file and return a string
func FileRead(filename string) string {
	reader := FileReader(filename)
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func RootDir() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}
