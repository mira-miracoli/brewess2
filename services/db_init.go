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
	malts := TableCreator(
		"malts",
		[]*models.Column{
			ColumnCreator("tiltle", "text"),
			ColumnCreator("ebc", "integer"),
			ColumnCreator("amount", "integer"),
		},
	)
	malts.CreateTable()
	hops := TableCreator(
		"hops",
		[]*models.Column{
			ColumnCreator("tiltle", "text"),
			ColumnCreator("iso", "real"),
			ColumnCreator("amount", "integer"),
		},
	)
	hops.CreateTable()
	yeast := TableCreator(
		"yeasts",
		[]*models.Column{
			ColumnCreator("tiltle", "text"),
			ColumnCreator("top", "boolean"),
			ColumnCreator("amount", "integer"),
		},
	)
	yeast.CreateTable()
	mash_steps := TableCreator(
		"mash_steps",
		[]*models.Column{
			ColumnCreator("time", "integer"),
			ColumnCreator("temperature", "real"),
		},
	)
	mash_steps.CreateTable()
	recipe := TableCreator(
		"recipe",
		[]*models.Column{
			ColumnCreator("cooking_time", "integer"),
			ColumnCreator("basic_info", "text"),
			ColumnCreator("malt_info", "text"),
			ColumnCreator("hop_info", "text"),
			ColumnCreator("mash_info", "text"),
			ColumnCreator("fermentation_info", "text"),
			ColumnCreator("ibu", "real"),
			ColumnCreator("ebc", "real"),
			ColumnCreator("og_target", "real"),
			ColumnCreator("cast_worth", "real"),
			ColumnCreator("cooking_time", "int"),
			ColumnCreator("sha", "real"),
		},
	)
	recipe.CreateTable()
	recipe_malt_relation := TableCreator(
		"recipe_malt_relation",
		[]*models.Column{
			ColumnCreator("recipe_id", "integer"),
			ColumnCreator("malt_id", "integer"),
			ColumnCreator("proportion", "real"),
		},
	)
	recipe_malt_relation.CreateTable()
	recipe_hop_relation := TableCreator(
		"recipe_malt_relation",
		[]*models.Column{
			ColumnCreator("recipe_id", "integer"),
			ColumnCreator("hop_id", "integer"),
			ColumnCreator("proportion", "real"),
			ColumnCreator("cooking_time", "integer"),
		},
	)
	recipe_hop_relation.CreateTable()
	recipe_yeast_relation := TableCreator(
		"recipe_yeast_relation",
		[]*models.Column{
			ColumnCreator("recipe_id", "integer"),
			ColumnCreator("yeast_id", "integer"),
			ColumnCreator("proportion", "real"),
		},
	)
	recipe_yeast_relation.CreateTable()
	recipe_mash_step_relation := TableCreator(
		"recipe_malt_relation",
		[]*models.Column{
			ColumnCreator("recipe_id", "integer"),
			ColumnCreator("mash_step_id", "integer"),
		},
	)
	recipe_mash_step_relation.CreateTable()
}
