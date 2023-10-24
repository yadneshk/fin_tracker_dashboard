package main

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

// Data structure for the database query result
type ChartData struct {
    Date  string  `json:"date"`
    Value float32 `json:"amount"`
}

func main() {
    // Initialize your database connection

    // Initialize a router (e.g., Gorilla Mux)
    router := mux.NewRouter()

    // Define an API endpoint to fetch chart data
    router.HandleFunc("/api/data", getChartData).Methods("GET")

    // Serve your web page and the Chart.js code
    http.Handle("/", router)
    http.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getChartData(w http.ResponseWriter, r *http.Request) {
    // Connect to your database
    db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/fin_tracker")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Execute a query to retrieve data from the database
    rows, err := db.Query("select txn_date, debit from expenses where debit is not null")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Iterate over the results and build a slice of ChartData
    var chartData []ChartData
    for rows.Next() {
        var data ChartData
        if err := rows.Scan(&data.Date, &data.Value); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        chartData = append(chartData, data)
    }

    // Serialize the data to JSON
    jsonData, err := json.Marshal(chartData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
