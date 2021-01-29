cmake_minimum_required(VERSION 3.5)

project(ShockChat_Qt LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# QtCreator supports the following variables for Android, which are identical to qmake Android variables.
# Check https://doc.qt.io/qt/deployment-android.html for more information.
# They need to be set before the find_package( ...) calls below.

#if(ANDROID)
#    set(ANDROID_PACKAGE_SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/android")
#    if (ANDROID_ABI STREQUAL "armeabi-v7a")
#        set(ANDROID_EXTRA_LIBS
#            ${CMAKE_CURRENT_SOURCE_DIR}/path/to/libcrypto.so
#            ${CMAKE_CURRENT_SOURCE_DIR}/path/to/libssl.so)
#    endif()
#endif()

find_package(QT NAMES Qt6 Qt5 COMPONENTS Widgets LinguistTools REQUIRED)
find_package(Qt${QT_VERSION_MAJOR} COMPONENTS Widgets LinguistTools REQUIRED)
find_package(Qt6 COMPONENTS Network REQUIRED)

include_directories("C:/openssl/include")
include_directories("F:/protobuf-3.14.0/src")
link_directories("C:/openssl/lib")
link_directories("D:/project/ShockChat/client")
#link_directories("F:/protobuf-3.14.0/src/build/Debug")
set(TS_FILES ShockChat_Qt_zh_CN.ts)

aux_source_directory(./protos myprotos)
aux_source_directory(./ src)
message("protos:${myprotos}")
set(PROJECT_SOURCES
        main.cpp
        mainwindow.cpp
        mainwindow.hpp
	registerwindow.hpp
	netthread.hpp
	Message.hpp
        mainwindow.ui
	registerwindow.ui
        ${TS_FILES}
	${myprotos}
	${src}
)
# 定义一些关于socket连接的宏
add_definitions(-DHOST="127.0.0.1")
add_definitions(-DPORT=8999)
add_definitions(-DTIME_OUT=800000)

if(${QT_VERSION_MAJOR} GREATER_EQUAL 6)
    qt_add_executable(ShockChat_Qt
        ${PROJECT_SOURCES}
    )

    qt_create_translation(QM_FILES ${CMAKE_SOURCE_DIR} ${TS_FILES})
else()
    if(ANDROID)
        add_library(ShockChat_Qt SHARED
            ${PROJECT_SOURCES}
        )
    else()
        add_executable(ShockChat_Qt
            ${PROJECT_SOURCES}
        )
    endif()

    qt5_create_translation(QM_FILES ${CMAKE_SOURCE_DIR} ${TS_FILES})
endif()

target_link_libraries(ShockChat_Qt PRIVATE Qt${QT_VERSION_MAJOR}::Widgets libssl libcrypto libprotobuf Qt6::Network)
#target_link_libraries(ShockChat_Qt PRIVATE Qt${QT_VERSION_MAJOR}::Widgets libssl libcrypto libprotobufd Qt6::Network)
