run:
	@go run ./cmd/app

build:
	@cd ./tmp; cmake ..; cmake --build .

test:
	@cd ./tmp; cmake --build . --target test;

package:
	@cd ./tmp; cmake --build . --target package	

all:
	@cd ./tmp; cmake ..; cmake --build .;cmake --build . --target test; cmake --build . --target package

