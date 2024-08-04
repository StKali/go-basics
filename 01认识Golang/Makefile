PROGRAM=helloworld
REMOVE=@rm -f
PRINT=@echo

build:
	@go build -o bin/$(PROGRAM) main.go
	$(PRINT) successfully build $(PROGRAM) binary.

clear:
	$(REMOVE) bin/*
	$(PRINT) Successfully cleaned up build intermediate files.
