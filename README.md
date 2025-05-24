# Company API

A RESTful API for company CRUD operations built with Echo framework.

## Project Structure

```
echopad/
├── api/
│   ├── controllers/
│   │   └── company_controller.go
│   ├── models/
│   │   └── company.go
│   └── routes/
│       └── company_routes.go
├── main.go
├── go.mod
└── go.sum
```

## API Endpoints

| Method | Endpoint              | Description               |
|--------|------------------------|---------------------------|
| GET    | /api/companies        | Get all companies         |
| GET    | /api/companies/:id    | Get a company by ID       |
| POST   | /api/companies        | Create a new company      |
| PUT    | /api/companies/:id    | Update a company by ID    |
| DELETE | /api/companies/:id    | Delete a company by ID    |

## Running the Application

```bash
go run main.go
```

The server will start on port 8080.

## API Examples

### Create a company (POST /api/companies)

Request:
```json
{
  "name": "ABC Corp",
  "address": "456 Business Ave, Suite 100",
  "phone": "555-123-4567",
  "email": "contact@abccorp.com",
  "website": "https://abccorp.com"
}
```

### Update a company (PUT /api/companies/:id)

Request:
```json
{
  "name": "ABC Corporation",
  "address": "456 Business Avenue, Suite 200",
  "phone": "555-123-9999",
  "email": "info@abccorp.com",
  "website": "https://abccorp.com"
}
```
