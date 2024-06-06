# These commands are used to execute the program
build:
	go build -o bin/main main.go
run:
	go run main.go
# test:
# 	go test -v ./...

---------------------------------------------------------
## These commands runs the docker restapi container
docker-build:
	docker build -t restapi:1.0.0 . 

docker-run:
	 docker run  --env-file ./.env   -d -p 8081:8081 --name go-app restapi:1.0.0 

docker-stop:
	docker stop rest-api-container && docker rm rest-api-container

----------------------------------------------------------
## These commands are used to run the application on docker compose.
compose-up:
	docker compose up -d 

compose-down:
	docker compose down


.PHONY: build run test docker-build docker-run docker-stop compose-up compose-down
