package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	paths, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, path := range paths {
		if absPath, err := filepath.Abs(path); err != nil {
			fmt.Printf("Failed to get absolute path of \"%s\": %v", path, err)
		} else {
			fmt.Println(absPath)
		}
	}
}

// ParseArgs parses command line arguments.
func ParseArgs() (paths []string, err error) {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse stdin: %v", err)
	}
	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		defer os.Stdin.Close()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			paths = append(paths, scanner.Text())
		}
		return paths, nil
	}

	if len(os.Args) < 2 {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		paths = append(paths, pwd)
		return paths, nil
	}

	for _, s := range os.Args[1:] {
		paths = append(paths, s)
	}
	return paths, nil
}
