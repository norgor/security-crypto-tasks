build:
	mkdir -p out

	# 1
	nasm -f elf64 1/hello.s -o out/hello.o
	ld out/hello.o -o out/1

	# 2
	gcc 2/main.c -o out/2_c
	g++ 2/main.cpp -o out/2_cpp
	rustc 2/main.rs -o out/2_rs

	# 3
	gcc -c -fPIC 3/a_function.c -o out/a_function.o
	gcc -c -fPIC 3/more_functions.c -o out/more_functions.o
	gcc -shared out/a_function.o out/more_functions.o -o out/libfunctions.so	
	gcc 3/main.c -L$(PWD)/out -lfunctions -o out/3
