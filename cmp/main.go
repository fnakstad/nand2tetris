package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	inFlag  = flag.String("in", "", ".jack file or directory to process")
	outFlag = flag.String("out", "", "file to save the resulting .xml file")
)

func main() {
	flag.Parse()
	if *inFlag == "" {
		log.Fatalf("-in flag is required")
	}
	if *outFlag == "" {
		log.Fatalf("-out flag is required")
	}

	jackFiles, err := getFilesWithExtension(*inFlag, ".jack")
	if err != nil {
		log.Fatalf("error getting files: %v", err)
	}

	if len(jackFiles) <= 0 {
		log.Fatal("no jack files to handle")
	}

	outf, err := os.Create(*outFlag)
	if err != nil {
		log.Fatalf("error creating out file: %v", err)
	}
	defer outf.Close()

	for _, jackFile := range jackFiles {
		if err := processJackFile(jackFile); err != nil {
			log.Fatalf("error processing jack file: %v", err)
		}
	}
}

func processJackFile(jackFile string) error {
	log.Println(jackFile)

	f, err := os.Open(jackFile)
	if err != nil {
		return err
	}
	defer f.Close()

	t := NewTokenizer(f)
	for t.Next() {
		//log.Println(t.Input())
	}
	if t.Err() != nil {
		return t.Err()
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
