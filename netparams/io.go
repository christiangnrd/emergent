// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netparams

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"cogentcore.org/core/gi"
	"cogentcore.org/core/glop/indent"
	"cogentcore.org/core/grows"
	"cogentcore.org/core/grows/jsons"
	"cogentcore.org/core/grows/tomls"
	"github.com/BurntSushi/toml"
	"github.com/emer/emergent/v2/params"
)

// WriteGoPrelude writes the start of a go file in package main that starts a
// variable assignment to given variable -- for start of SaveGoCode methods.
func WriteGoPrelude(w io.Writer, varNm string) {
	w.Write([]byte("// File generated by netparams.SaveGoCode\n\n"))
	w.Write([]byte("package main\n\n"))
	w.Write([]byte(`import "github.com/emer/emergent/v2/params"`))
	w.Write([]byte(`import "github.com/emer/emergent/v2/netparams"`))
	w.Write([]byte("\n\nvar " + varNm + " = "))
}

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sets) OpenJSON(filename gi.Filename) error {
	*pr = make(Sets) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sets) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Sets) OpenTOML(filename gi.Filename) error {
	*pr = make(Sets) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Sets) SaveTOML(filename gi.Filename) error {
	// return tomls.Save(pr, string(filename)) // pelletier/go-toml produces bad output on maps
	return grows.Save(pr, string(filename), func(w io.Writer) grows.Encoder {
		return toml.NewEncoder(w)
	})
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sets) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("netparams.Sets{\n"))
	depth++
	for nm, st := range *pr {
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(`"` + nm + `": `))
		st.WriteGoCode(w, depth)
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}\n"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sets) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sets) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	params.WriteGoPrelude(fp, "SavedParamsSets")
	pr.WriteGoCode(fp, 0)
	return nil
}

/*
var SetsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsAll", ki.Props{
			"desc":        "between all sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsFirst", ki.Props{
			"desc":        "between first set (e.g., the Base set) and rest of sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports all the cases where the same param path is being set to different values within different sheets in given set",
			"icon":        "search",
			"show-return": true,
			"Args": ki.PropSlice{
				{"Set Name", ki.Props{}},
			},
		}},
	},
}
*/
