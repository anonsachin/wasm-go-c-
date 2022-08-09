#include<iostream>
#include <emscripten.h>

extern "C" {
    int Add(int x, int y);
}

EMSCRIPTEN_KEEPALIVE
int Add(int x, int y){
    return x + y;
}