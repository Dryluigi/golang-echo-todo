# Todo List REST API

Demo REST API sederhana menggunakan Golang dan framework Echo.

## Setup

### Database
1. Buat database MySQL baru dengan nama DB `todos`.
2. Buat tabel baru dengan nama `todos`. Isi tabel dengan kolom ID (integer), Title (varchar), Description (varchar), dan Done (tinyint)

### Running
1. Jalankan `go run main.go`

## Endpoints

- Buat todo baru : `POST /todos`
- List semua todo : `GET /todos`
- Hapus todo : `DELETE /todos/{id}`
- Update todo : `PATCH /todos/{id}`
- Check todo : `PATCH /todos/{id}/check`
