package es_queries

import "github.com/olivere/elastic"

func (eq ESQueries) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	// Create empty slice of needed query
	// []
	equalsQueries := make([]elastic.Query, 0)
	for _, q := range eq.Equals {
		// Append the object of query to slice
		// [{"field":"value"}]
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(q.Field, q.Value))
	}
	// The sliced of query object reserved all the query
	query.Must(equalsQueries...)
	return query
}
