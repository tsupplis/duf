.PHONY: clean push push-openbsd-amd64 push-netbsd-amd64 push-freebsd-amd64 push-linux-amd64 local

local: duf

all: local duf-openbsd-amd64 duf-netbsd-amd64 duf-freebsd-amd64 \
	duf-linux-amd64 duf-linux-arm64 duf-linux-riscv64 duf-dragonfly-amd64 \
	duf-illumos-amd64 duf-solaris-amd64 duf-windows-amd64.exe
	
duf: 
	go build -ldflags="-s -w" -o $@ $*

duf-windows-amd64.exe: main.go 
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-openbsd-amd64: main.go 
	GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-dragonfly-amd64: main.go  
	GOOS=dragonfly GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-netbsd-amd64: main.go  
	GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-solaris-amd64: main.go 
	GOOS=solaris GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-illumos-amd64: main.go 
	GOOS=illumos GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-linux-riscv64: main.go 
	GOOS=linux GOARCH=riscv64 go build -ldflags="-s -w" -o $@ $*

duf-freebsd-amd64: main.go 
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-linux-arm64: main.go 
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $@ $*

duf-linux-amd64: main.go 
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $@ $*

duf-linux-ppc64le: main.go 
	GOOS=linux GOARCH=ppc64le go build -ldflags="-s -w" -o $@ $*

clean:
	rm -f duf duf-openbsd-amd64 duf-netbsd-amd64 duf-dragonfly-amd64 \
	duf-freebsd-amd64 duf-linux-amd64 duf-linux-ppc64le \
	duf-linux-riscv64 duf-illumos-amd64 duf-solaris-amd64 duf-windows-amd64
