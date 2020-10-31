# Go parameters
MAIN_PATH=cmd/deliver-app/main.go
BINARY_NAME=$(BINARY_PATH)/app
BINARY_PATH=bin

run:
	go build -o $(BINARY_NAME) -race $(MAIN_PATH)
	./$(BINARY_NAME)

clean:
	go clean $(MAIN_PATH)
	rm -f $(BINARY_PATH)/*