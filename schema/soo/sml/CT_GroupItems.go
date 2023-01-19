// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/pevecyan/gooxml"
)

type CT_GroupItems struct {
	// Items Created Count
	CountAttr *uint32
	// No Value
	M []*CT_Missing
	// Numeric Value
	N []*CT_Number
	// Boolean
	B []*CT_Boolean
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
	// Date Time
	D []*CT_DateTime
}

func NewCT_GroupItems() *CT_GroupItems {
	ret := &CT_GroupItems{}
	return ret
}

func (m *CT_GroupItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "ma:m"}}
		for _, c := range m.M {
			e.EncodeElement(c, sem)
		}
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "ma:n"}}
		for _, c := range m.N {
			e.EncodeElement(c, sen)
		}
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "ma:b"}}
		for _, c := range m.B {
			e.EncodeElement(c, seb)
		}
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "ma:e"}}
		for _, c := range m.E {
			e.EncodeElement(c, see)
		}
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "ma:s"}}
		for _, c := range m.S {
			e.EncodeElement(c, ses)
		}
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "ma:d"}}
		for _, c := range m.D {
			e.EncodeElement(c, sed)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_GroupItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "m"}:
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "n"}:
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "b"}:
				tmp := NewCT_Boolean()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "e"}:
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "s"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "d"}:
				tmp := NewCT_DateTime()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.D = append(m.D, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_GroupItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupItems and its children
func (m *CT_GroupItems) Validate() error {
	return m.ValidateWithPath("CT_GroupItems")
}

// ValidateWithPath validates the CT_GroupItems and its children, prefixing error messages with path
func (m *CT_GroupItems) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.D {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/D[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
