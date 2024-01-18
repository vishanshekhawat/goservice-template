#----------------------------------------
tidy:
	go mod tidy
	go mod vendor

	
run-user-service-local:
	go run app/services/user-service/main.go 


test-endpoint-local:
	curl -il localhost:3000/api/api1/v2/12

