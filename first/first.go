package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type ExpensesResponse struct {
	Id          int
	Date        string
	Description string
	Amount      int
}

type ErrorResponce struct {
	Message string
	Code    string
}
type ExpensesCreateRequest struct {
	Id          int
	Date        string
	Description string
	Amount      int
}

func expenses(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), "select id, date::varchar, description, amount from expenses")
	if err != nil {
		fmt.Fprint(w, "yalnyslyk: ", err.Error())
		return
	}
	var list []ExpensesResponse
	for rows.Next() {
		var res ExpensesResponse
		err = rows.Scan(&res.Id, &res.Date, &res.Description, &res.Amount)
		if err != nil {
			fmt.Fprint(w, "Yalnyslyk: ", err.Error())
		}
		list = append(list, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func ExpensesAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		json.NewEncoder(w).Encode(ErrorResponce{"only works PUT method", "400"})
		return
	}
	var req ExpensesCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	idStr := strconv.Itoa(req.Id)
	// dateStr := strconv.Itoa(req.Date)
	amountStr := strconv.Itoa(req.Amount)
	_, err = db.Exec(context.Background(), "insert into expenses (Id, Date, Description, Amount)values ('"+idStr+"', '"+req.Date+"', '"+req.Description+"', '"+amountStr+"');")
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	json.NewEncoder(w).Encode(true)
}
func deleteexpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		json.NewEncoder(w).Encode(ErrorResponce{"only works delete method", "400"})
		return
	}
	var req ExpensesResponse
	// json.NewDecoder(r.Body).Decode(&req)
	id := r.FormValue("id")
	fmt.Println(req.Id, id)
	_, err := db.Exec(context.Background(), "delete from expenses where id=$1", id)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	json.NewEncoder(w).Encode(true)
}
func summaryexpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-Type", "application/json")
	var sum int
	rows := db.QueryRow(context.Background(), "select sum(amount) from expenses")
	// if err != nil {
	// 	json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
	// 	return
	// }
	err := rows.Scan(&sum)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	json.NewEncoder(w).Encode(sum)
}
func getexpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(ErrorResponce{"only works get method", "400"})
		return
	}
	var req ExpensesResponse
	id := r.FormValue("id")
	// decoder:=json.NewDecoder(r.Body)
	// err:=decoder.Decode(&req)
	// if err!=nil{
	// 	json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
	// }
	// idstr:=strconv.Itoa(req.Id)
	// id:=r.FormValue("id")
	rows := db.QueryRow(context.Background(), "select  id, date::varchar, description, amount from expenses where id=$1", id)

	err := rows.Scan(&req.Id, &req.Date, &req.Description, &req.Amount)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	json.NewEncoder(w).Encode(req)

}
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		json.NewEncoder(w).Encode(ErrorResponce{"only works PUT method", "400"})
		return
	}
	var req ExpensesCreateRequest
	// json.NewDecoder(r.Body).Decode(&req)

	// if err != nil {
	// 	json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
	// 	return
	// }
	// idStr := strconv.Itoa(req.Id)
	// // dateStr := strconv.Itoa(req.Date)
	// amountStr := strconv.Itoa(req.Amount)
	json.NewDecoder(r.Body).Decode(&req)
	_, err := db.Exec(context.Background(), "update expenses set date=$1, description=$2, amount=$3 where id=$4", &req.Date, &req.Description, &req.Amount, &req.Id)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponce{err.Error(), "400"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

var db *pgx.Conn

func connectDB(config string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect")
		os.Exit(1)
	}
	return conn
}
func main() {
	db = connectDB("postgres://postgres:@localhost:5432")
	defer db.Close(context.Background())
	http.HandleFunc("/expenses", expenses)
	http.HandleFunc("/expenses/add", ExpensesAdd)
	http.HandleFunc("/expenses/del", deleteexpenses)
	http.HandleFunc("/expenses/summary", summaryexpenses)
	http.HandleFunc("/getexpense", getexpense)
	http.HandleFunc("/updateexpense", update)

	http.ListenAndServe(":8080", nil)
}
