package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const (
	vmext          = ".vm"
	globalFuncName = "global"
)

var (
	inFlag        = flag.String("in", "", "vm file or directory to process")
	outFlag       = flag.String("out", "", "file to save the resulting .asm file")
	bootstrapFlag = flag.Bool("bootstrap", true, "whether to add bootstrap asm")
)

func main() {
	flag.Parse()

	if *inFlag == "" {
		log.Fatalf("-in flag is required")
	}
	if *outFlag == "" {
		log.Fatalf("-out flag is required")
	}

	vmfiles, err := getVMFiles(*inFlag)
	if err != nil {
		log.Fatalf("error getting VM files: %v", err)
	}

	if len(vmfiles) <= 0 {
		log.Fatal("no vm files to handle")
	}

	outf, err := os.Create(*outFlag)
	if err != nil {
		log.Fatalf("error creating out file: %v", err)
	}
	defer outf.Close()

	cw := NewCodeWriter(outf)

	if *bootstrapFlag {
		err = cw.WriteBootstrap()
		err = cw.WriteCall("Sys.init", 0)
		if err != nil {
			log.Fatalf("error writing bootstrap: %v", err)
		}
	}

	for _, vmfile := range vmfiles {
		if err := processVMFile(vmfile, cw); err != nil {
			log.Fatalf("error processing vm file: %v", err)
		}
	}
}

func processVMFile(vmfile string, cw *CodeWriter) error {
	log.Println(vmfile)
	fid := strings.TrimSuffix(
		path.Base(vmfile),
		path.Ext(vmfile),
	)
	f, err := os.Open(vmfile)
	if err != nil {
		return err
	}
	defer f.Close()

	p := NewParser(f)
	currentFunc := globalFuncName // tracks which function is currently being processed
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
			err := cw.WritePushPop(p.CommandType(), segmentType, p.Arg2(), fid)
			if err != nil {
				return err
			}
		case CommandTypeLabel:
			err := cw.WriteLabel(currentFunc, p.Arg1())
			if err != nil {
				return err
			}
		case CommandTypeIf:
			err := cw.WriteIf(currentFunc, p.Arg1())
			if err != nil {
				return err
			}
		case CommandTypeGoto:
			err := cw.WriteGoto(currentFunc, p.Arg1())
			if err != nil {
				return err
			}
		case CommandTypeFunction:
			err := cw.WriteFunction(p.Arg1(), p.Arg2())
			if err != nil {
				return err
			}
			currentFunc = p.Arg1()
		case CommandTypeReturn:
			err := cw.WriteReturn()
			if err != nil {
				return err
			}
		case CommandTypeCall:
			err := cw.WriteCall(p.Arg1(), p.Arg2())
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
