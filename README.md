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

## Example

```shell
221b bake -k shflfhje -s test.sh
go: added golang.org/x/sys v0.10.0
[+] file compiled to ./test.exe
```