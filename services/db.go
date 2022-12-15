package services

import (
	"brewess2/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var dbconn *sqlx.DB

func SetDB(db *sqlx.DB) {
	dbconn = db
}

func DeleteByID(w http.ResponseWriter, r *http.Request, table string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := 0
	sqlStmt := `DELETE FROM` + ` table ` + `WHERE id=$1 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("Deleted record ID is:", id)
	json.NewEncoder(w).Encode(id)
}

func InitDB() {
	hops := TableCreator("hops", []*models.Column{ColumnCreator("iso", "real"), ColumnCreator("amount", "integer")})
	hops.CreateResourceTable()
	yeast := TableCreator("yeasts", []*models.Column{ColumnCreator("top", "boolean"), ColumnCreator("amount", "integer")})
	yeast.CreateResourceTable()
}
