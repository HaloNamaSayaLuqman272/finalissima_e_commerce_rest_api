# finalissima_e_commerce_rest_api
## REST API E-COMMERCE PROJECT

## 🔑 Key Features
    - 💾 API Connection
    - 🔐 Basic Authentication and JWT
    - 🗄️ Product Management
    - 📨 Purchase Management

## 👫 Role
    - 🧍‍♂️ Admin
    - 🧍‍♀️ User

## 🖨️ Tech Stacks
    - Echo: Web Framework
    - GORM: ORM (Object-Relational Mapping) Library
    - Validator: Validate Library
    - Zap: Logging
    - Viper: Configuration

## 🖥️ How to Use
    - 📱 Create a custom type in database.
        -- create role type
        CREATE TYPE role AS ENUM ('user', 'admin');

        -- create purchase status type
        CREATE TYPE purchase_status AS ENUM (
            'pending', 'paid', 'on_delivery', 'cancelled', 'received'
        );
    - 📲 Generate admmin by code and .env
        go run ./cmd/generate/.
    - 💻 Run the application
        go run ./cmd/api/.

## 📑 Additional Notes
    📄 The docs directory contains Postman collection examples including:
        - 💾 OpenRouter Playground: Sample request of OpenRouter API
        - 🚚 RajaOngkir Playground: Sample request of RajaOngkir API