desloader: main.go
	 GO111MODULE=on go build -o desloader

build: desloader

clean:
	 rm -f desloader
