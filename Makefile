
# validation
default:
	go vet ./client/
	go build ./client/
	go build ./examples/listvol.go

clean:
	rm -f *~ */*~
	rm -f listvol
	go clean
