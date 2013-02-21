// Package atom implements the syndication format Atom.
//
// http://tools.ietf.org/html/rfc4287
package atom

import (
	"bytes"
	"encoding/xml"
	"io"
)

const MIME = "application/atom+xml"

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	URI  string `xml:"href,attr"`
}

type Person struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri,omitempty"`
	Email string `xml:"email,omitempty"`
}

type Category struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr,omitempty"`
	Label  string `xml:"label,attr,omitempty"`
}

type Common struct {
	ID           string     `xml:"id,omitempty"`
	Title        string     `xml:"title"`
	Subtitle     string     `xml:"subtitle,omitempty"`
	Categories   []Category `xml:"category"`
	Author       *Person    `xml:"author"`
	Contributors []Person   `xml:"contributor"`
	Updated      string     `xml:"updated,omitempty"`
	Published    string     `xml:"published,omitempty"`
	Links        []Link     `xml:"link"`
	Rights       string     `xml:"rights,omitempty"`
}

type Content struct {
	Type string `xml:"type,attr,omitempty"`
	Data string `xml:",chardata"`
}

type Feed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Common
	Entries []Entry `xml:"entry"`
}

func (f *Feed) marshal(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "\t")
	return enc.Encode(f)
}

func (f *Feed) WriteTo(w io.Writer) (int64, error) {
	b := new(bytes.Buffer)
	if err := f.marshal(b); err != nil {
		return 0, err
	}
	return b.WriteTo(w)
}

func (f *Feed) String() string {
	b := new(bytes.Buffer)
	if err := f.marshal(b); err != nil {
		return err.Error()
	}
	return b.String()
}

type Entry struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom entry"`
	Common
	Source  string  `xml:"source,omitempty"`
	Content Content `xml:"content"`
}
