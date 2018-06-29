package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/mem"
)

// Data to export
type Data struct {
	TYPE string `json:"type,omitempty"`
	Data string `json:"data,omitempty"`
}

func getCPU(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	result := fmt.Sprintf("Total: %v, UsedPercent:%f%%\n", v.Total, v.UsedPercent)
	fmt.Printf("\n%s\n", time.Now().String())
	fmt.Println(v)
	json.NewEncoder(w).Encode(Data{TYPE: "1", Data: result})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info/cpu", getCPU).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
