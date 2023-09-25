cmake_minimum_required(VERSION 3.2)

project(task)

find_package(Threads REQUIRED)
find_package(Boost COMPONENTS unit_test_framework REQUIRED)



add_executable(task main.cpp)
add_library(lib sqrteq.cpp)

add_executable(test_sqrteq test_sqrteq.cpp)

set_target_properties(task lib test_sqrteq PROPERTIES
        CXX_STANDARD 17
        CXX_STANDARD_REQUIRED ON
        COMPILE_OPTIONS "-Wpedantic;-Wall;-Wextra"
)

set_target_properties(test_sqrteq PROPERTIES
        COMPILE_DEFINITIONS BOOST_TEST_DYN_LINK
        INCLUDE_DIRECTORIES ${Boost_INCLUDE_DIR}
)

target_link_libraries(task
        lib
        Threads::Threads
 )

target_link_libraries(test_sqrteq
        ${Boost_LIBRARIES}
        lib
)

install(TARGETS task RUNTIME DESTINATION bin)

set(CPACK_GENERATOR DEB)

set(CPACK_PACKAGE_VERSION_MAJOR "1")
set(CPACK_PACKAGE_VERSION_MINOR "1")
set(CPACK_PACKAGE_VERSION_PATCH "1")



set(CPACK_PACKAGE_CONTACT avp365@yandex.ru)

set(CPACK_DEBIAN_PACKAGE_SHLIBDEPS ON)

include(CPack)

enable_testing()

add_test(test_sqrteq test_sqrteq)
