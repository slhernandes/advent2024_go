tidy:
	go mod tidy

test_small:
	go test -run TestSmall

test:
	go test

run: tidy test

debug: tidy test_small



.PHONY: run tidy test test_small debug
.SILENT:
