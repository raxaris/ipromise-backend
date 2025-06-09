package main

import (
	"fmt"
	"github.com/raxaris/ipromise-backend/cmd/test/testdata"
	"log"

	"github.com/raxaris/ipromise-backend/config"
)

func main() {
	db := config.ConnectDB()

	err := testdata.LoadTestPostTree(db)
	if err != nil {
		log.Fatal("❌ Failed to load test data:", err)
	}

	fmt.Println("✅ Test data successfully loaded!")
}
