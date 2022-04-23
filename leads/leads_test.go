package leads

// import (
// 	"cdb"
// 	"log"
// 	"testing"
// 	"unicode"
// )

// const dsn = "root:password@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

// func BenchmarkLeads(b *testing.B) {
// 	db, err := Init(dsn)
// 	if err != nil {
// 		log.Fatalf("Cant init data base error: %s", err.Error())
// 	}
// 	l := &Leads{DB: db.DB}
// 	for i := 0; i < b.N; i++ {
// 		_, err := l.Leads(50, 0, "")
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
