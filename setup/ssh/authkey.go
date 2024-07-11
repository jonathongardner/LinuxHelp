package ssh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type AuthKey struct {
	file   *os.File
	dryRun bool
	keys   map[string]bool
}
type key interface {
	GetKey() string
	// Fingerprint() string
	String() string
}

func NewAuthKey(path string, dryRun bool) (*AuthKey, error) {
	// BUG: dry run true will still create directories/files should fix
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, fmt.Errorf("error creating authorized key folders: %v (%v)", path, err)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("error opening authorized key file: %v (%v)", path, err)
	}

	// Read authkeys
	return &AuthKey{f, dryRun, make(map[string]bool)}, nil
}

func (au *AuthKey) Clear() error {
	if au.dryRun {
		return nil
	}
	return au.file.Truncate(0)
}

func (au *AuthKey) Parse() error {
	_, err := au.file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file to parse start %v", err)
	}

	scanner := bufio.NewScanner(au.file)
	for scanner.Scan() {
		key := strings.SplitN(scanner.Text(), " ", 3)
		if len(key) >= 2 {
			au.keys[strings.Join(key[0:2], " ")] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file to parse %v", err)
	}

	_, err = au.file.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("error seeking file to parse end %v", err)
	}

	return nil
}

// return false if already saved
func (au *AuthKey) Add(key key) (bool, error) {
	if au.keys[key.GetKey()] {
		return false, nil
	}

	if !au.dryRun {
		_, err := au.file.WriteString(key.String())
		if err != nil {
			return false, fmt.Errorf("error adding key %v", err)
		}
		_, err = au.file.WriteString("\n")
		if err != nil {
			return false, fmt.Errorf("error adding newline key %v", err)
		}
	}

	au.keys[key.GetKey()] = true
	return true, nil
}

func (au *AuthKey) Close() error {
	return au.file.Close()
}
