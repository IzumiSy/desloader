desload: main.go
	 GO111MODULE=on go build -o desload

build: desload

clean:
	 rm -f desload
