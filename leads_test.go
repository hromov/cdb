package cdb

import (
	"log"
	"testing"

	"github.com/hromov/cdb/leads"
)

func BenchmarkLeads(b *testing.B) {
	db, err := Init(dsnForTests)
	if err != nil {
		log.Fatalf("Cant init data base error: %s", err.Error())
	}
	l := &leads.Leads{DB: db.DB}
	for i := 0; i < b.N; i++ {
		_, err := l.List(50, 0, "")
		if err != nil {
			panic(err)
		}
	}
}
