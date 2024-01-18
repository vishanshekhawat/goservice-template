#----------------------------------------
tidy:
	go mod tidy
	go mod vendor

	
run-user-service-local:
	go run app/services/user-service/main.go 
