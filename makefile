GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
  GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=mybinary
    BINARY_UNIX=$(BINARY_NAME)_unix
    
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
  test: 
            $(GOTEST) -v ./...
    clean: 
            $(GOCLEAN)
            rm -f $(BINARY_NAME)
            rm -f $(BINARY_UNIX)
    run:
            $(GOBUILD) -o $(BINARY_NAME) -v ./...
            ./$(BINARY_NAME)
    deps:
            $(GOGET) github.com/boltdb/bolt
            $(GOGET) github.com/gorilla/csrf
    


