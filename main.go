package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"udacity-crm-backend/model"
)

var customers = []*model.Customer{
	{
		Id:        1,
		Name:      "Jynn",
		Role:      "Price analyst",
		Email:     "jynn@smu.com",
		Phone:     "9923 4342",
		Contacted: false,
	},
	{
		Id:        2,
		Name:      "Charlie",
		Role:      "Backend developer",
		Email:     "charlie@bytedance.com",
		Phone:     "9923 1232",
		Contacted: true,
	},
	{
		Id:        3,
		Name:      "Gloria",
		Role:      "Evangelist",
		Email:     "gloria96@smu.com",
		Phone:     "9432 4342",
		Contacted: false,
	},
}

var nextID = 4

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "http server provide backend api for CRM")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var idStr string
	var ok bool

	if idStr, ok = mux.Vars(r)["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Search for customer with input ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("error converting id param")
		return
	}
	for _, customer := range customers {
		if customer.Id == id {
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "no customer with id=%d exists", id)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req map[string]string
	json.Unmarshal(reqBody, &req)
	newCustomer := &model.Customer{
		Id:        nextID,
		Name:      req["name"],
		Role:      req["role"],
		Email:     req["email"],
		Phone:     req["phone"],
		Contacted: false,
	}
	customers = append(customers, newCustomer)
	nextID += 1

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Search for customer with inout ID
	var idStr string
	var ok bool
	if idStr, ok = mux.Vars(r)["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(idStr)
	customer := getCustomerByID(id)

	// Update fields
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req map[string]string
	json.Unmarshal(reqBody, &req)

	if newName, ok := req["name"]; ok {
		customer.Name = newName
	}
	if newRole, ok := req["role"]; ok {
		customer.Role = newRole
	}
	if newEmail, ok := req["email"]; ok {
		customer.Email = newEmail
	}
	if newPhone, ok := req["phone"]; ok {
		customer.Phone = newPhone
	}
	if newContacted, ok := req["contacted"]; ok && newContacted == "true" {
		customer.Contacted = true
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var idStr string
	var ok bool
	if idStr, ok = mux.Vars(r)["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(idStr)
	deleteCustomerByID(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func getCustomerByID(id int) *model.Customer {
	for _, customer := range customers {
		if customer.Id == id {
			return customer
		}
	}
	return nil
}

func deleteCustomerByID(id int) {
	var updatedCustomers []*model.Customer

	for _, customer := range customers {
		if customer.Id != id {
			updatedCustomers = append(updatedCustomers, customer)
		}
	}
	customers = updatedCustomers
}

func main() {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("http server running at port 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("error running http server")
	}
}
