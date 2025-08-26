package raft

import "log"

// Debugging
const Debug = 0

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}

// DPrintLog logs the Raft state for debugging
func DPrintLog(rf *Raft) {
	if Debug > 0 {
		log.Printf("[term %d]: Raft[%d] [state %d] log: %+v", rf.currentTerm, rf.me, rf.state, rf.log)
	}
}
