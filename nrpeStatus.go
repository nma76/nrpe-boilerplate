package main

type ReturnStatusEnum int

// All available return codes
const (
	OK       ReturnStatusEnum = 0
	WARNING  ReturnStatusEnum = 1
	CRITICAL ReturnStatusEnum = 2
	UNKNOWN  ReturnStatusEnum = 3
)
