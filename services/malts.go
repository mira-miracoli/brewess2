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

func GetAllMalts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var malts = models.GetMalts()

	sqlStmt := `SELECT * FROM malts`
	rows, err := dbconn.Queryx(sqlStmt)

	if err == nil {
		var tempMalt = models.GetMalt()

		for rows.Next() {
			err = rows.StructScan(&tempMalt)
			malts = append(malts, tempMalt)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("No rows returned.")
				http.Error(w, err.Error(), 204)
			}
		case nil:
			json.NewEncoder(w).Encode(&malts)
		default:
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetMalt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var searchmalt = models.GetMalt()

	sqlStmt := `SELECT * FROM malts WHERE id=$1`
	row := dbconn.QueryRowx(sqlStmt, id)
	switch err := row.StructScan(&searchmalt); err {
	case sql.ErrNoRows:
		{
			log.Println("No rows returned.")
			http.Error(w, err.Error(), 204)
		}
	case nil:
		json.NewEncoder(w).Encode(&searchmalt)
	default:
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateMalt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var malt = models.GetMalt()
	var id int

	_ = json.NewDecoder(r.Body).Decode(&malt)

	sqlStmt := `INSERT INTO malts(title, EBC) VALUES($1,$2) RETURNING id`
	err := dbconn.QueryRow(sqlStmt, malt.Title, malt.EBC).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	malt.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(w).Encode(&malt)
}

func UpdateMalt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var malt = models.GetMalt()
	_ = json.NewDecoder(r.Body).Decode(&malt)
	malt.ID, _ = strconv.Atoi(params["id"])

	id := 0
	sqlStmt := `UPDATE malts SET title=$1, EBC=$2 WHERE id=$3 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, malt.Title, malt.EBC, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("Updated record ID is:", id)
	json.NewEncoder(w).Encode(&malt)
}

func DeleteMalt(w http.ResponseWriter, r *http.Request) {
	DeleteByID(w, r, "malts")
}
