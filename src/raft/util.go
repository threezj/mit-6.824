package raft

import "log"

// Debugging
const Debug = -3

func Log(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}


func Log2(format string, a ...interface{}) (n int, err error) {
	if Debug > -1 {
		log.Printf(format, a...)
	}
	return
}

func Log3(format string, a ...interface{}) (n int, err error) {
	if Debug > -2 {
		log.Printf(format, a...)
	}
	return
}
