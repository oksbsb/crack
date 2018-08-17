package util

import (
	"bufio"
	"os"
	"strings"
)

func Makelist(f string) ([]string, error) {
	var filelist = make([]string, 0)

	_, err := os.Stat(f)
	if err != nil {
		return append(filelist, f), nil
	}

	return makefilelist(f)
}

func makefilelist(f string) ([]string, error) {
	var filelist = make([]string, 0)
	file, err := os.Open(f)
	if err != nil {
		return filelist, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			filelist = append(filelist, line)
		}
	}
	return filelist, nil
}
