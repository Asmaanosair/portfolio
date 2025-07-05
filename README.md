# 🚀 Simple Go  Portfolio Project

This is a simple web server project built using **Go (Golang)**.  
It serves multiple static HTML pages with custom error handling using Go’s built-in `net/http` package.

The server listens on port `:9090` and serves a small personal portfolio-style website.

---

## 🌐 Pages Included

- `/home` → index.html (Home page)
- `/about` → about.html
- `/contact` → contact.html
- `/experience` → experience.html
- `/form` → form.html
- Custom error pages:
  - 404 Not Found → 404.html
  - 405 Method Not Allowed → 405.html
  - 500 Internal Server Error → 500.html

---

## 🔧 Tech Stack

- **Language:** Go
- **HTTP Package:** `net/http`
- **File Reading:** `os.ReadFile`
- **Static Directory:** `/statics`
- **Port:** `:9090`

---

## 🗂️ Project Structure

