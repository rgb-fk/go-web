package main

/*
#include <stdio.h>
#include <sys/time.h>
#include <sys/resource.h>

int rlimit_init() {
    printf("setting rlimit\n");

    struct rlimit limit;

    if (getrlimit(RLIMIT_NOFILE, &limit) == -1) {
        printf("getrlimit error\n");
        return 1;
    }

    limit.rlim_cur = limit.rlim_max = 50000;

    if (setrlimit(RLIMIT_NOFILE, &limit) == -1) {
        printf("setrlimit error\n");
        return 1;
    }

    printf("set limit ok\n");
    return 0;
}
*/
import "C"
import (
	"runtime"

	"github.com/everywan/go-web/controller"
)

func main() {
	// go1.5之后已经自动设置MAXPROCS为CPU实际核心数. 此处不再需要
	runtime.GOMAXPROCS(runtime.NumCPU())
	C.rlimit_init()

	controller.StartWebServer()
}
