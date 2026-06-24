package main

import (
	"log"
	"gorm.io/driver/postgres" // Or mysql, depending on your setup
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "query",          // Where the query code will live
		ModelPkgPath: "models",         // Where your existing structs live
		Mode:         gen.WithDefaultQuery | gen.WithGeneric,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	// 4. Run it
	g.Execute()
}
