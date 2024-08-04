update_cli:
	cd script && go build -o ../cli cli.go && cd -
	@echo Successfully updated cli.
