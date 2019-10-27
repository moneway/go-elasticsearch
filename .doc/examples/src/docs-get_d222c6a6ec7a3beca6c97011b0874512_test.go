// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L275>
//
// --------------------------------------------------------------------------------
// GET twitter/_source/1/?_source_includes=*.id&_source_excludes=entities
// --------------------------------------------------------------------------------

func Test_docs_get_d222c6a6ec7a3beca6c97011b0874512(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d222c6a6ec7a3beca6c97011b0874512[]
	res, err := es.GetSource(
		"twitter",
		"1",
		es.GetSource.WithSourceExcludes("entities"),
		es.GetSource.WithSourceIncludes("*.id"),
		es.GetSource.WithPretty(),
	)
	// end:d222c6a6ec7a3beca6c97011b0874512[]
	if err != nil {
		fmt.Println("Error getting the response:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Println(res)
}
