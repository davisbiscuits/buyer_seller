# REST API example application

## Run the app

    docker-compose up

## Run the tests

    go test ./...

# REST API

The REST API to the example app is described below.

## Register Buyer

### Request

`POST /register`

    curl --location 'localhost:8080/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "test",
        "email": "test@g.c",
        "type": "buyer"
    }'

### Response

    Status: 201
    Body: {"id": 1, "message": "SuccessFully added User" }

## Register Seller

### Request

`POST /register`

    curl --location 'localhost:8080/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "test",
        "email": "test@g.c",
        "type": "seller"
    }'

### Response

    Status: 201
    Body: {"id": 1, "message": "SuccessFully added User" }


## Add Buyer Requirement

### Request

`POST /buyer/add-requirement`

    curl --location 'localhost:8080/buyer/add-requirement' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "test",
        "quantity": 10,
        "price": 100,
        "user_id": 1
    }'

### Response

    Status: 200
    Body: {"id":1, "message": "SuccessFully added Requirement" }

## Find Sellers for a buyer product

### Request

`GET /buyer/match`

    curl --location 'localhost:8080/buyer/match' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "product_id": 1
    }'

### Response

    Status: 200
    Body: [{"id": 10, name: "test", email: "test@g.c"}]

## Place order

### Request

`POST /buyer/place-order`

    curl --location 'localhost:8080/buyer/place-order' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "user_id": 1,
    }'

### Response

    Status: 200
    Body: {"id": 2, "message": "SuccessFully placed Order"}

## Get list orders by buyer

### Request

`GET /buyer/orders`

    curl --location 'localhost:8080/buyer/orders' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "user_id": 1,
    }'

### Response

    Status: 200
    Body: [{"id":1,"seller":{"id":1,"name":"seller_name","email":"seller@email.com"},"requirement":{"id":2,"name":"req_name","quantity":10,"price":10},"status":"pending"}]


## Add Seller Product

### Request

`POST /seller/add-product`

    curl --location 'localhost:8080/seller/add-product' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "test",
        "quantity": 10,
        "price": 100,
        "user_id": 1
    }'

### Response

    Status: 200
    Body: {"id":1, "message": "SuccessFully added Product" }

## Accept Order

### Request

`POST /seller/accept-order`

    curl --location 'localhost:8080/seller/accept-order' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "order_id": 1,
        "user_id": 10
    }'

### Response

    Status: 200
    Body: {"message": "SuccessFully accepted Order" }

## Reject Order

### Request

`POST /seller/reject-order`

    curl --location 'localhost:8080/seller/reject-order' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "order_id": 1
    }'

### Response

    Status: 200
    Body: {"message": "SuccessFully rejected Order" }

## Get list orders for seller

### Request

`GET /seller/orders`

    curl --location 'localhost:8080/seller/orders' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "user_id": 1,
    }'

### Response

    Status: 200
    Body: [{"id":1,"seller":{"id":1,"name":"seller_name","email":"seller@email.com"},"requirement":{"id":2,"name":"req_name","quantity":10,"price":10},"status":"pending"}]
