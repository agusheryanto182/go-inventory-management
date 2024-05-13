## What Is Inventory Management?
This is the backend used to serve customers

## Key Technologies

1. **Go**: Go (or Golang) is the programming language that can develop an API with high performance and scalable.

2. **Echo**: Echo is a high-performance, minimalist web framework for the Go programming language. It is known for its simplicity, speed, and robustness. Echo is commonly used for building web applications and APIs due to its fast routing capabilities and low memory footprint. It provides a clean and elegant API for handling HTTP requests, routing, middleware, and more.

3. **PostgreSQL**: PostgreSQL is a powerful, open-source relational database management system. It is used for storing and managing data related to products, customers, orders, and more in the online store application.

4. **JWT (JSON Web Tokens)**: JWT is a standard for securely transmitting information between parties as a JSON object. In the online store application, JWT is used for implementing authentication and authorization mechanisms.

5. **Pgx**: Pgx is a pure Go driver and toolkit for PostgreSQL. The pgx driver is a low-level, high performance interface that exposes PostgreSQL-specific features such as LISTEN / NOTIFY and COPY. It also includes an adapter for the standard database/sql interface.

6. **Sqlx**: Sqlx is a library for Go that provides extensions to the standard database/sql package. It aims to simplify working with SQL databases in Go by offering additional features and utilities. Sqlx includes functions for working with query results, struct scanning, named parameters, and more. It enhances the productivity of Go developers when interacting with SQL databases like PostgreSQL by reducing boilerplate code and providing a more intuitive API.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/agusheryanto182/go-inventory-management.git
   ```

2. Go to folder go-inventory-management 

   ```bash
   cd go-inventory-management
   ```

3. Set up env, copy the code and then paste it in your terminal

4. Create a new database

5. Copy this and paste it in terminal for migrate database

   ```bash
   make migrate-dev
   ```

7. ```bash
   go run .
   ```

## Documentation

1. Postman

   ```bash
   https://documenter.getpostman.com/view/32137512/2sA3JNafN6
   ```

2. Database

   ```bash
   https://dbdiagram.io/d/EniQilo-Stroe-66388de847ef755ec6ec15d1
   ```
