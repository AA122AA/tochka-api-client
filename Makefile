PROJECTNAME=$(shell basename "$(PWD)")

## build: Compile the binary.
build:
	@mkdir -p bin/
	@echo "Building tochka client..."
	@/usr/local/go/bin/go build -o ./bin/client $(PWD)/cmd/tochkaClient
	@strip ./bin/client

## start: Start the binary. Runs `build` internally.
start: build
	bin/client 

## run: Run the binary. 
run: 
	bin/client 

## clean: Remove the binary.
clean:
	@rm -rf bin/


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
