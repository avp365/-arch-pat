run:
	@mkdir tmp -p; cd ./tmp; cmake .. ;  cmake --build . --clean-first ; ./task

build:
	@mkdir tmp -p; cd ./tmp; cmake .. ; cmake --build .

test:
	@mkdir tmp -p; cd ./tmp; cmake .. ;  cmake --build . --clean-first ; cmake --build . --target test; cat ./Testing/Temporary/LastTest.log 

package:
	@cd ./tmp; cmake --build . --target package	

all:
	@cd ./tmp; cmake ..; cmake --build .;cmake --build . --target test; cmake --build . --target package

