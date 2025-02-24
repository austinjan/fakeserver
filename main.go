package main

import (
	"encoding/json"
	"fakeserver/data"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var employees []data.Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Handle query parameters
	query := r.URL.Query()
	department := query.Get("department")
	jobTitle := query.Get("job_title")
	name := query.Get("name")
	yearsOfService := query.Get("years_of_service")
	salary := query.Get("salary")

	filteredEmployees := make([]data.Employee, 0)
	for _, emp := range employees {
		// Apply filters
		if department != "" && !strings.EqualFold(emp.Department, department) {
			continue
		}
		if jobTitle != "" && !strings.EqualFold(emp.JobTitle, jobTitle) {
			continue
		}
		if name != "" && !strings.Contains(strings.ToLower(emp.Name), strings.ToLower(name)) {
			continue
		}

		// Handle years of service filter
		if yearsOfService != "" {
			hireYear := strings.Split(emp.HireDate, "-")[0]
			currentYear := time.Now().Year()
			empYears, _ := strconv.Atoi(hireYear)
			years := currentYear - empYears

			if !compareValues(years, yearsOfService) {
				continue
			}
		}

		// Handle salary filter
		if salary != "" {
			if !compareValues(emp.Salary, salary) {
				continue
			}
		}

		filteredEmployees = append(filteredEmployees, emp)
	}

	json.NewEncoder(w).Encode(filteredEmployees)
}

// Helper function to compare values with operators
func compareValues(value interface{}, condition string) bool {
	operators := []string{">", "<", ">=", "<=", "="}
	var op string
	var numStr string

	for _, operator := range operators {
		if strings.HasPrefix(condition, operator) {
			op = operator
			numStr = strings.TrimSpace(condition[len(operator):])
			break
		}
	}

	if op == "" || numStr == "" {
		return false
	}

	targetNum, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return false
	}

	var valueNum float64
	switch v := value.(type) {
	case int:
		valueNum = float64(v)
	case float64:
		valueNum = v
	default:
		return false
	}

	switch op {
	case ">":
		return valueNum > targetNum
	case "<":
		return valueNum < targetNum
	case ">=":
		return valueNum >= targetNum
	case "<=":
		return valueNum <= targetNum
	case "=":
		return valueNum == targetNum
	default:
		return false
	}
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID format"})
		return
	}

	for _, emp := range employees {
		if emp.ID == id {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})
}

func getEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dept := params["department"]

	filteredEmployees := make([]data.Employee, 0)
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, dept) {
			filteredEmployees = append(filteredEmployees, emp)
		}
	}

	json.NewEncoder(w).Encode(filteredEmployees)
}

func main() {
	// Generate and save employee data
	employees = data.GenerateEmployees(50)
	err := data.SaveEmployeesToFile(employees, "employees.json")
	if err != nil {
		log.Fatal("Error saving employees data:", err)
	}

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", getEmployeeByID).Methods("GET")
	r.HandleFunc("/api/employees/department/{department}", getEmployeesByDepartment).Methods("GET")

	log.Printf("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
