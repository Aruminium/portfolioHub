package utils

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// RespondWithError エラー情報をJSONで返す
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:passw0rd@/portfolio_hub?charset=utf8&parseTime=True&loc=Local")
	// 接続に失敗したらエラーログを出して終了する
	if err != nil {
			log.Fatalf("DB connection failed %v", err)
	}
	db.LogMode(true)

	return db
}

func GetID(r *http.Request) (id string) {
	vars := mux.Vars(r)
	return vars["id"]
}