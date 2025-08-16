# Go Monitoring Service

A simple monitoring system built with **Go + Gin + GORM + JWT**, allowing users to:

* Register/Login (JWT-based authentication with cookies).
* Add and manage monitors (HTTP, DNS/SSL, DB, Dead Man’s Switch).
* Automatically run checks on each monitor at user-defined intervals.
* Store and display check results (status, latency, details) in a dashboard.

---

## Features

✅ **User Authentication** (Register, Login, Logout)

✅ **JWT-secured routes** with middleware

✅ **Monitor Management**

  * Add new monitors with type, target, method, expected status, interval, etc.
    
✅ **Background Scheduler**

  * Runs monitor checks automatically in the background based on their interval
    
✅ **Results Storage**

  * Saves status, latency, and details in a `results` table
    
✅ **Dashboard**

  * Each user can view all their monitors and latest results

---

## Project Structure

```
go-monitoring-service/
├── controllers/        # Request handlers (auth, dashboard, monitor)
├── middlewares/        # Auth middleware
├── models/             # Database models (User, Monitor, Result)
├── services/           # Background jobs & monitoring logic
├── templates/          # HTML templates (login, register, dashboard, etc.)
├── initializers/       # DB + env initialization
├── main.go             # Entry point
```

---

## How It Works

1. **User registers & logs in** → gets JWT token in cookie.
2. **User adds a monitor** (e.g., HTTP check for `https://example.com`).
3. **Scheduler runs** → periodically checks all active monitors:

   * Runs HTTP request / DNS check / DB check, etc.
   * Saves result into the `results` table.
   * Updates monitor’s `LastCheckedAt`.
4. **Dashboard** → Displays user’s monitors + latest results.


---

## Installation

### 1. Clone repository

```bash
git clone https://github.com/yourusername/go-monitoring-service.git
cd go-monitoring-service
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set environment variables (`.env`)

```
DB_URL=postgres://user:password@localhost:5432/monitoring
SECRET=your_jwt_secret
```

### 4. Run the service

```bash
go run main.go
```

---

## Roadmap

* [ ] Implement DNS/SSL monitor checks
* [ ] Implement DB monitor checks
* [ ] Implement Dead Man’s Switch monitor
* [ ] Improve UI for dashboard (live refresh with HTMX/WebSockets)
* [ ] Dockerize project
