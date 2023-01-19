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

type CT_QueryTable struct {
	// QueryTable Name
	NameAttr string
	// First Row Column Titles
	HeadersAttr *bool
	// Row Numbers
	RowNumbersAttr *bool
	// Disable Refresh
	DisableRefreshAttr *bool
	// Background Refresh
	BackgroundRefreshAttr *bool
	// First Background Refresh
	FirstBackgroundRefreshAttr *bool
	// Refresh On Load
	RefreshOnLoadAttr *bool
	// Grow Shrink Type
	GrowShrinkTypeAttr ST_GrowShrinkType
	// Fill Adjacent Formulas
	FillFormulasAttr *bool
	// Remove Data On Save
	RemoveDataOnSaveAttr *bool
	// Disable Edit
	DisableEditAttr *bool
	// Preserve Formatting On Refresh
	PreserveFormattingAttr *bool
	// Adjust Column Width On Refresh
	AdjustColumnWidthAttr *bool
	// Intermediate
	IntermediateAttr *bool
	// Connection Id
	ConnectionIdAttr uint32
	// QueryTable Refresh Information
	QueryTableRefresh *CT_QueryTableRefresh
	// Future Feature Data Storage Area
	ExtLst                      *CT_ExtensionList
	AutoFormatIdAttr            *uint32
	ApplyNumberFormatsAttr      *bool
	ApplyBorderFormatsAttr      *bool
	ApplyFontFormatsAttr        *bool
	ApplyPatternFormatsAttr     *bool
	ApplyAlignmentFormatsAttr   *bool
	ApplyWidthHeightFormatsAttr *bool
}

func NewCT_QueryTable() *CT_QueryTable {
	ret := &CT_QueryTable{}
	return ret
}

func (m *CT_QueryTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.HeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headers"},
			Value: fmt.Sprintf("%d", b2i(*m.HeadersAttr))})
	}
	if m.RowNumbersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rowNumbers"},
			Value: fmt.Sprintf("%d", b2i(*m.RowNumbersAttr))})
	}
	if m.DisableRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disableRefresh"},
			Value: fmt.Sprintf("%d", b2i(*m.DisableRefreshAttr))})
	}
	if m.BackgroundRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "backgroundRefresh"},
			Value: fmt.Sprintf("%d", b2i(*m.BackgroundRefreshAttr))})
	}
	if m.FirstBackgroundRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstBackgroundRefresh"},
			Value: fmt.Sprintf("%d", b2i(*m.FirstBackgroundRefreshAttr))})
	}
	if m.RefreshOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshOnLoad"},
			Value: fmt.Sprintf("%d", b2i(*m.RefreshOnLoadAttr))})
	}
	if m.GrowShrinkTypeAttr != ST_GrowShrinkTypeUnset {
		attr, err := m.GrowShrinkTypeAttr.MarshalXMLAttr(xml.Name{Local: "growShrinkType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillFormulas"},
			Value: fmt.Sprintf("%d", b2i(*m.FillFormulasAttr))})
	}
	if m.RemoveDataOnSaveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "removeDataOnSave"},
			Value: fmt.Sprintf("%d", b2i(*m.RemoveDataOnSaveAttr))})
	}
	if m.DisableEditAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disableEdit"},
			Value: fmt.Sprintf("%d", b2i(*m.DisableEditAttr))})
	}
	if m.PreserveFormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveFormatting"},
			Value: fmt.Sprintf("%d", b2i(*m.PreserveFormattingAttr))})
	}
	if m.AdjustColumnWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "adjustColumnWidth"},
			Value: fmt.Sprintf("%d", b2i(*m.AdjustColumnWidthAttr))})
	}
	if m.IntermediateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "intermediate"},
			Value: fmt.Sprintf("%d", b2i(*m.IntermediateAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
		Value: fmt.Sprintf("%v", m.ConnectionIdAttr)})
	if m.AutoFormatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFormatId"},
			Value: fmt.Sprintf("%v", *m.AutoFormatIdAttr)})
	}
	if m.ApplyNumberFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyNumberFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyNumberFormatsAttr))})
	}
	if m.ApplyBorderFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyBorderFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyBorderFormatsAttr))})
	}
	if m.ApplyFontFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFontFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyFontFormatsAttr))})
	}
	if m.ApplyPatternFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyPatternFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyPatternFormatsAttr))})
	}
	if m.ApplyAlignmentFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyAlignmentFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyAlignmentFormatsAttr))})
	}
	if m.ApplyWidthHeightFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyWidthHeightFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyWidthHeightFormatsAttr))})
	}
	e.EncodeToken(start)
	if m.QueryTableRefresh != nil {
		sequeryTableRefresh := xml.StartElement{Name: xml.Name{Local: "ma:queryTableRefresh"}}
		e.EncodeElement(m.QueryTableRefresh, sequeryTableRefresh)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_QueryTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "adjustColumnWidth" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdjustColumnWidthAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "intermediate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IntermediateAttr = &parsed
			continue
		}
		if attr.Name.Local == "rowNumbers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowNumbersAttr = &parsed
			continue
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ConnectionIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "backgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
			continue
		}
		if attr.Name.Local == "fillFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FillFormulasAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstBackgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FirstBackgroundRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
			continue
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "disableRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "preserveFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormattingAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "growShrinkType" {
			m.GrowShrinkTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "disableEdit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableEditAttr = &parsed
			continue
		}
		if attr.Name.Local == "headers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadersAttr = &parsed
			continue
		}
		if attr.Name.Local == "removeDataOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RemoveDataOnSaveAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
			continue
		}
	}
lCT_QueryTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryTableRefresh"}:
				m.QueryTableRefresh = NewCT_QueryTableRefresh()
				if err := d.DecodeElement(m.QueryTableRefresh, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_QueryTable %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_QueryTable and its children
func (m *CT_QueryTable) Validate() error {
	return m.ValidateWithPath("CT_QueryTable")
}

// ValidateWithPath validates the CT_QueryTable and its children, prefixing error messages with path
func (m *CT_QueryTable) ValidateWithPath(path string) error {
	if err := m.GrowShrinkTypeAttr.ValidateWithPath(path + "/GrowShrinkTypeAttr"); err != nil {
		return err
	}
	if m.QueryTableRefresh != nil {
		if err := m.QueryTableRefresh.ValidateWithPath(path + "/QueryTableRefresh"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
