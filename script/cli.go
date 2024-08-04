package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/stkali/utility/tool"
	"golang.org/x/sync/errgroup"
)

const (
	ReadMeFileName = "README.md"
	MakeFileName   = "Makefile"
	GoModFileName  = "go.mod"
	MainFileName   = "main.go"
	FilePermission = 0o644
)

var (
	NotFoundChapterNameError = errors.New("cannot found chapter name")
	EmptyChapterNameError    = errors.New("empty chapter name")
)

type Options struct {
	Name string
	Path string
}

func printUsage() {
	_, _ = fmt.Fprintf(
		os.Stdout,
		"cli [CHAPTER_NAME]\n    Creates chapter template.\nSample:\n    cli 01认识Golang\n")
}

// getCh
func getChapterName() (string, error) {

	if len(os.Args) < 2 {
		return "", NotFoundChapterNameError
	}
	chapterName := strings.Trim(os.Args[1], " ")
	if chapterName == "" {
		return "", EmptyChapterNameError
	}
	if chapterName[0] == '-' {
		printUsage()
		os.Exit(0)
	}
	return chapterName, nil
}

type FileMaker func(options *Options) error

func createFile(file string, content string) error {
	fd, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, FilePermission)
	if err != nil {
		return err
	}
	_, err = io.WriteString(fd, content)
	if err != nil {
		return err
	}
	return nil
}

func createMakefile(options *Options) error {
	makefile := filepath.Join(options.Path, MakeFileName)
	tmpl := `PROGRAM=%s
REMOVE=@rm -f
PRINT=@echo
		
build:
	@go build -o bin/$(PROGRAM) main.go
	$(PRINT) successfully build $(PROGRAM) at 'bin'.
		
clear:
	$(REMOVE) bin/*
	$(PRINT) successfully clear the build workspace.
`
	content := fmt.Sprintf(tmpl, options.Name)
	return createFile(makefile, content)
}

func createGoModFile(options *Options) error {
	modFile := filepath.Join(options.Path, GoModFileName)
	tmpl := `module github.com/stkali/go-basics/%s

go 1.22.5
`
	content := fmt.Sprintf(tmpl, options.Name)
	return createFile(modFile, content)
}

func createMainFile(options *Options) error {
	mainFile := filepath.Join(options.Path, MainFileName)
	content := `package main

func main() {

}
`
	return createFile(mainFile, content)
}

func createReadmeFile(options *Options) error {
	readmeFile := filepath.Join(options.Path, ReadMeFileName)
	content := fmt.Sprintf("# %s\n", options.Name)
	return createFile(readmeFile, content)
}

// cli create basic template for vidoa
func cli() error {

	chapterName, err := getChapterName()
	if err != nil {
		return fmt.Errorf("failed to get chapter name, err: %s", err)
	}

	chapterPath := tool.ToAbsPath(chapterName)
	options := &Options{
		Name: filepath.Base(chapterPath),
		Path: chapterPath,
	}

	// create chapter folder
	err = os.MkdirAll(chapterName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to make chapter folder, err: %s", err)
	}

	fileMakers := []FileMaker{
		createGoModFile,
		createMakefile,
		createMainFile,
		createReadmeFile,
	}

	g, ctx := errgroup.WithContext(context.Background())

	for index := range fileMakers {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return fileMakers[index](options)
			}
		})
	}
	return g.Wait()
}

func main() {
	err := cli()
	if err != nil {

		_, _ = fmt.Fprintf(os.Stderr, "cli error: %s\n", err)
	}
}
