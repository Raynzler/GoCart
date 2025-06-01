# 🚛GoCart

A full-stack Shopping application using Go (Gin) as the backend and React for the frontend. Stripe is integrated to handle secure payments. (Built for understanding Payments and Integration)



### 🚀 Features

- 🧠 Built with Go (Gin) backend and React frontend
- 💳 Stripe payment integration (card input, client secret)
- 📦 RESTful API for product checkout

### 🛠️ Tech Stack

- **Backend**: Go (Gin), Stripe SDK
- **Frontend**: React, Stripe Elements

## 📚API Reference

#### Payment intent creation check 

```http
  POST /create-payment-intent
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `product_id` | `string` | This gives back reply as to specify whether customer's payment got accepted in backend |

#### Health status

```http
  GET /health
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | Get health status of backend server |



## 💻Run Locally

Clone the project

```bash
  git clone https://github.com/Raynzler/StripeGO-PayApp.git
```

Go to the project directory & start backend

```bash
cd backend\server
go run main.go
```

Install dependencies & start frontend

```bash
cd ../../frontend
npm install
npm start
```

Server Started, shop & Checkout

```bash
Shop your favourite item
```

## License

This project is licensed under the [MIT License](./LICENSE).

## ⚠️ Disclaimer

This software is provided "as-is" without any warranties.  
The author is not responsible or liable for any damages or losses resulting from its use, including misuse of real credit card information.  
Please use only Stripe's official test card numbers during development and testing.