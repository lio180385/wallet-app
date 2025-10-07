# 💰 Wallet App (Golang Test Project)

A simple digital wallet backend service built with **Go + Gin + GORM + SQLite**.

## 🚀 Features
- Withdraw funds from wallet (`POST /withdraw`)
- Check balance (`GET /balance/:id`)
- Data stored in SQLite database

## ⚙️ How to Run
1. Clone this repo:
   ```bash
   git clone https://github.com/yourusername/wallet-app.git
   cd wallet-app
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the app:
   ```bash
   go run main.go
   ```

## 🧪 API Examples

### ✅ Check Balance
**GET** `/balance/1`

**Response:**
```json
{
  "user_id": 1,
  "name": "Leo",
  "balance": 100000
}
```

### 💸 Withdraw
**POST** `/withdraw`

**Body:**
```json
{
  "user_id": 1,
  "amount": 5000
}
```

**Response:**
```json
{
  "message": "withdrawal successful",
  "balance": 95000
}
```

## 🗄 Database
SQLite file: `wallet.db` (auto-created in root directory)

## 🕒 Deadline
Completed within 2 days as requested.
