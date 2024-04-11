clean:
	@echo "Cleaning up"
	rm -rf bin/*
build-vm:
	@echo "Building VM"
	go build -o bin/flux.exe cmd/vm/main.go