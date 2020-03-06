BIN_NAME=app

all: test

install:
	echo install
	go build -o $(BIN_NAME)
	./app

test:
	go test ./...

uninstall:
	echo uninstall

clean:
	$(RM) $(BIN_NAME)