// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
}

func NewCT_Keywords() *CT_Keywords {
	ret := &CT_Keywords{}
	return ret
}
func (m *CT_Keywords) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml:lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Value != nil {
		sevalue := xml.StartElement{Name: xml.Name{Local: "cp:value"}}
		e.EncodeElement(m.Value, sevalue)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Keywords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
		}
	}
lCT_Keywords:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "value":
				tmp := NewCT_Keyword()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Value = append(m.Value, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Keywords
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Keywords) Validate() error {
	return m.ValidateWithPath("CT_Keywords")
}
func (m *CT_Keywords) ValidateWithPath(path string) error {
	for i, v := range m.Value {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Value[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
