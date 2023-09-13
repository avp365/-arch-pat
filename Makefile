run:
	@go run ./cmd/app

build:
	@cd ./tmp; cmake ..; cmake --build .;cmake --build . --target test; cmake --build . --target package

