
# validation
default:
	go vet
	go build

clean:
	rm -f *~
