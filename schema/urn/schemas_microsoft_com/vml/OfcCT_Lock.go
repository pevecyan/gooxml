// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/pevecyan/gooxml/schema/soo/ofc/sharedTypes"
)

type OfcCT_Lock struct {
	PositionAttr      sharedTypes.ST_TrueFalse
	SelectionAttr     sharedTypes.ST_TrueFalse
	GroupingAttr      sharedTypes.ST_TrueFalse
	UngroupingAttr    sharedTypes.ST_TrueFalse
	RotationAttr      sharedTypes.ST_TrueFalse
	CroppingAttr      sharedTypes.ST_TrueFalse
	VerticiesAttr     sharedTypes.ST_TrueFalse
	AdjusthandlesAttr sharedTypes.ST_TrueFalse
	TextAttr          sharedTypes.ST_TrueFalse
	AspectratioAttr   sharedTypes.ST_TrueFalse
	ShapetypeAttr     sharedTypes.ST_TrueFalse
	ExtAttr           ST_Ext
}

func NewOfcCT_Lock() *OfcCT_Lock {
	ret := &OfcCT_Lock{}
	return ret
}

func (m *OfcCT_Lock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PositionAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PositionAttr.MarshalXMLAttr(xml.Name{Local: "position"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SelectionAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.SelectionAttr.MarshalXMLAttr(xml.Name{Local: "selection"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.GroupingAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.GroupingAttr.MarshalXMLAttr(xml.Name{Local: "grouping"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UngroupingAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UngroupingAttr.MarshalXMLAttr(xml.Name{Local: "ungrouping"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RotationAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.RotationAttr.MarshalXMLAttr(xml.Name{Local: "rotation"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CroppingAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.CroppingAttr.MarshalXMLAttr(xml.Name{Local: "cropping"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VerticiesAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.VerticiesAttr.MarshalXMLAttr(xml.Name{Local: "verticies"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AdjusthandlesAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AdjusthandlesAttr.MarshalXMLAttr(xml.Name{Local: "adjusthandles"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.TextAttr.MarshalXMLAttr(xml.Name{Local: "text"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AspectratioAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AspectratioAttr.MarshalXMLAttr(xml.Name{Local: "aspectratio"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShapetypeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ShapetypeAttr.MarshalXMLAttr(xml.Name{Local: "shapetype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Lock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "position" {
			m.PositionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "selection" {
			m.SelectionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "grouping" {
			m.GroupingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ungrouping" {
			m.UngroupingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rotation" {
			m.RotationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cropping" {
			m.CroppingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "verticies" {
			m.VerticiesAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "adjusthandles" {
			m.AdjusthandlesAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "text" {
			m.TextAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "aspectratio" {
			m.AspectratioAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "shapetype" {
			m.ShapetypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Lock: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Lock and its children
func (m *OfcCT_Lock) Validate() error {
	return m.ValidateWithPath("OfcCT_Lock")
}

// ValidateWithPath validates the OfcCT_Lock and its children, prefixing error messages with path
func (m *OfcCT_Lock) ValidateWithPath(path string) error {
	if err := m.PositionAttr.ValidateWithPath(path + "/PositionAttr"); err != nil {
		return err
	}
	if err := m.SelectionAttr.ValidateWithPath(path + "/SelectionAttr"); err != nil {
		return err
	}
	if err := m.GroupingAttr.ValidateWithPath(path + "/GroupingAttr"); err != nil {
		return err
	}
	if err := m.UngroupingAttr.ValidateWithPath(path + "/UngroupingAttr"); err != nil {
		return err
	}
	if err := m.RotationAttr.ValidateWithPath(path + "/RotationAttr"); err != nil {
		return err
	}
	if err := m.CroppingAttr.ValidateWithPath(path + "/CroppingAttr"); err != nil {
		return err
	}
	if err := m.VerticiesAttr.ValidateWithPath(path + "/VerticiesAttr"); err != nil {
		return err
	}
	if err := m.AdjusthandlesAttr.ValidateWithPath(path + "/AdjusthandlesAttr"); err != nil {
		return err
	}
	if err := m.TextAttr.ValidateWithPath(path + "/TextAttr"); err != nil {
		return err
	}
	if err := m.AspectratioAttr.ValidateWithPath(path + "/AspectratioAttr"); err != nil {
		return err
	}
	if err := m.ShapetypeAttr.ValidateWithPath(path + "/ShapetypeAttr"); err != nil {
		return err
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
