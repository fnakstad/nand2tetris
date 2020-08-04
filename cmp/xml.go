package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func (t Token) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var val string
	switch t.Type {
	case TokenTypeSymbol:
		val = string(t.Symbol)
	case TokenTypeIdentifier:
		val = t.Identifier
	case TokenTypeKeyword:
		val = string(t.Keyword)
	case TokenTypeStringConst:
		val = t.StringVal
	case TokenTypeIntConst:
		val = strconv.Itoa(t.IntVal)
	default:
		return errors.New("unmarshalable token")
	}

	s := xml.StartElement{Name: xml.Name{Local: string(t.Type)}}
	err := e.EncodeToken(s)
	if err != nil {
		return err
	}
	err = e.EncodeToken(xml.CharData(fmt.Sprintf("%[2]s%[1]s%[2]s", val, " ")))
	if err != nil {
		return err
	}
	return e.EncodeToken(s.End())
}

func MarshalTokens(tokens []Token, w io.Writer) error {
	type container struct {
		XMLName xml.Name `xml:"tokens"`
		Tokens  []Token
	}

	x, err := xml.MarshalIndent(container{Tokens: tokens}, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(x)
	return err
}
