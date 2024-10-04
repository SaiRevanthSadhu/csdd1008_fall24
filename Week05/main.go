package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Car struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var (
	cars   = make(map[int]Car)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/cars", carsHandler)
	http.HandleFunc("/cars/", carHandler) // For specific car actions
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func carsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCars(w)
	case http.MethodPost:
		createCar(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func carHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getCar(w, id)
	case http.MethodPut:
		updateCar(w, r, id)
	case http.MethodDelete:
		deleteCar(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getCars(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()

	carList := make([]Car, 0, len(cars))
	for _, car := range cars {
		carList = append(carList, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carList)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	car.ID = nextID
	nextID++
	cars[car.ID] = car
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	car, found := cars[id]
	if !found {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	car, found := cars[id]
	if !found {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cars[id] = car
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	if _, found := cars[id]; !found {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	delete(cars, id)
	w.WriteHeader(http.StatusNoContent)
}

func parseID(path string) (int, error) {
	var id int
	_, err := fmt.Sscanf(path, "/cars/%d", &id)
	if err != nil {
		return 0, fmt.Errorf("invalid car ID")
	}
	return id, nil
}
