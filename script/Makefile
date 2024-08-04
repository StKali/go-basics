PROGRAM=cli
REMOVE=@rm -f
PRINT=@echo
		
build_cli:
	@go build -o bin/$(PROGRAM) cli.go
	$(PRINT) successfully build $(PROGRAM) at 'bin'.
		
clear:
	$(REMOVE) bin/*
	$(PRINT) successfully clear the build workspace.

