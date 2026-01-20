.PHONY: install dev build test clean

install:
	@echo "Installing backend dependencies..."
	go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

dev-backend:
	@echo "Starting Go backend..."
	go run main.go

dev-frontend:
	@echo "Starting Vue frontend..."
	cd frontend && npm run dev

dev:
	@echo "Starting both backend and frontend..."
	# Using make -j to run in parallel
	make -j 2 dev-backend dev-frontend

build:
	@echo "Building frontend..."
	cd frontend && npm run build
	@echo "Building backend..."
	go build -o bin/server main.go

clean:
	rm -rf frontend/dist
	rm -rf bin/
