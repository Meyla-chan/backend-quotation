package handlers

import (
	"backend-quotation/config"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func setupTestDB(t *testing.T) (sqlmock.Sqlmock, *gin.Engine) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Gagal membuat mock: %s", err)
	}

	dialector := postgres.New(postgres.Config{Conn: dbMock})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	config.DB = db

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return mock, r
}


func TestGetBankAccounts(t *testing.T) {
	mock, r := setupTestDB(t)
	r.GET("/bank-accounts", GetBankAccounts)

	rows := sqlmock.NewRows([]string{"id", "bank_name"}).AddRow(1, "BCA")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bank_accounts"`)).WillReturnRows(rows)

	req, _ := http.NewRequest("GET", "/bank-accounts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBankAccountByID(t *testing.T) {
	mock, r := setupTestDB(t)
	r.GET("/bank-accounts/:id", GetBankAccountByID)

	rows := sqlmock.NewRows([]string{"id", "bank_name"}).AddRow(1, "BCA")
	
	
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bank_accounts" WHERE "bank_accounts"."id" = $1 ORDER BY "bank_accounts"."id" LIMIT $2`)).
		WithArgs("1", 1). // <-- Tambahkan angka 1 di sini untuk Limit
		WillReturnRows(rows)

	req, _ := http.NewRequest("GET", "/bank-accounts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBankAccount(t *testing.T) {
	mock, r := setupTestDB(t)
	r.POST("/bank-accounts", CreateBankAccount)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "bank_accounts"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	body := map[string]interface{}{"bank_name": "Mandiri", "account_name": "Melania", "account_number": "123", "company_id": 8}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/bank-accounts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateBankAccount(t *testing.T) {
	mock, r := setupTestDB(t)
	r.PUT("/bank-accounts/:id", UpdateBankAccount)

	
	rows := sqlmock.NewRows([]string{"id", "bank_name"}).AddRow(1, "BCA")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bank_accounts"`)).WillReturnRows(rows)

	
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "bank_accounts"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	body := map[string]interface{}{"bank_name": "BCA Digital"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/bank-accounts/1", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBankAccount(t *testing.T) {
	mock, r := setupTestDB(t)
	r.DELETE("/bank-accounts/:id", DeleteBankAccount)

	
	rows := sqlmock.NewRows([]string{"id", "bank_name"}).AddRow(1, "BCA")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bank_accounts"`)).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bank_accounts"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	req, _ := http.NewRequest("DELETE", "/bank-accounts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}