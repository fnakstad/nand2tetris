package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const (
	vmext = ".vm"
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
	for _, vmfile := range vmfiles {
		if err := processVMFile(vmfile, cw); err != nil {
			log.Fatalf("error processing vm file: %v", err)
		}
	}
}

func processVMFile(vmfile string, cw *CodeWriter) error {
	log.Println(vmfile)
	f, err := os.Open(vmfile)
	if err != nil {
		return err
	}
	defer f.Close()

	p := NewParser(f)
	for p.Parse() {
		if _, ok := arithmeticCommand[p.CommandType()]; ok {
			cmdType := CommandType(p.Arg1())
			err := cw.WriteArithmetic(cmdType)
			if err != nil {
				return err
			}
		}
		switch p.CommandType() {
		case CommandTypePush, CommandTypePop:
			segmentType := SegmentType(p.Arg1())
			err := cw.WritePushPop(p.CommandType(), segmentType, p.Arg2())
			if err != nil {
				return err
			}
			// default:
			// 	return fmt.Errorf("unimplemented command: %d", p.CommandType())
		}
		// log.Println("h")
		// log.Println(p.CommandType())
		// log.Println(p.Arg1())
		// log.Println(p.Arg2())
	}
	if err := p.Err(); err != nil {
		return err
	}

	return nil
}

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
