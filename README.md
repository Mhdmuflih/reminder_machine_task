# Reminder System

A simple Reminder Management API built with Golang (Gin framework) and PostgreSQL, with full CRUD operations for Tasks and Reminder Rules, and an audit log system to track all actions.

---

## Features

- **Manage Tasks:** create, update, delete, view all.
- **Manage Reminder Rules:** create, update, activate/deactivate, delete, view all.
- **Audit Logs:** Automatically records every create/update/delete action for Tasks and Rules.

**Built using:** Gin framework, GORM, and PostgreSQL.

---

## Tech Stack

- **Backend:** Golang (Gin)  
- **Database:** PostgreSQL  
- **ORM:** GORM  
- **Environment Management:** Godotenv  

---

## Project Structure
reminder/
├── config/ # Database connection and config
├── controllers/ # Task, Rule, and Log controllers
├── models/ # Task, ReminderRule, AuditLog models
├── repositories/ # DB access logic for Task, Rule, Log
├── routes/ # API routes for Task, Rule, Log
├── services/ # Business logic for Task, Rule, Log
├── main.go # Entry point
└── .env # Environment variables


## Installation

### Clone the repository

```bash
git clone <your-repo-url>
cd reminder


Install dependencies
go mod tidy
Set up .env

Create a .env file in the root:

PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=reminder_db
Run the application
go run main.go

Server will run at http://localhost:8080



Database

Make sure PostgreSQL is running.
Create database manually or let GORM migrate tables automatically.

Tables:
| Column     | Type |
| ---------- | ---- |
| id (PK)    |      |
| title      |      |
| due_at     |      |
| status     |      |
| created_at |      |
| updated_at |      |


reminder_rules
| Column         | Type |
| -------------- | ---- |
| id (PK)        |      |
| name           |      |
| minutes_before |      |
| is_active      |      |
| created_at     |      |
| updated_at     |      |


audit_logs
| Column     | Type |
| ---------- | ---- |
| id (PK)    |      |
| action     |      |
| user       |      |
| created_at |      |


API Endpoints

Tasks
| Method | Route      | Description   |
| ------ | ---------- | ------------- |
| GET    | /tasks/    | Get all tasks |
| POST   | /tasks/    | Create a task |
| PUT    | /tasks/:id | Update a task |
| DELETE | /tasks/:id | Delete a task |



Reminder Rules
| Method | Route                 | Description       |
| ------ | --------------------- | ----------------- |
| GET    | /rules/               | Get all rules     |
| POST   | /rules/               | Create a rule     |
| PUT    | /rules/:id            | Update a rule     |
| PATCH  | /rules/:id/active     | Activate a rule   |
| PATCH  | /rules/:id/deactivate | Deactivate a rule |
| DELETE | /rules/:id            | Delete a rule     |


Audit Logs
| Method | Route | Description        |
| ------ | ----- | ------------------ |
| GET    | /logs | Get all audit logs |


Usage Example
Create Task

POST /tasks/
Content-Type: application/json

{
  "title": "Finish Backend API",
  "due_at": "2026-02-08T20:00:00Z",
  "status": "pending"
}


### Response Example

```json
{
  "task": {
    "id": 1,
    "title": "Finish Backend API",
    "due_at": "2026-02-08T20:00:00Z",
    "status": "pending",
    "created_at": "2026-02-07T23:00:00Z",
    "updated_at": "2026-02-07T23:00:00Z"
  }
}

**Audit Log created automatically:**

```json
{
  "action": "Created Task: Finish Backend API",
  "user": "Muflih",
  "created_at": "2026-02-07T23:00:00Z"
}



✅ Notes:  

1. Make sure you **use three backticks (\`\`\`)** before and after the JSON.  
2. Specify `json` right after the first backticks for proper syntax highlighting.  
3. Do **not leave extra spaces before the opening `{`**.  

If you want, I can rewrite your **entire README.md** properly formatted so that **all JSON examples, bash commands, and project structure** show correctly without any “error” in GitHub or VS Code preview.  

Do you want me to do that?