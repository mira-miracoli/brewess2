package tests

import (
	"brewess2/services"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetAllMalts(t *testing.T) {
	db, mock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	services.SetDB(db)
	mock.ExpectQuery("SELECT * FROM malts;").
		WillReturnRows(sqlxmock.NewRows([]string{"id", "title", "ebc", "amount"}).
			AddRow("1", "Münchner", "4", 3400))
	r, err := http.NewRequest("GET", "/malt", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	services.GetAllMalts(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	returned := string(body)
	expected := `[{"id":1,"title":"Münchner","EBC":"4","amount":3400}]` + "\n"
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, returned)
}

func TestGetMalt(t *testing.T) {
	db, mock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	services.SetDB(db)
	mock.ExpectQuery("SELECT * FROM malts WHERE id=$1").
		WillReturnRows(sqlxmock.NewRows([]string{"id", "title", "ebc", "amount"}).
			AddRow("1", "Münchner", "4", 3400))
	r, err := http.NewRequest("GET", "/malts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	services.GetMalt(w, r)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	returned := string(body)
	expected := `{"id":1,"title":"Münchner","EBC":"4","amount":3400}` + "\n"
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, returned)
}
