
# validation
default:
	go vet ./client/
	go build ./client/

clean:
	rm -f *~
