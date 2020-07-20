package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("not enough arguments")
	}

	vmfiles, err := getVMFiles(os.Args[1])
	if err != nil {
		log.Fatalf("error getting VM files: %v", err)
	}

	if len(vmfiles) <= 0 {
		log.Fatal("no vm files to handle")
	}

	outf, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("error creating out file: %v", err)
	}
	defer outf.Close()

	cw := NewCodeWriter(outf)

	// log.Println(vmfiles)
	for _, vmfile := range vmfiles {
		log.Println(vmfile)

		f, err := os.Open(vmfile)
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}
		defer f.Close()

		p := NewParser(f)
		for p.Parse() {
			switch p.CommandType() {
			case CommandTypeArithmetic:
				cw.WriteArithmetic(p.Arg1())
			case CommandTypePush, CommandTypePop:
				cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
			default:
				log.Fatalf("unimplemented command: %d", p.CommandType())
			}
			// log.Println("h")
			// log.Println(p.CommandType())
			// log.Println(p.Arg1())
			// log.Println(p.Arg2())
		}
		if err := p.Err(); err != nil {
			log.Fatalf("error parsing: %v", err)
		}
	}
}

const (
	vmext = ".vm"
)

func getVMFiles(p string) ([]string, error) {
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
			if path.Ext(name) == vmext {
				valid = append(valid, path.Join(p, name))
			}
		}

		return valid, nil
	}

	ext := path.Ext(stat.Name())
	if ext != vmext {
		return nil, fmt.Errorf("given file has invalid extension: %s", ext)
	}

	return []string{p}, nil
}
