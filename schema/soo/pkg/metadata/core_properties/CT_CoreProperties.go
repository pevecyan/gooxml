// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties

import (
	"encoding/xml"
	"time"

	"github.com/pevecyan/gooxml"
)

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *gooxml.XSDAny
	Creator        *gooxml.XSDAny
	Description    *gooxml.XSDAny
	Identifier     *gooxml.XSDAny
	Keywords       *CT_Keywords
	Language       *gooxml.XSDAny
	LastModifiedBy *string
	LastPrinted    *time.Time
	Modified       *gooxml.XSDAny
	Revision       *string
	Subject        *gooxml.XSDAny
	Title          *gooxml.XSDAny
	Version        *string
}

func NewCT_CoreProperties() *CT_CoreProperties {
	ret := &CT_CoreProperties{}
	return ret
}

func (m *CT_CoreProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Category != nil {
		secategory := xml.StartElement{Name: xml.Name{Local: "cp:category"}}
		gooxml.AddPreserveSpaceAttr(&secategory, *m.Category)
		e.EncodeElement(m.Category, secategory)
	}
	if m.ContentStatus != nil {
		secontentStatus := xml.StartElement{Name: xml.Name{Local: "cp:contentStatus"}}
		gooxml.AddPreserveSpaceAttr(&secontentStatus, *m.ContentStatus)
		e.EncodeElement(m.ContentStatus, secontentStatus)
	}
	if m.Created != nil {
		secreated := xml.StartElement{Name: xml.Name{Local: "dcterms:created"}}
		e.EncodeElement(m.Created, secreated)
	}
	if m.Creator != nil {
		secreator := xml.StartElement{Name: xml.Name{Local: "dc:creator"}}
		e.EncodeElement(m.Creator, secreator)
	}
	if m.Description != nil {
		sedescription := xml.StartElement{Name: xml.Name{Local: "dc:description"}}
		e.EncodeElement(m.Description, sedescription)
	}
	if m.Identifier != nil {
		seidentifier := xml.StartElement{Name: xml.Name{Local: "dc:identifier"}}
		e.EncodeElement(m.Identifier, seidentifier)
	}
	if m.Keywords != nil {
		sekeywords := xml.StartElement{Name: xml.Name{Local: "cp:keywords"}}
		e.EncodeElement(m.Keywords, sekeywords)
	}
	if m.Language != nil {
		selanguage := xml.StartElement{Name: xml.Name{Local: "dc:language"}}
		e.EncodeElement(m.Language, selanguage)
	}
	if m.LastModifiedBy != nil {
		selastModifiedBy := xml.StartElement{Name: xml.Name{Local: "cp:lastModifiedBy"}}
		gooxml.AddPreserveSpaceAttr(&selastModifiedBy, *m.LastModifiedBy)
		e.EncodeElement(m.LastModifiedBy, selastModifiedBy)
	}
	if m.LastPrinted != nil {
		selastPrinted := xml.StartElement{Name: xml.Name{Local: "cp:lastPrinted"}}
		e.EncodeElement(m.LastPrinted, selastPrinted)
	}
	if m.Modified != nil {
		semodified := xml.StartElement{Name: xml.Name{Local: "dcterms:modified"}}
		e.EncodeElement(m.Modified, semodified)
	}
	if m.Revision != nil {
		serevision := xml.StartElement{Name: xml.Name{Local: "cp:revision"}}
		gooxml.AddPreserveSpaceAttr(&serevision, *m.Revision)
		e.EncodeElement(m.Revision, serevision)
	}
	if m.Subject != nil {
		sesubject := xml.StartElement{Name: xml.Name{Local: "dc:subject"}}
		e.EncodeElement(m.Subject, sesubject)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "dc:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.Version != nil {
		seversion := xml.StartElement{Name: xml.Name{Local: "cp:version"}}
		gooxml.AddPreserveSpaceAttr(&seversion, *m.Version)
		e.EncodeElement(m.Version, seversion)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CoreProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CoreProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "category"}:
				m.Category = new(string)
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "contentStatus"}:
				m.ContentStatus = new(string)
				if err := d.DecodeElement(m.ContentStatus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "created"}:
				m.Created = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Created, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "creator"}:
				m.Creator = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Creator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "description"}:
				m.Description = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "identifier"}:
				m.Identifier = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Identifier, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "keywords"}:
				m.Keywords = NewCT_Keywords()
				if err := d.DecodeElement(m.Keywords, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "language"}:
				m.Language = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Language, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastModifiedBy"}:
				m.LastModifiedBy = new(string)
				if err := d.DecodeElement(m.LastModifiedBy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastPrinted"}:
				m.LastPrinted = new(time.Time)
				if err := d.DecodeElement(m.LastPrinted, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "modified"}:
				m.Modified = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Modified, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "revision"}:
				m.Revision = new(string)
				if err := d.DecodeElement(m.Revision, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "subject"}:
				m.Subject = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Subject, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "title"}:
				m.Title = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "version"}:
				m.Version = new(string)
				if err := d.DecodeElement(m.Version, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_CoreProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CoreProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CoreProperties and its children
func (m *CT_CoreProperties) Validate() error {
	return m.ValidateWithPath("CT_CoreProperties")
}

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (m *CT_CoreProperties) ValidateWithPath(path string) error {
	if m.Keywords != nil {
		if err := m.Keywords.ValidateWithPath(path + "/Keywords"); err != nil {
			return err
		}
	}
	return nil
}
