all:test

test:
	go build -o main *.go
clean:
	rm main
