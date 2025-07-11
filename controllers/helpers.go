package controllers

import "strconv"

func Atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
