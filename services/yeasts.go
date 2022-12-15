package services

import (
	"brewess2/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllYeasts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var yeasts = models.GetYeasts()

	sqlStmt := `SELECT * FROM yeasts`
	rows, err := dbconn.Queryx(sqlStmt)

	if err == nil {
		var tempYeast = models.GetYeast()

		for rows.Next() {
			err = rows.StructScan(&tempYeast)
			yeasts = append(yeasts, tempYeast)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("No rows returned.")
				http.Error(w, err.Error(), 204)
			}
		case nil:
			json.NewEncoder(w).Encode(&yeasts)
		default:
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetYeast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var searchyeast = models.GetYeast()

	sqlStmt := `SELECT * FROM yeasts WHERE id=$1`
	row := dbconn.QueryRowx(sqlStmt, id)
	switch err := row.StructScan(&searchyeast); err {
	case sql.ErrNoRows:
		{
			log.Println("No rows returned.")
			http.Error(w, err.Error(), 204)
		}
	case nil:
		json.NewEncoder(w).Encode(&searchyeast)
	default:
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateYeast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var yeast = models.GetYeast()
	var id int

	_ = json.NewDecoder(r.Body).Decode(&yeast)

	sqlStmt := `INSERT INTO yeasts(title, top) VALUES($1,$2) RETURNING id`
	err := dbconn.QueryRow(sqlStmt, yeast.Title, yeast.Top).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	yeast.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(w).Encode(&yeast)
}

func UpdateYeast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var yeast = models.GetYeast()
	_ = json.NewDecoder(r.Body).Decode(&yeast)
	yeast.ID, _ = strconv.Atoi(params["id"])

	id := 0
	sqlStmt := `UPDATE yeasts SET title=$1, top=$2 WHERE id=$3 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, yeast.Title, yeast.Top, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("Updated record ID is:", id)
	json.NewEncoder(w).Encode(&yeast)
}

func DeleteYeast(w http.ResponseWriter, r *http.Request) {
	DeleteByID(w, r, "yeasts")
}
