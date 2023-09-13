run:
	@cd ./tmp;./task

build:
	@mkdir tmp -p;cd ./tmp; cmake ..; cmake --build .

test:
	@cd ./tmp; cmake --build . --target test;

package:
	@cd ./tmp; cmake --build . --target package	

all:
	@cd ./tmp; cmake ..; cmake --build .;cmake --build . --target test; cmake --build . --target package

