package main

import (
	"fmt"
	"log"

	"net/http"

	"encoding/json"


	"github.com/jinzhu/gorm"


	_ "github.com/jinzhu/gorm/dialects/postgres"

	"golang.org/x/crypto/bcrypt"
)

type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

type RequestError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}



var healthStauts = HealthCheck{Version: "1.0", Status: "ok"}

var db *gorm.DB
var err error

func main() {

	//set api version 
	
	//start mux
	//start db
	////register routes passing db instance 
	//init cors middlaware 
	//listen 
	

}

func checkApiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStauts)
}






}




