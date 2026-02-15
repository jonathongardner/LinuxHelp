package dot

import (
	"fmt"
	"path/filepath"
)

// TODO: file opening/touching could be improved/reduced but its fast enough and run once a blue moon so who cares...
func Export(homedir, project string) error {
	if err := exportBin(homedir); err != nil {
		return err
	}
	if err := sourceBashRC(homedir, project); err != nil {
		return err
	}
	if err := exportGitConfig(homedir, project); err != nil {
		return err
	}
	if err := sourceVimRc(homedir, project); err != nil {
		return err
	}
	return nil
}

func sourceBashRC(homedir, path string) error {
	if !proceed("Source bash rc") {
		return nil
	}

	myBashRC := filepath.Join(path, ".jonathon_bash_profile")
	fmt.Printf("Sourcing %v\n", myBashRC)
	if err := isFile(myBashRC); err != nil {
		return err
	}
	toAdd := fmt.Sprintf("source %v", myBashRC)

	bashRC := filepath.Join(homedir, ".bashrc")
	if err := addToFile(bashRC, toAdd); err != nil {
		return err
	}

	return nil
}

func exportBin(homedir string) error {
	if !proceed("Export bin") {
		return nil
	}

	myBin := filepath.Join(homedir, "bin")
	fmt.Printf("Exporting %v\n", myBin)
	if err := isDir(myBin); err != nil {
		return err
	}
	toAdd := fmt.Sprintf("export PATH=\"$PATH:%v\"", myBin)

	bashRC := filepath.Join(homedir, ".bashrc")
	if err := addToFile(bashRC, toAdd); err != nil {
		return err
	}

	return nil
}

func exportGitConfig(homedir, path string) error {
	if !proceed("Include gitconfig") {
		return nil
	}

	myGitConfig := filepath.Join(path, ".gitconfig")
	fmt.Printf("Including %v\n", myGitConfig)
	if err := isFile(myGitConfig); err != nil {
		return err
	}
	toAdd := fmt.Sprintf("[include]\n  path = %v", myGitConfig)

	gitConfig := filepath.Join(homedir, ".gitconfig")
	if err := createOrAddToFile(gitConfig, toAdd); err != nil {
		return err
	}

	return nil
}

func sourceVimRc(homedir, path string) error {
	if !proceed("Source vmrc") {
		return nil
	}

	myVimRc := filepath.Join(path, ".jonathon_vimrc")
	fmt.Printf("Sourcing %v\n", myVimRc)
	if err := isFile(myVimRc); err != nil {
		return err
	}
	toAdd := fmt.Sprintf("source %v", myVimRc)

	vimRc := filepath.Join(homedir, ".vimrc")
	if err := createOrAddToFile(vimRc, toAdd); err != nil {
		return err
	}

	return nil
}
