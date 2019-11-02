package util

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"runtime"
	"time"
)

func PrintOS() {
	//go switch 不需要break。因为自动添加了，但是如果有fallthrough, 则会向下执行多个label
	//case 无须为常量和整数
	//从上置下，短路
	switch os := runtime.GOOS; os {
	case "darwin":
		log.Println("Mac OS X.")
	case "linux":
		log.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		log.Printf("%s.\n", os)
	}
}

func PrintOsInfo() {
	log.Printf("os.Getpagesize() = %#v\n", os.Getpagesize())
	log.Printf("有效 group id os.Getegid() = %#v\n", os.Getegid())
	log.Printf("uid os.Geteuid() = %#v\n", os.Geteuid())
	log.Printf("group id os.Getgid() = %#v\n", os.Getgid())
	ints, err := os.Getgroups()
	if err != nil {
		LogErr(err)
		return
	}
	log.Printf("os.Getgroups() = %#v\n", ints)
	log.Printf("process id os.Getpid() = %#v\n", os.Getpid())
	log.Printf("process caller pid os.Getppid() = %#v\n", os.Getppid())
	log.Printf("os.Getuid() = %#v\n", os.Getuid())
	dir, err := os.Getwd()
	if err != nil {
		LogErr(err)
		return
	}
	log.Printf("os.Getwd() = %#v\n", dir)
	//log.Printf("os.Getpagesize() = %v\n", os.Getegid())
	time.Sleep(5 * time.Minute)
}

func getid(s string) {
	pid := unix.Getpid()
	log.Println(s+" Getpid", pid)
	log.Println(s+" Getppid", unix.Getppid())

	pgid, _ := unix.Getpgid(pid)
	log.Println(s+" Getpgid", pgid)

	//log.Println(s+" Gettid", unix.ge())

	sid, _ := unix.Getsid(pid)
	log.Println(s+" Getsid", sid)

	log.Println(s+" Getegid", unix.Getegid())
	log.Println(s+" Geteuid", unix.Geteuid())
	log.Println(s+" Getgid", unix.Getgid())
	log.Println(s+" Getuid", unix.Getuid())
}
