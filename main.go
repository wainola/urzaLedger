package main

import (
	"fmt"
	"net/http"
	"urza/models"
	"urza/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	appDB := &models.UrzaDB{DB: utils.InstanceDbConnection()}

	defer appDB.DB.Close()

	r := buildRouter(appDB)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
