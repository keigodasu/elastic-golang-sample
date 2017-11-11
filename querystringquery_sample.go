package main

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
	}

	runner := QueryRunner{client: client, ctx: ctx}

	// xx AND xx
	runner.RunQuery("priority: 1 AND CreatedAt:2017-11-06")

	// xx OR xx
	runner.RunQuery("priority: 1 OR CreatedAt:2017-11-06")

	// NOT xx
	runner.RunQuery("priority: 1 OR NOT CreatedAt:2017-11-06")

	// (xx AND xx) OR (xx AND xx)
	runner.RunQuery("(priority:1 AND CreatedAt:2017-11-06) OR (CreatedAt:2017-11-03)")

	// range
	runner.RunQuery("CreatedAt: [2017-11-01 TO 2017-11-04]")
	runner.RunQuery("CreatedAt:>=2017-11-01 AND CreatedAt:<=2017-11-04")
	runner.RunQuery("CreatedAt:>2017-11-04")

	// wildcards
	runner.RunQuery("comment:moga*")
	runner.RunQuery("comment:ugau?a")

	//regular
	runner.RunQuery("comment:/hogeho[gu]e/")
}
