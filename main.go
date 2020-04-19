package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Send the Headers back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(r.Header)
	if err != nil{
		fmt.Printf("%v", err)
		return
	}
}

func main(){
	if len(os.Args) < 2{
		fmt.Println("Port number required to start proxy gateway")
		os.Exit(1)
	}

	// Get the port from the command line
	port := os.Args[1]
	port = strings.TrimRight(port, "\r\n")

	// Check if port is valid integer
	// Checking if the port is in correct range and other conditions are left out intentionally
	if _, err := strconv.Atoi(port); err != nil{
		fmt.Println("Invalid port specified, received error: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Proxy Gateway Started on port: " + port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":" + port, nil)

	// The line of execution will not come here, if it does then its an error
	fmt.Println(err)
}