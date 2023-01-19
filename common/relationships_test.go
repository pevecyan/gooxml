// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/pevecyan/gooxml/common"
	"github.com/pevecyan/gooxml/testhelper"
	"github.com/pevecyan/gooxml/zippkg"
)

func TestRelationshipsUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/rels.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := common.NewRelationships()
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r.X()); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "rels.xml", got.Bytes())
}
