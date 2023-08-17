.PHONY: dc run 

dc: 
	docker-compose up --remove-orphans --build

run:
	go build -o app cmd/bot/main.go