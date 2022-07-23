run-order:
	@echo "Running... Service Order"
	cd ./order-service && go run server.go

run-user:
	@echo "Running... Service User"
	cd ./user-service && go run server.go

run-auth:
	@echo "Running... Service Auth"
	cd ./auth-service && go run server.go

run-api:
	@echo "Running... API Services"
	cd ./api-gateway && go run main.go