package main

//#include "postgres.h"
//#include "utils/builtins.h"
//#include "fmgr.h"
//#include <string.h>
/*
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo CFLAGS:-I/usr/include/postgresql/9.5/server/
*/
import "C"
import (
	"fmt"
	"strconv"
	"time"
)

//export lastWeekday
func lastWeekday() *C.text {
	today := time.Now()
	lastWeekday := today.AddDate(0, 0, -1)
	if lastWeekday.Weekday() == 0 {
		lastWeekday = lastWeekday.AddDate(0, 0, -2)
	} else if lastWeekday.Weekday() == 6 {
		lastWeekday = lastWeekday.AddDate(0, 0, -1)
	}
	return C.cstring_to_text(C.CString(lastWeekday.Format("2006-01-02")))
}

//export nsToTime
func nsToTime(ns uint64) *C.text {
	if ns > 86400000000000 {
		return C.cstring_to_text(C.CString("00:00:00.000000"))
	}
	us := ns / 1000
	hour := us / 3600000000
	us = us - hour*3600000000
	min := us / 60000000
	us = us - min*60000000
	sec := us / 1000000
	us = us - sec*1000000
	return C.cstring_to_text(C.CString(fmt.Sprintf("%02d:%02d:%02d.%06d", hour, min, sec, us)))
}

//export nsStringToTime
func nsStringToTime(nsText *C.text) *C.text {
	nsString := C.GoString(C.text_to_cstring(nsText))
	ns, _ := strconv.ParseUint(nsString,10,64)
	return nsToTime(ns)
}

func main() {
}
