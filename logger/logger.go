package logger

import (
	"fmt"
	"os"
)

func Debug(msg string) {
	if DebugMode {
		fmt.Printf("[*] %s\n", msg)
	}
}

func Info(msg string) {
	fmt.Printf("[+] %s\n", msg)
}

func Warn(msg string) {
	fmt.Printf("[^] %s\n", msg)
}

func Error(err error) {
	fmt.Printf("[!] %v\n", err)
}

func Fatal(err error) {
	Error(err)
	os.Exit(1)
}
