# Running Application with Docker Compose

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/zaenalarifin12/

API endpoint
https://documenter.getpostman.com/view/4549598/2sA3Bn6CX1


# API Endpoints

## Authentication API

### 1. `/api/v1/login`
- **Description:** This API is used for user authentication and login.
- **Method:** POST

### 2. `/api/v1/register`
- **Description:** Allows users to register and create a new account.
- **Method:** POST

## Product API

### 1. `/api/v1/products/list`
- **Description:** Retrieves a list of all available products.
- **Authentication:** Required
- **Level:** 1

### 2. `/api/v1/products/{id}`
- **Description:** Retrieves details of a specific product by its ID.
- **Method:** GET
- **Authentication:** Required
- **Level:** 1


### 3. `/api/v1/products`
- **Description:** Allows admin users to create a new product.
- **Method:** POST
- **Authentication:** Required
- **Level:** 1

### 4. `/api/v1/products/{id}`
- **Description:** Allows admin users to update details of an existing product.
- **Method:** PUT
- **Authentication:** Required
- **Level:** 1

### 5. `/api/v1/products/{id}`
- **Description:** Allows admin users to delete a product by its ID.
- **Method:** DELETE
  - **Authentication:** Required
- **Level:** 2

                                 +---------------------------+
                                 |            API            |
                                 +--------------+------------+
                                                |
                                                |
                                 +--------------v------------+
                                 |   Authentication Service  |
                                 +--------------+------------+
                                                |
                                                |
               +-----------------------------+  |  +-------------------------------+
               |        /api/v1/login        |  |  |       /api/v1/register        |
               |      (POST Method)          |  |  |     (POST Method)             |
               |  Authentication Required    |  |  |   Public Access               |
               +-----------------------------+  |  +-------------------------------+
                                                |
                                                |
                                 +--------------v------------+
                                 |       Product Service      |
                                 +--------------+------------+
                                                 |
              +--------------------------------+ |  +--------------------------------+
              |  /api/v1/products/list         | |  |  /api/v1/products/{id}         |
              |   (GET Method)                 | |  |  (GET Method)                  |
              |  Authentication Required       | |  |  Authentication Required       |
              |  Access Level: 1               | |  |  Access Level: 1               |
              +--------------------------------+ |  +--------------------------------+
                                                 |
                                                 |
              +--------------------------------+ |  +--------------------------------+
              |  /api/v1/products              | |  |  /api/v1/products/{id}         |
              |   (POST Method)                | |  |  (PUT Method)                  |
              |  Authentication Required       | |  |  Authentication Required       |
              |  Access Level: 1               | |  |  Access Level: 2               |
              +--------------------------------+ |  +--------------------------------+
                                                 |
                                                 |
                                   +--------------v------------+
                                   |       PostgreSQL DB       |
                                   +---------------------------+
