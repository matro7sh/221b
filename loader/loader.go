package loader

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cmepw/221b/logger"
	"github.com/cmepw/221b/templates"
)

type Loader interface {
	Load(content []byte) ([]byte, error)
	Compile(path string, content []byte) error
}

type baseLoader struct{}

const (
	windowsExt = ".exe"
	tmpFile    = "tmp.go"
)

func (b baseLoader) Compile(path string, content []byte) error {
	outputPath := strings.TrimSuffix(path, filepath.Ext(path)) + windowsExt

	dir := "/tmp/test"
	if err := os.MkdirAll(dir, 0750); err != nil {
		logger.Error(fmt.Errorf("could not create temporary directory"))
		return err
	}

	defer func() {
		_ = os.RemoveAll(dir)
	}()

	// Set environment
	logger.Debug("write content to temporary file")
	if err := os.WriteFile(filepath.Join(dir, tmpFile), content, 0666); err != nil {
		logger.Error(fmt.Errorf("could not write tmp file"))
		return err
	}

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(templates.GoMod), 0666); err != nil {
		logger.Error(fmt.Errorf("could not write tmp go.mod file"))
		return err
	}

	initCmd := exec.Command("go", "get", "-u", "golang.org/x/sys/windows")
	initCmd.Dir = dir
	initCmd.Stderr = os.Stderr
	initCmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	if err := initCmd.Run(); err != nil {
		logger.Error(fmt.Errorf("could not install dependency"))
		return err
	}

	logger.Debug("dependency installed")

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	buildCmd := exec.Command(
		"go",
		"build",
		"-ldflags",
		"-s -w -H=windowsgui",
		"-o",
		filepath.Join(pwd, outputPath),
		filepath.Join(dir, tmpFile),
	)
	buildCmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	buildCmd.Stderr = os.Stderr
	buildCmd.Dir = dir

	if err := buildCmd.Run(); err != nil {
		logger.Error(fmt.Errorf("failed to compile"))
		return err
	}

	logger.Info(fmt.Sprintf("file compiled to %s", filepath.Join(pwd, outputPath)))
	return nil
}
