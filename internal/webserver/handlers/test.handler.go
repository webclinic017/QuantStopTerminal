package handlers

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
)

func Test(db *sql.DB, usr *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	//log.Println("all")
	return write.Success()
}
