
all: build

build:
	go build

test:
	go test -v ./

clean:
	rm -f jc jc.test
