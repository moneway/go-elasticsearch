// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L320>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/1
// {
//     "counter" : 1,
//     "tags" : ["red"]
// }
// --------------------------------------------------------------------------------

func Test_docs_get_5eabcdbf61bfcb484dc694f25c2bba36(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5eabcdbf61bfcb484dc694f25c2bba36[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "counter": 1,
		  "tags": [
		    "red"
		  ]
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithPretty(),
	)
	// end:5eabcdbf61bfcb484dc694f25c2bba36[]
	if err != nil {
		fmt.Println("Error getting the response:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Println(res)
}
