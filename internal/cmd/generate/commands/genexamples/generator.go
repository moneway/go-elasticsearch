// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package genexamples

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// SrcGenerator represents the generator for Go source files.
//
type SrcGenerator struct {
	b       bytes.Buffer
	Example Example
}

// DocGenerator represents the generator for Go source files.
//
type DocGenerator struct {
	b       bytes.Buffer
	Example Example
}

func (g SrcGenerator) Filename() string {
	return fmt.Sprintf(
		"%s_%s_test.go",
		strings.ReplaceAll(strings.TrimSuffix(g.Example.SourceLocation.File, ".asciidoc"), "/", "-"),
		g.Example.Digest)
}

func (g SrcGenerator) Output() io.Reader {
	var out bytes.Buffer

	out.WriteString(`package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

`)

	out.WriteString(fmt.Sprintf("// Generated from %s\n//\n", g.Example.GithubURL()))

	out.WriteString(fmt.Sprintf(`func Test__%s(t *testing.T) {`+"\n", g.Example.Digest))
	out.WriteString("\t")
	if !g.Example.IsTranslated() {
		out.WriteString("// ")
	}
	out.WriteString(`es, _ := elasticsearch.NewDefaultClient()` + "\n")

	if !g.Example.IsTranslated() {
		out.WriteString("\n\tt.Error(`")

		out.WriteString(strings.Split(g.Example.Source, "\n")[0])
		out.WriteString("`)\n")
		out.WriteString("\t// " + strings.Repeat("-", 80) + "\n")
		for _, l := range strings.Split(g.Example.Source, "\n") {
			out.WriteString("\t// " + l + "\n")
		}
		out.WriteString("\t//" + strings.Repeat("-", 80) + "\n")
	} else {
		src, err := g.Example.Translated()
		if err != nil {
			out.WriteString("\t" + `panic("Error translating example")`)
		} else {
			out.WriteString(src)
		}
	}

	out.WriteString("}\n")

	return &out
}