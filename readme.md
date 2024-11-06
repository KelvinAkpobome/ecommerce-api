Ecommerce API

An API for managing an e-commerce platform, including user management, product handling, and order processing.

Features
- User Management: Register, authenticate, and manage user accounts.
- Product Management: Add, update, delete, and fetch products.
- Order Management: Place, update, fetch, and cancel orders.
- Status Management: Track and update the status of orders.

Tech Stack
- Go: The backend is built with Go (Golang).
- Gin: A web framework for building the RESTful API.
- GORM: ORM for interacting with the database.
- PostgreSQL: The database for storing user, product, and order data.
- Validator: For input validation.
- JWT: For user authentication (if applicable).

Installation

Prerequisites
Ensure that you have the following installed:
- Go (version 1.18 or higher)
- PostgreSQL (for database)
- Git (for cloning the repo)

Steps
1. Clone the repository:
   git clone https://github.com/kelvinakpobome/ecommerce-api.git

2. Navigate into the project directory:
   cd ecommerce-api

3. Install dependencies:
   go mod tidy

4. Set up environment variables for the database connection, JWT secrets, etc. Create a .env file with the following:

   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=ecommerce_db
   JWT_SECRET=yoursecretkey

5. Run the application:
   go run main.go

   The application should now be running on localhost:8080.

API Endpoints
(Prefix `/api`)

User Endpoints
- POST /register  
  Register a new user.  
  Body: { "email": "string", "password": "string" }

- POST /login  
  Login to the platform and get a JWT token.  
  Body: { "email": "string", "password": "string" }

Product Endpoints
- GET /products  
  Fetch all products.  
  Response: [{ "id": 1, "name": "Product", "price": 99.99 }]

- GET /products/:id  
  Fetch a specific product by ID.  
  Response: {"id": 1, "name": "Product", "price": 99.99}

- POST /products  
  Add a new product.  
  Body: { "name": "Product", "price": 99.99 }

- PUT /products/:id  
  Update an existing product.  
  Body: { "name": "Updated Product", "price": 89.99 }

- DELETE /products/:id  
  Delete a product.

Order Endpoints
- POST /orders  
  Place a new order.  
  Body: { "product_ids": [1, 2] }

- GET /orders  
  Get all orders for the logged-in user.  
  Response: [{ "id": 1, "status": 0, "products": [...] }]

- PUT /orders/:id  
  Cancel an order.  


Admin Endpoints (Prefix `/api/admin`)
- GET /products  
  Fetch all products (Admin only).

- POST /products  
  Add a new product (Admin only).  
  Body: { "Name": "Product", "Price": 99.99 }

- PUT /products/:id  
  Update an existing product (Admin only).  
  Body: { "Name": "Updated Product", "Price": 89.99 }

- DELETE /products/:id  
  Delete a product (Admin only).

- PUT /api/admin/orders/:id 
  Update the status of an order (Admin only).  
  Body: { "status": 1 }

Authentication
This API uses JWT for authentication. Once a user registers and logs in, a JWT token is generated. This token should be included in the Authorization header for protected routes.

Example header for authenticated requests:
Authorization: Bearer <jwt_token>

Configuration
Set up the following environment variables for the project:
- DB_HOST: Database host (e.g., localhost).
- DB_PORT: Database port (e.g., 5432).
- DB_USER: Database username.
- DB_PASSWORD: Database password.
- DB_NAME: Database name.
- JWT_SECRET: Secret key for signing JWT tokens.

Database Setup
1. Install PostgreSQL and create a new database.
2. Configure the database connection in the environment variables.
3. Run migrations to set up the database schema.
   If you're using GORM with auto migrations, the schema will automatically be created when the application starts.

Contributing
Feel free to fork the repository and create a pull request with your contributions. Please ensure that your code follows the project's style guidelines and includes tests where applicable.

License
This project is licensed under the MIT License.

---
Notes
- Product Management: Admin-only access to manage products. You can create, update, delete, and list products.
- Order Management: Users can place orders, and admins can update order status.
- Error Handling: The API returns detailed error messages to help debug any issues that occur.
