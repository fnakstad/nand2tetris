package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var (
	inFlag   = flag.String("in", "", ".jack file or directory to process")
	toutFlag = flag.String("tout", "", "file to save the resulting tokens .xml file")
)

func main() {
	flag.Parse()
	if *inFlag == "" {
		log.Fatalf("-in flag is required")
	}
	if *toutFlag == "" {
		log.Fatalf("-tout flag is required")
	}

	jackFiles, err := getFilesWithExtension(*inFlag, ".jack")
	if err != nil {
		log.Fatalf("error getting files: %v", err)
	}

	if len(jackFiles) <= 0 {
		log.Fatal("no jack files to handle")
	}

	toutf, err := os.Create(*toutFlag)
	if err != nil {
		log.Fatalf("error creating out file: %v", err)
	}
	defer toutf.Close()

	for _, jackFile := range jackFiles {
		if err := processJackFile(jackFile, toutf); err != nil {
			log.Fatalf("error processing jack file: %v", err)
		}
	}
}

func processJackFile(jackFile string, w io.Writer) error {
	log.Println(jackFile)

	f, err := os.Open(jackFile)
	if err != nil {
		return err
	}
	defer f.Close()

	t := NewTokenizer(f)
	tokens := make([]Token, 0)
	for t.Next() {
		tokens = append(tokens, t.Token())
	}
	if t.Err() != nil {
		return t.Err()
	}

	err = MarshalTokens(tokens, w)
	if err != nil {
		return err
	}

	return nil
}

func getFilesWithExtension(p, validExt string) ([]string, error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		names, err := file.Readdirnames(0)
		if err != nil {
			return nil, err
		}
		var valid []string
		for _, name := range names {
			if path.Ext(name) == validExt {
				valid = append(valid, path.Join(p, name))
			}
		}

		return valid, nil
	}

	ext := path.Ext(stat.Name())
	if ext != validExt {
		return nil, fmt.Errorf("given file has invalid extension: %s", ext)
	}

	return []string{p}, nil
}
