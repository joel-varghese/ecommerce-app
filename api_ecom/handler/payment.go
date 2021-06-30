package handler

import (
	"api_ecom/model"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	uuid "github.com/uuid"
)

func MakePayment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	payment := model.Payment{}
	var price float64
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO payment(id,username,price,user_id) VALUES(?,?,?,?);")
	if err != nil {
		panic(err)
	}

	prc, err := db.QueryContext(ctx, "select price from cart where userid=?", payment.User_ID)
	if err != nil {
		panic(err)
	}
	uuid, _ := uuid.NewUUID()
	res, err := stmt.Exec(uuid, payment.Username,prc,payment.User_ID)

	if err != nil && res != nil {
		panic(err)
	}
	err := db.QueryRowContext(ctx, "DELETE FROM cart WHERE userid=?;", payment.User_ID).Scan(&cartID.CartID)
	respondJSON(w, http.StatusCreated, user)
}