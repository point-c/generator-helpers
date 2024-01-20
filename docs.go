//go:build docs

package generator_helpers

//go:generate rm -rf docs
//go:generate go run "github.com/johnstarich/go/gopages" -base /generator-helpers -internal -out docs -source-link "https://github.com/point-c/generator-helpers/blob/main/{{.Path}}{{if .Line}}#L{{.Line}}{{end}}"

import _ "github.com/johnstarich/go/gopages"
