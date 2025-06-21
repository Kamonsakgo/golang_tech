# Simple E-commerce API with Affiliate System

## 📋 คำอธิบายโปรเจค

โปรเจคนี้เป็น **RESTful API** ที่พัฒนาด้วยภาษา **Go (Golang)** สำหรับระบบ E-commerce พร้อมระบบ Affiliate Marketing โดยมีคุณสมบัติหลักดังนี้:

- 🛒 **ระบบจัดการสินค้า** (Product Management)
- 👥 **ระบบจัดการผู้ใช้** (User Management) 
- 💰 **ระบบจัดการยอดเงิน** (Balance Management)
- 🤝 **ระบบ Affiliate** (Affiliate System)
- 💼 **ระบบคอมมิชชั่น** (Commission System)

## 🛠️ เทคโนโลยีที่ใช้

- **Backend Framework**: Go (Golang) + Gin Web Framework
- **Database**: PostgreSQL
- **ORM/Query Builder**: SQLC (SQL Compiler)
- **Database Migration**: Go-migrate
- **Container**: Docker
- **API Format**: RESTful API (JSON)

## 📊 โครงสร้างฐานข้อมูล

โปรเจคนี้ประกอบด้วย 4 ตารางหลัก:

1. **users** - จัดเก็บข้อมูลผู้ใช้และยอดเงิน
2. **product** - จัดเก็บข้อมูลสินค้าและราคา
3. **affiliate** - จัดเก็บข้อมูล affiliate และเปอร์เซ็นต์คอมมิชชั่น
4. **commission** - จัดเก็บข้อมูลคอมมิชชั่นจากการขาย

## 🚀 API Endpoints

### 👥 User Management
- `POST /createuser` - สร้างผู้ใช้ใหม่
- `GET /user/all` - ดูรายการผู้ใช้ทั้งหมด
- `GET /user/:id` - ดูข้อมูลผู้ใช้ตาม ID
- `PATCH /user/deduct/balance/:id` - หักยอดเงินผู้ใช้
- `PATCH /user/add/balance/:id` - เพิ่มยอดเงินผู้ใช้

### 🛒 Product Management
- `POST /product` - เพิ่มสินค้าใหม่
- `GET /product/list` - ดูรายการสินค้าทั้งหมด
- `GET /product/:id` - ดูข้อมูลสินค้าตาม ID

### 🤝 Affiliate Management
- `POST /affiliate` - เพิ่ม affiliate ใหม่
- `GET /affiliate/list` - ดูรายการ affiliate ทั้งหมด
- `GET /affiliate/:id` - ดูข้อมูล affiliate ตาม ID

### 💼 Commission Management
- `GET /commission/:id` - ดูข้อมูลคอมมิชชั่นตาม ID
- `GET /commission/list` - ดูรายการคอมมิชชั่นทั้งหมด

### 🛍️ Purchase
- `POST /buyproduct` - ซื้อสินค้า (พร้อมคำนวณคอมมิชชั่น)

## 📦 การติดตั้งและรันโปรเจค

### ความต้องการของระบบ
- Go 1.23.2 หรือใหม่กว่า
- Docker
- PostgreSQL
- Go-migrate tool

### ขั้นตอนการติดตั้ง

1. **Clone โปรเจค**
```bash
git clone https://github.com/Kamonsakgo/golang_tech.git
cd simple
```

2. **สร้าง PostgreSQL Database ด้วย Docker**
```bash
make postgres
make createdb
```

3. **รันการ Migration**
```bash
make migrateup
```

4. **Generate Code จาก SQL**
```bash
make sqlc
```

5. **รัน Server**
```bash
make server
```

Server จะรันที่ `http://localhost:8080`

## 📁 โครงสร้างโปรเจค

```
simple/
├── api/                    # API Handlers
│   ├── account.go         # User API handlers
│   ├── affiliate.go       # Affiliate API handlers
│   ├── commission.go      # Commission API handlers
│   ├── product.go         # Product API handlers
│   └── server.go          # Server setup และ routes
├── db/
│   ├── migration/         # Database migration files
│   ├── query/             # SQL query files
│   └── sqlc/              # Generated Go code จาก SQLC
├── util/                  # Utility functions
├── main.go                # Entry point
├── go.mod                 # Go dependencies
├── Makefile               # Automation commands
└── sqlc.yaml              # SQLC configuration
```

## 🔧 คำสั่ง Makefile

- `make postgres` - สร้าง PostgreSQL container
- `make createdb` - สร้าง database
- `make dropdb` - ลบ database
- `make migrateup` - รัน migration ขึ้น
- `make migratedown` - รัน migration ลง
- `make sqlc` - Generate Go code จาก SQL
- `make server` - รัน server
- `make test` - รันการทดสอบ

## 🌐 การตั้งค่า Database

แก้ไขการตั้งค่าการเชื่อมต่อฐานข้อมูลในไฟล์ `main.go`:

```go
const (
    dbDriver      = "postgres"
    dbSource      = "postgresql://root:1234@localhost:5432/simple?sslmode=disable"
    ServerAddress = "0.0.0.0:8080"
)
```

## 📝 การใช้งาน API

### ตัวอย่างการสร้างผู้ใช้ใหม่
```bash
curl -X POST http://localhost:8080/createuser \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "balance": 1000.00}'
```

### ตัวอย่างการซื้อสินค้า
```bash
curl -X POST http://localhost:8080/buyproduct \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user-uuid", "product_id": "product-uuid", "quantity": 2}'
```

## 🧪 การทดสอบ

รันการทดสอบด้วยคำสั่ง:
```bash
make test
```

## 👨‍💻 ผู้พัฒนา

- GitHub: [Kamonsakgo](https://github.com/Kamonsakgo)
- Repository: [golang_tech](https://github.com/Kamonsakgo/golang_tech)

## 📄 License

โปรเจคนี้อยู่ภายใต้ MIT License

---

**หมายเหตุ**: โปรเจคนี้พัฒนาขึ้นเพื่อการศึกษาและเป็นตัวอย่างการสร้าง E-commerce API ด้วย Go และ PostgreSQL 