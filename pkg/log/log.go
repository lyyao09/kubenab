package log

import "log"

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Printf(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
