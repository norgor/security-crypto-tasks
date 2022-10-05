cmake_minimum_required(VERSION 3.1)

project(5)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c11 -Wall -Wextra")

add_executable(main main.c)

enable_testing()
add_executable(utility_fuzzer_test utility_fuzzer_test.c)
target_compile_options(utility_fuzzer_test PRIVATE -fsanitize=address,fuzzer)
target_link_options(utility_fuzzer_test PRIVATE -fsanitize=address,fuzzer)
