# CRM Backend API

### This project starts a http server that provides API to manage customers with CRUD operations

### Instructions to run http server
- Navigate to directory where *main.go* is in.
- Run `go run main.go` to start the server running

### API endpoints
- Get all customers: `GET /customers`
- Get customer by ID: `GET /customers/{id}`
- Add new customer: `POST /customers`
  - Body: { 'name': '{newName}', 'role': '{newRole}', 'email': '{newEmail}' }
  - Note that contacted will be *false* for new customers
- Update customer by ID: `PATCH /customers/{id}`
  - Body: { 'name': '{newName}', 'role': '{newRole}', 'email': '{newEmail}', 'contacted': 'true/false' }
- Delete customer by ID: `DELETE /customers/{id}`

