// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/pevecyan/gooxml"
	"github.com/pevecyan/gooxml/schema/soo/dml"
)

type CT_AbsoluteAnchor struct {
	Pos        *dml.CT_Point2D
	Ext        *dml.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

func NewCT_AbsoluteAnchor() *CT_AbsoluteAnchor {
	ret := &CT_AbsoluteAnchor{}
	ret.Pos = dml.NewCT_Point2D()
	ret.Ext = dml.NewCT_PositiveSize2D()
	ret.ClientData = NewCT_AnchorClientData()
	return ret
}

func (m *CT_AbsoluteAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sepos := xml.StartElement{Name: xml.Name{Local: "xdr:pos"}}
	e.EncodeElement(m.Pos, sepos)
	seext := xml.StartElement{Name: xml.Name{Local: "xdr:ext"}}
	e.EncodeElement(m.Ext, seext)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	seclientData := xml.StartElement{Name: xml.Name{Local: "xdr:clientData"}}
	e.EncodeElement(m.ClientData, seclientData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AbsoluteAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = dml.NewCT_Point2D()
	m.Ext = dml.NewCT_PositiveSize2D()
	m.ClientData = NewCT_AnchorClientData()
lCT_AbsoluteAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "pos"}:
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "ext"}:
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "sp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "grpSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "graphicFrame"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cxnSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "pic"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "contentPart"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.ContentPart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "clientData"}:
				if err := d.DecodeElement(m.ClientData, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_AbsoluteAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AbsoluteAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AbsoluteAnchor and its children
func (m *CT_AbsoluteAnchor) Validate() error {
	return m.ValidateWithPath("CT_AbsoluteAnchor")
}

// ValidateWithPath validates the CT_AbsoluteAnchor and its children, prefixing error messages with path
func (m *CT_AbsoluteAnchor) ValidateWithPath(path string) error {
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.ClientData.ValidateWithPath(path + "/ClientData"); err != nil {
		return err
	}
	return nil
}
