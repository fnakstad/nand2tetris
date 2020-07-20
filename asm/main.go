package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"fnakstad.com/asm/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	p := parser.New(file)
	st := NewSymbolTable()
	err = addLabels(p, st)
	if err != nil {
		log.Fatalf("error building symbol table: %v", err)
	}

	p.Rewind()
	err = translate(p, st, os.Stdout)
	if err != nil {
		log.Fatalf("error translating: %v", err)
	}
	log.Println("great success")
}

type SymbolTable map[string]uint16

func NewSymbolTable() SymbolTable {
	return SymbolTable(map[string]uint16{
		"SP":     0x0000,
		"LCL":    0x0001,
		"ARG":    0x0002,
		"THIS":   0x0003,
		"THAT":   0x0004,
		"R0":     0x0000,
		"R1":     0x0001,
		"R2":     0x0002,
		"R3":     0x0003,
		"R4":     0x0004,
		"R5":     0x0005,
		"R6":     0x0006,
		"R7":     0x0007,
		"R8":     0x0008,
		"R9":     0x0009,
		"R10":    0x000a,
		"R11":    0x000b,
		"R12":    0x000c,
		"R13":    0x000d,
		"R14":    0x000e,
		"R15":    0x000f,
		"SCREEN": 0x4000,
		"KBD":    0x6000,
	})
}

func addLabels(p *parser.Parser, st SymbolTable) error {

	var l uint16 = 0x0
	for p.Parse() {
		if p.CommandType() == parser.CommandTypeL {
			if _, ok := st[p.Symbol()]; !ok {
				st[p.Symbol()] = l
			}
		} else {
			l++
		}
	}
	if err := p.Err(); err != nil {
		return err
	}

	return nil
}

func translate(p *parser.Parser, st SymbolTable, w io.Writer) error {
	var i uint16 = 0x10
	for p.Parse() {
		// Populate symbol table if encountering an A command (not decimal)
		if p.CommandType() == parser.CommandTypeA && p.Symbol() != "" {
			if _, ok := st[p.Symbol()]; !ok {
				st[p.Symbol()] = i
				i++
			}
		}

		if p.CommandType() == parser.CommandTypeA {
			var address uint16
			if p.Symbol() != "" {
				address = st[p.Symbol()]
			} else {
				address = p.Decimal()
			}
			_, err := fmt.Fprintf(w, "0%015b\n", address)
			if err != nil {
				return err
			}
		}

		if p.CommandType() == parser.CommandTypeC {
			_, err := fmt.Fprintf(w, "111%s%s%s\n", p.Comp(), p.Dest(), p.Jump())
			if err != nil {
				return err
			}
		}

	}
	if err := p.Err(); err != nil {
		return err
	}
	return nil
}
