package main

import (
	"context"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

//QueyRunner is query client for Elasticsearch
type QueryRunner struct {
	client *elastic.Client
	ctx    context.Context
}

//RunQuery query to Elasticsearch
func (q QueryRunner) RunQuery(query string) {
	queryStringQuery := elastic.NewQueryStringQuery(query)
	searchResult, err := q.client.Search().Index("items").Query(queryStringQuery).Do(q.ctx)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Hist: %d \n", searchResult.TotalHits())
}
