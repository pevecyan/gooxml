// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/pevecyan/gooxml"
)

type CT_Bar struct {
	BarPr *CT_BarPr
	E     *CT_OMathArg
}

func NewCT_Bar() *CT_Bar {
	ret := &CT_Bar{}
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_Bar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BarPr != nil {
		sebarPr := xml.StartElement{Name: xml.Name{Local: "m:barPr"}}
		e.EncodeElement(m.BarPr, sebarPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Bar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
lCT_Bar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "barPr"}:
				m.BarPr = NewCT_BarPr()
				if err := d.DecodeElement(m.BarPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Bar %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Bar
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Bar and its children
func (m *CT_Bar) Validate() error {
	return m.ValidateWithPath("CT_Bar")
}

// ValidateWithPath validates the CT_Bar and its children, prefixing error messages with path
func (m *CT_Bar) ValidateWithPath(path string) error {
	if m.BarPr != nil {
		if err := m.BarPr.ValidateWithPath(path + "/BarPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
