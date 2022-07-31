## API Gateway
- Connect Order Service and User Service and Auth Service

## SET ENV
```sh
ORDER_GRPC=localhost:3000
USER_GRPC=localhost:3001
AUTH_GRPC=localhost:4000
PORT=1000
```

## ENDPOINT
```sh
Auth
- POST api/register
- POST api/login
- POST api/validate

Order
-GET api/orders

User
-GET api/users
```
