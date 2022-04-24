package cdb

import (
	"log"
	"testing"

	"github.com/hromov/cdb/leads"
	"github.com/hromov/cdb/models"
)

func BenchmarkLeads(b *testing.B) {
	db, err := Init(dsnForTests)
	if err != nil {
		log.Fatalf("Cant init data base error: %s", err.Error())
	}
	l := &leads.Leads{DB: db.DB}
	for i := 0; i < b.N; i++ {
		_, err := l.List(models.ListFilter{Limit: 50, Offset: 0, Query: ""})
		if err != nil {
			panic(err)
		}
	}
}
