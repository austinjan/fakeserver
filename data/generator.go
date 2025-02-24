package data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Employee struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Gender            string  `json:"gender"`
	Phone             string  `json:"phone"`
	Birthday          string  `json:"birthday"`
	Email             string  `json:"email"`
	JobTitle          string  `json:"job_title"`
	Department        string  `json:"department"`
	HireDate          string  `json:"hire_date"`
	Salary            float64 `json:"salary"`
	WorkStatus        string  `json:"work_status"`
	BankInfo          string  `json:"bank_info"`
	SupervisorID      int     `json:"supervisor_id"`
	TotalSpecialLeave int     `json:"total_special_leave"`
	UsedSpecialLeave  int     `json:"used_special_leave"`
}

var departments = []string{"Engineering", "Sales", "Marketing", "HR", "Finance", "Operations"}
var jobTitles = map[string][]string{
	"Engineering": {"Software Engineer", "Senior Engineer", "Tech Lead", "DevOps Engineer"},
	"Sales":       {"Sales Representative", "Sales Manager", "Account Executive"},
	"Marketing":   {"Marketing Specialist", "Marketing Manager", "Content Writer"},
	"HR":         {"HR Specialist", "HR Manager", "Recruiter"},
	"Finance":    {"Accountant", "Financial Analyst", "Finance Manager"},
	"Operations": {"Operations Manager", "Project Manager", "Business Analyst"},
}

var firstNames = []string{"John", "Jane", "Michael", "Emily", "David", "Sarah", "James", "Emma", "William", "Olivia"}
var lastNames = []string{"Smith", "Johnson", "Brown", "Davis", "Wilson", "Anderson", "Taylor", "Thomas", "Moore", "Martin"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateName() string {
	return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
}

func generatePhone() string {
	return fmt.Sprintf("+1-%d%d%d-%d%d%d-%d%d%d%d",
		rand.Intn(10), rand.Intn(10), rand.Intn(10),
		rand.Intn(10), rand.Intn(10), rand.Intn(10),
		rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))
}

func generateEmail(name string) string {
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com"}
	nameNormalized := strings.ToLower(strings.Replace(name, " ", ".", -1))
	return fmt.Sprintf("%s@%s", nameNormalized, domains[rand.Intn(len(domains))])
}

func generateDate(startYear, endYear int) string {
	min := time.Date(startYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(endYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

func GenerateEmployees(count int) []Employee {
	employees := make([]Employee, count)

	for i := 0; i < count; i++ {
		dept := departments[rand.Intn(len(departments))]
		jobs := jobTitles[dept]
		name := generateName()
		totalLeave := rand.Intn(16) + 10 // 10-25 days
		usedLeave := rand.Intn(totalLeave + 1)

		employees[i] = Employee{
			ID:                i + 1,
			Name:              name,
			Gender:            []string{"Male", "Female"}[rand.Intn(2)],
			Phone:             generatePhone(),
			Birthday:          generateDate(1970, 2000),
			Email:             generateEmail(name),
			JobTitle:          jobs[rand.Intn(len(jobs))],
			Department:        dept,
			HireDate:          generateDate(2015, 2023),
			Salary:            float64(rand.Intn(70000) + 30000),
			WorkStatus:        "Active",
			BankInfo:          fmt.Sprintf("Bank-%d-%d", rand.Intn(5)+1, rand.Intn(100000)),
			SupervisorID:      rand.Intn(count) + 1,
			TotalSpecialLeave: totalLeave,
			UsedSpecialLeave:  usedLeave,
		}
	}

	return employees
}

func SaveEmployeesToFile(employees []Employee, filename string) error {
	data, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}