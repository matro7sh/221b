package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cmepw/221b/encryption"
	"github.com/cmepw/221b/logger"
	"github.com/cmepw/221b/templates"
)

// Method gather all loading and encryption method.
var Method = map[string]Loader{
	"xor":      Xor{},
	"aes":      Aes{},
	"chacha20": ChaCha20{},
}

type Loader interface {
	Load(content, key []byte) ([]byte, error)
	Compile(path string, content []byte) error
	encryption.Encryption
}

type baseLoader struct{}

const (
	tmpFile = "tmp.go"
	tmpDir  = "/tmp/221b-compile"
)

var (
	base_path, _ = os.Getwd()
)

func (b baseLoader) Compile(outputPath string, content []byte) error {
	if err := b.setupTmpDir(content); err != nil {
		return err
	}
	defer func() {
		logger.Debug(fmt.Sprintf("cleanup temporary dir %s", tmpDir))
		_ = os.RemoveAll(tmpDir)
	}()

	err := b.execCmd("go", "get", "-u", "golang.org/x/sys/windows", "golang.org/x/crypto")
	if err != nil {
		logger.Error(fmt.Errorf("could not install dependency"))
		return err
	}

	logger.Debug("dependency installed")
	logger.Debug("start compiling binary")

	relOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}

	_ = b.execCmd(
		"go",
		"generate",
	)

	err = b.execCmd(
		"go",
		"build",
		"-ldflags",
		"-s -w -H=windowsgui",
		"-o",
		relOutputPath,
	)
	if err != nil {
		logger.Error(fmt.Errorf("failed to compile"))
		return err
	}

	logger.Info(fmt.Sprintf("file compiled to %s", relOutputPath))

	return nil
}

func (b baseLoader) execCmd(name string, args ...string) error {
	logger.Debug(fmt.Sprintf("execute command %s", name))

	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	cmd.Stderr = os.Stderr
	cmd.Dir = tmpDir

	return cmd.Run()
}

func (b baseLoader) setupTmpDir(goFile []byte) error {
	logger.Debug(fmt.Sprintf("setup temporary directory %s", tmpDir))

	if err := os.MkdirAll(tmpDir, 0750); err != nil {
		logger.Error(fmt.Errorf("could not create temporary directory"))
		return err
	}

	data, err := ioutil.ReadFile(filepath.Join(base_path, "versioninfo.json"))
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile(filepath.Join(tmpDir, "versioninfo.json"), data, 0666)

	if err := os.WriteFile(filepath.Join(tmpDir, tmpFile), goFile, 0666); err != nil {
		logger.Error(fmt.Errorf("could not write tmp file"))
		return err
	}

	if err := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(templates.GoMod), 0666); err != nil {
		logger.Error(fmt.Errorf("could not write tmp go.mod file"))
		return err
	}

	return nil
}
