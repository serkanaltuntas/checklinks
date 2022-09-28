# checklinks

Get all links from a web resource and check if any of them broken.

## Build:
```bash
cd <into the code>
go build
```

## Usage:
```bash
./checklinks https://www.example.com/
```

Results will be presented in your current working directory with name 'results.txt'.

## Note:

Solves DNS related issues if used before getting built.

```bash
GODEBUG=netdns=go
```

or 

```bash
GODEBUG=netdns=cgo
```

