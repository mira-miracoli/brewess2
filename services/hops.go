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

func GetAllHops(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var hops = models.GetHops()

	sqlStmt := `SELECT * FROM hops`
	rows, err := dbconn.Queryx(sqlStmt)

	if err == nil {
		var tempHop = models.GetHop()

		for rows.Next() {
			err = rows.StructScan(&tempHop)
			hops = append(hops, tempHop)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("No rows returned.")
				http.Error(w, err.Error(), 204)
			}
		case nil:
			json.NewEncoder(w).Encode(&hops)
		default:
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetHop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var searchhop = models.GetHop()

	sqlStmt := `SELECT * FROM hops WHERE id=$1`
	row := dbconn.QueryRowx(sqlStmt, id)
	switch err := row.StructScan(&searchhop); err {
	case sql.ErrNoRows:
		{
			log.Println("No rows returned.")
			http.Error(w, err.Error(), 204)
		}
	case nil:
		json.NewEncoder(w).Encode(&searchhop)
	default:
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateHop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var hop = models.GetHop()
	var id int

	_ = json.NewDecoder(r.Body).Decode(&hop)

	sqlStmt := `INSERT INTO hops(title, iso, amount) VALUES($1,$2,$3) RETURNING id`
	err := dbconn.QueryRow(sqlStmt, hop.Title, hop.Iso, hop.Amount).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	hop.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(w).Encode(&hop)
}

func UpdateHop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var hop = models.GetHop()
	_ = json.NewDecoder(r.Body).Decode(&hop)
	hop.ID, _ = strconv.Atoi(params["id"])

	id := 0
	sqlStmt := `UPDATE hops SET title=$1, iso=$2, amount=$3 WHERE id=$4 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, hop.Title, hop.Iso, hop.Amount, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("Updated record ID is:", id)
	json.NewEncoder(w).Encode(&hop)
}

func DeleteHop(w http.ResponseWriter, r *http.Request) {
	DeleteByID(w, r, "hops")
}
