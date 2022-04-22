package cdb

import (
	"log"
	"testing"
)

func BenchmarkLeads(b *testing.B) {
	if err := Init(dsn); err != nil {
		log.Fatalf("Cant init data base error: %s", err.Error())
	}
	for i := 0; i < b.N; i++ {
		_, err := Leads(50, 0, "")
		if err != nil {
			panic(err)
		}
	}
}
