# Catetduit API

API sederhana menggunakan Go dan Chi framework untuk mengembalikan data user static.

## Struktur Project

```
catetduit/
├── cmd/
│   └── api/
│       └── main.go          # Entry point aplikasi
├── internal/
│   └── user/
│       ├── model.go         # Model/struct User
│       ├── handler.go       # HTTP handlers untuk user
│       └── routes.go        # Routing untuk user endpoints
├── pkg/
│   └── logger/              # (untuk future use)
├── go.mod
└── README.md
```

## Teknologi

- Go 1.24
- Chi Router v5.2.3

## Cara Menjalankan

### 1. Install dependencies
```bash
go mod tidy
```

### 2. Jalankan aplikasi
```bash
go run ./cmd/api/main.go
```

Atau build terlebih dahulu:
```bash
go build -o bin/api.exe ./cmd/api
.\bin\api.exe
```

### 3. Test endpoint
```bash
curl http://localhost:8080/api/user
```

## Endpoints

### GET /api/user
Mengembalikan data user static (dummy data)

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

## Penjelasan Struktur

### Pemisahan Routes

Routes telah dipisahkan ke dalam file terpisah untuk modularitas yang lebih baik:

- **`internal/user/routes.go`**: Berisi function `RegisterRoutes()` yang mendaftarkan semua endpoint terkait user
- **`cmd/api/main.go`**: Hanya setup router, middleware, dan memanggil `user.RegisterRoutes()`

Keuntungan struktur ini:
- Setiap module (user, product, dll) bertanggung jawab atas routes-nya sendiri
- Main.go tetap clean dan tidak membengkak
- Mudah untuk scaling - tinggal tambah module baru
- Testable - bisa test routes per module

### Menambah Module Baru

Untuk menambah fitur baru (misal: product), ikuti langkah berikut:

1. Buat folder `internal/product`
2. Buat file:
   - `model.go` - struct Product
   - `handler.go` - handler functions
   - `routes.go` - RegisterRoutes function
3. Di `main.go` tambah: `product.RegisterRoutes(r)`

## Development

Server berjalan di port `:8080` dengan middleware:
- Logger - mencatat setiap HTTP request
- Recoverer - recover dari panic

