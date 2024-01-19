#----------------------------------------
tidy:
	go mod tidy
	go mod vendor

	
run-user-service-local:
	go run app/services/user-service/main.go 


test-endpoint-local:
	curl -il localhost:3000/test

test-get-users:
	curl --location 'http://localhost:3000/api/v1/get-users' --header 'Content-Type: application/json' --data '{"token":"2121"}'