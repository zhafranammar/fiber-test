# API Documentation

## Endpoint List

| Method | Endpoint                   | Description                                              |
| ------ | -------------------------- | -------------------------------------------------------- |
| GET    | `/public`                  | Mengambil data produk publik                             |
| GET    | `/protected`               | Mengambil data produk yang hanya bisa diakses dengan JWT |
| GET    | `/generate-token/:user_id` | Menghasilkan JWT berdasarkan user_id                     |

---

## 1. Public Endpoint

### **GET `/public`**

#### **Deskripsi:**

Endpoint ini bisa diakses oleh siapa saja dan akan mengembalikan daftar produk publik.

#### **Response:**

```json
{
  "status": true,
  "message": "This is a public endpoint",
  "data": [
    { "id": 1, "name": "Laptop ASUS ROG", "price": 25000000 },
    { "id": 2, "name": "Keyboard Mechanical", "price": 1500000 },
    { "id": 3, "name": "Mouse Gaming", "price": 500000 }
  ]
}
```

---

## 2. Generate JWT Token

### **GET `/generate-token/:user_id`**

#### **Deskripsi:**

Endpoint ini menghasilkan token JWT berdasarkan `user_id` yang diberikan.

#### **Request Parameter:**

| Parameter | Tipe | Deskripsi                        |
| --------- | ---- | -------------------------------- |
| `user_id` | int  | ID pengguna untuk generate token |

#### **Response (Success):**

```json
{
  "status": true,
  "message": "Token generated successfully",
  "data": { "token": "<JWT_TOKEN>" }
}
```

#### **Response (Invalid User ID):**

```json
{
  "status": false,
  "message": "Invalid user ID",
  "data": null
}
```

#### **Response (Failed to Generate Token):**

```json
{
  "status": false,
  "message": "Failed to generate token",
  "data": null
}
```

---

## 3. Protected Endpoint

### **GET `/protected`**

#### **Deskripsi:**

Endpoint ini hanya dapat diakses dengan token JWT yang valid.

#### **Header:**

| Key             | Value                |
| --------------- | -------------------- |
| `Authorization` | `Bearer <JWT_TOKEN>` |

#### **Response (Success - Valid Token):**

```json
{
  "status": true,
  "message": "This is a protected endpoint",
  "data": [
    { "id": 101, "name": "Laptop MSI", "price": 12000000 },
    { "id": 102, "name": "Keyboard Mechanical", "price": 200000 },
    { "id": 103, "name": "Mouse Gaming", "price": 120000 }
  ]
}
```

#### **Response (Unauthorized - No Token):**

```json
{
  "status": false,
  "message": "Unauthorized",
  "data": null
}
```

#### **Response (Unauthorized - Invalid Token):**

```json
{
  "status": false,
  "message": "Invalid or expired token",
  "data": null
}
```

---

## Testing API

### **Jalankan Server**

```bash
make run
```

### **Gunakan `curl` untuk Testing**

- **Public Endpoint:**
  ```bash
  curl -X GET http://localhost:3000/public
  ```
- **Generate Token:**
  ```bash
  curl -X GET http://localhost:3000/generate-token/123
  ```
- **Akses Protected Endpoint dengan Token:**

  ```bash
  curl -X GET http://localhost:3000/protected -H "Authorization: Bearer <JWT_TOKEN>"
  ```

  ***

## Testing

### **Jalankan Test Scenario**

```bash
make test
```
