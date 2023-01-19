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
	"strconv"

	"github.com/pevecyan/gooxml"
)

type Headers struct {
	CT_RevisionHeaders
}

func NewHeaders() *Headers {
	ret := &Headers{}
	ret.CT_RevisionHeaders = *NewCT_RevisionHeaders()
	return ret
}

func (m *Headers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:headers"
	return m.CT_RevisionHeaders.MarshalXML(e, start)
}

func (m *Headers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_RevisionHeaders = *NewCT_RevisionHeaders()
	for _, attr := range start.Attr {
		if attr.Name.Local == "exclusive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ExclusiveAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastGuid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastGuidAttr = &parsed
			continue
		}
		if attr.Name.Local == "shared" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SharedAttr = &parsed
			continue
		}
		if attr.Name.Local == "diskRevisions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DiskRevisionsAttr = &parsed
			continue
		}
		if attr.Name.Local == "history" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HistoryAttr = &parsed
			continue
		}
		if attr.Name.Local == "trackRevisions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TrackRevisionsAttr = &parsed
			continue
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
			continue
		}
		if attr.Name.Local == "revisionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RevisionIdAttr = &pt
			continue
		}
		if attr.Name.Local == "version" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.VersionAttr = &pt
			continue
		}
		if attr.Name.Local == "keepChangeHistory" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.KeepChangeHistoryAttr = &parsed
			continue
		}
		if attr.Name.Local == "protected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ProtectedAttr = &parsed
			continue
		}
		if attr.Name.Local == "preserveHistory" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PreserveHistoryAttr = &pt
			continue
		}
	}
lHeaders:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "header"}:
				tmp := NewCT_RevisionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Header = append(m.Header, tmp)
			default:
				gooxml.Log("skipping unsupported element on Headers %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lHeaders
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Headers and its children
func (m *Headers) Validate() error {
	return m.ValidateWithPath("Headers")
}

// ValidateWithPath validates the Headers and its children, prefixing error messages with path
func (m *Headers) ValidateWithPath(path string) error {
	if err := m.CT_RevisionHeaders.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
