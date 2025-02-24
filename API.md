# Employee Management API Documentation

## Base URL
`http://localhost:8080`

## Endpoints

### List All Employees
**GET** `/api/employees`

Retrieve a list of all employees.

#### Examples
```bash
# Get all employees
curl http://localhost:8080/api/employees

# Filter by department
curl "http://localhost:8080/api/employees?department=Engineering"

# Filter by job title
curl "http://localhost:8080/api/employees?job_title=Software%20Engineer"

# Filter by name (partial match)
curl "http://localhost:8080/api/employees?name=John"

# Filter by years of service (greater than 5 years)
curl "http://localhost:8080/api/employees?years_of_service=>5"

# Filter by salary (between 50000 and 80000)
curl "http://localhost:8080/api/employees?salary=>50000&salary=<80000"

# Combine multiple filters
curl "http://localhost:8080/api/employees?department=Engineering&salary=>60000"
```

#### Query Parameters
- `department` (optional): Filter employees by department
- `job_title` (optional): Filter employees by job title
- `name` (optional): Filter employees by name
- `years_of_service` (optional): Filter employees by years of service with comparison operators (>n, <n, >=n, <=n, =n)
- `salary` (optional): Filter employees by salary with comparison operators (>n, <n, >=n, <=n, =n)
- `limit` (optional): Limit the number of results (default: 50)
- `offset` (optional): Number of records to skip (default: 0)

#### Response
```json
{
  "total": 50,
  "employees": [
    {
      "id": 1,
      "name": "John Wilson",
      "gender": "Male",
      "phone": "+1-019-985-7274",
      "birthday": "1999-01-22",
      "email": "john.wilson@outlook.com",
      "job_title": "Recruiter",
      "department": "HR",
      "hire_date": "2019-12-01",
      "salary": 88422,
      "work_status": "Active",
      "bank_info": "Bank-5-80263",
      "supervisor_id": 35,
      "total_special_leave": 13,
      "used_special_leave": 9
    }
  ]
}
```

### Get Employee by ID
**GET** `/api/employees/{id}`

Retrieve details of a specific employee.

#### Example
```bash
# Get employee with ID 1
curl http://localhost:8080/api/employees/1
```

#### Response
```json
{
  "id": 1,
  "name": "John Wilson",
  "gender": "Male",
  "phone": "+1-019-985-7274",
  "birthday": "1999-01-22",
  "email": "john.wilson@outlook.com",
  "job_title": "Recruiter",
  "department": "HR",
  "hire_date": "2019-12-01",
  "salary": 88422,
  "work_status": "Active",
  "bank_info": "Bank-5-80263",
  "supervisor_id": 35,
  "total_special_leave": 13,
  "used_special_leave": 9
}
```

### Create Employee
**POST** `/api/employees`

Create a new employee record.

#### Example
```bash
curl -X POST http://localhost:8080/api/employees \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Employee",
    "gender": "Female",
    "phone": "+1-123-456-7890",
    "birthday": "1990-01-01",
    "email": "new.employee@company.com",
    "job_title": "Software Engineer",
    "department": "Engineering",
    "hire_date": "2023-01-01",
    "salary": 75000,
    "work_status": "Active",
    "bank_info": "Bank-1-12345",
    "supervisor_id": 1,
    "total_special_leave": 15,
    "used_special_leave": 0
  }'
```

#### Request Body
```json
{
  "name": "New Employee",
  "gender": "Female",
  "phone": "+1-123-456-7890",
  "birthday": "1990-01-01",
  "email": "new.employee@company.com",
  "job_title": "Software Engineer",
  "department": "Engineering",
  "hire_date": "2023-01-01",
  "salary": 75000,
  "work_status": "Active",
  "bank_info": "Bank-1-12345",
  "supervisor_id": 1,
  "total_special_leave": 15,
  "used_special_leave": 0
}
```

#### Response
```json
{
  "id": 51,
  "message": "Employee created successfully"
}
```

### Update Employee
**PUT** `/api/employees/{id}`

Update an existing employee record.

#### Example
```bash
curl -X PUT http://localhost:8080/api/employees/1 \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+1-987-654-3210",
    "email": "updated.email@company.com",
    "salary": 80000,
    "work_status": "Active"
  }'
```

#### Request Body
```json
{
  "phone": "+1-987-654-3210",
  "email": "updated.email@company.com",
  "salary": 80000,
  "work_status": "Active"
}
```

#### Response
```json
{
  "message": "Employee updated successfully"
}
```

### Delete Employee
**DELETE** `/api/employees/{id}`

Delete an employee record.

#### Example
```bash
curl -X DELETE http://localhost:8080/api/employees/1
```

#### Response
```json
{
  "message": "Employee deleted successfully"
}
```

### Get Employee Leave Balance
**GET** `/api/employees/{id}/leave`

Retrieve the leave balance for a specific employee.

#### Example
```bash
curl http://localhost:8080/api/employees/1/leave
```

#### Response
```json
{
  "employee_id": 1,
  "total_special_leave": 13,
  "used_special_leave": 9,
  "remaining_special_leave": 4
}
```

## Error Responses

In case of errors, the API will return appropriate HTTP status codes along with error messages:

```json
{
  "error": "Error message description",
  "status": 400
}
```

### Common Status Codes
- 200: Success
- 201: Created
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Rate Limiting

API requests are limited to 100 requests per minute per IP address. When exceeded, the API will return a 429 (Too Many Requests) status code.

## Authentication

All API endpoints require authentication using an API key that should be included in the request header:

```
Authorization: Bearer YOUR_API_KEY
```

