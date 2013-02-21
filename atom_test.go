package atom

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	f := Feed{
		Common: Common{
			Title: "test",
		},
		Entries: []Entry{
			{
				Common: Common{
					Title: "e1",
					Author: &Person{
						Name: "e1author",
					},
				},
				Content: Content{
					Type: "html",
					Data: "<h1>Hi</h1>",
				},
			},
		},
	}
	t.Logf(&f)
}
