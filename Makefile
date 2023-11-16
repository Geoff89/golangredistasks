worker:
	go run ./worker/server.go

client:
	go run ./client/main.go	

.PHONY:	worker client
