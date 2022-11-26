package main

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	id string
	name string
	photoUrl string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", findAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", findById).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	db := utils.GetConnection()
	defer db.Close()

	var userList []User
	db.Find(&userList)
	utils.RespondWithJSON(w, http.StatusOK, userList)
}

func findById(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	var user User
	// IDで検索しに行く
	db.Where("id = ?", id).Find(&user)
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// リクエストボディ取得
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "JSON Unmarshaling failed .")
		return
	}

	db := utils.GetConnection()
	defer db.Close()

	db.Create(&user)
	utils.RespondWithJSON(w, http.StatusOK, user)
}