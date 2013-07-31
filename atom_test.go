// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
