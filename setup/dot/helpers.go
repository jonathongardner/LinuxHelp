package dot

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

func isFile(file string) error {
	stats, err := os.Stat(file)
	if err == nil {
		if stats.IsDir() {
			return fmt.Errorf("not a file %v", file)
		}
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file doesnt exist %v", file)
	} else {
		return fmt.Errorf("error checking if a directory %v (%v)", file, err)
	}
}

func isDir(file string) error {
	stats, err := os.Stat(file)
	if err == nil {
		if stats.Mode().IsRegular() {
			return fmt.Errorf("not a directory %v", file)
		}
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("directory doesnt exist %v", file)
	} else {
		return fmt.Errorf("error checking if a directory %v (%v)", file, err)
	}
}

func proceed(question string) bool {
	fmt.Printf("%v (Y/y)?\n", question)
	var res string
	fmt.Scanln(&res)
	return strings.EqualFold(res, "yes") || strings.EqualFold(res, "y")
}

func createOrAddToFile(fileToAddTo, line string) error {
	file, err := os.OpenFile(fileToAddTo, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return writeToFileIfNotThere(file, line)
}

func addToFile(fileToAddTo, line string) error {
	if err := isFile(fileToAddTo); err != nil {
		return err
	}

	file, err := os.OpenFile(fileToAddTo, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file %v (%v)", fileToAddTo, err)
	}
	defer file.Close()
	return writeToFileIfNotThere(file, line)
}

func writeToFileIfNotThere(f *os.File, line string) error {
	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(f)

	if strings.Contains(buf.String(), line) {
		fmt.Printf("Skipping %v... already included\n", f.Name())
		return nil
	}

	for l, s := range []string{line, "\n"} {
		_, err := f.WriteString(s)
		if err != nil {
			return fmt.Errorf("error wirting %v (%v)", l, err)
		}
	}

	return nil
}
