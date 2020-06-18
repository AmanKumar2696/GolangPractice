package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Account struct {
	AccountNumber int     `json:"AccountNumber"`
	AccountType   string  `json:"AccountType"`
	Name          string  `json:"Name"`
	Gender        string  `json:"Gender"`
	Contact       string  `json:"Contact"`
	Balance       float32 `json:"Balance"`
}

func dbConnect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Database Connection")
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success")
	fmt.Fprintf(w, "<h1>Success</h1>")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) { //Insert new value
	jsonBody, err := ioutil.ReadAll(r.Body) //Read the body of the request
	if err != nil {
		fmt.Fprintf(w, "Error in reading request")
		return
	}
	account := Account{}                     //create instance of structure
	err = json.Unmarshal(jsonBody, &account) //parse JSON to Structure, passed by reference
	fmt.Println(account.AccountNumber, account.AccountType, account.Name,
		account.Gender, account.Contact, account.Balance)
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp") //Connect to database
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() //close database object before the function returns
	fmt.Println("Success")
	fmt.Fprintf(w, "<h1>Success</h1>")
	insert, err2 := db.Query("INSERT INTO Account VALUES(?,?,?,?,?,?)", account.AccountNumber, account.AccountType, account.Name,
		account.Gender, account.Contact, account.Balance)
	if err2 != nil {
		panic(err2.Error())
	}
	defer insert.Close()
}

func ReadAll(w http.ResponseWriter, r *http.Request) { //Read All
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp") //Connect to database
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("Select * from Account")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var account Account
		err = result.Scan(&account.AccountNumber, &account.AccountType, &account.Name, &account.Gender, &account.Contact,
			&account.Balance)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(account)
		json.NewEncoder(w).Encode(account)
	}
}

func ReadOne(w http.ResponseWriter, r *http.Request) { //Read One
	query := mux.Vars(r)["AccountNumber"]
	accNum, err := strconv.Atoi(query)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var account Account
	err = db.QueryRow("SELECT * FROM Account where AccountNumber = ?", accNum).Scan(&account.AccountNumber, &account.AccountType,
		&account.Name, &account.Gender, &account.Contact, &account.Balance)
	json.NewEncoder(w).Encode(account)

}

func UpdateNum(w http.ResponseWriter, r *http.Request) { //Update Contact Number
	query := strings.Split(mux.Vars(r)["Contact"], ",")
	contact := query[0]
	acc := query[1]
	accNum, err3 := strconv.Atoi(acc)
	if err3 != nil {
		fmt.Println("Error in Conversion", err3)
	}
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	update, err2 := db.Query("UPDATE Account SET Contact = ? WHERE AccountNumber=?", contact, accNum)
	if err2 != nil {
		panic(err2.Error())
	}
	defer update.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) { //Delete
	query := mux.Vars(r)["AccountNumber"]
	accNum, err := strconv.Atoi(query)
	db, err2 := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp")
	if err != nil {
		panic(err2.Error())
	}
	defer db.Close()
	del, err3 := db.Query("DELETE FROM Account WHERE AccountNumber=?", accNum)
	if err3 != nil {
		panic(err3.Error())
	}
	defer del.Close()
}
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/insert", CreateAccount).Methods("POST") //method only accessible by POST request
	router.HandleFunc("/showall", ReadAll).Methods("GET")
	router.HandleFunc("/showone/{AccountNumber}", ReadOne).Methods("GET")
	router.HandleFunc("/updatenum/{Contact}", UpdateNum).Methods("PUT")
	router.HandleFunc("/delete/{AccountNumber}", Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
