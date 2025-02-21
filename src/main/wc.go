package main

import "os"
import "fmt"
import "../mapreduce"
import "container/list"
import "unicode"
import "strings"
import "strconv"


func puncCheck( r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsNumber(r)

}


// our simplified version of MapReduce does not supply a
// key to the Map function, as in the paper; only a value,
// which is a part of the input file content. the return
// value should be a list of key/value pairs, each represented
// by a mapreduce.KeyValue



func Map(value string) *list.List {
	var kvplist = list.New()
	var strarr =  strings.FieldsFunc( value ,puncCheck)
	for i:= 0; i< len(strarr) ; i++ {
		var kvp  mapreduce.KeyValue
		kvp.Key = strarr[i]
		kvp.Value =  "1"
		kvplist.PushBack(kvp)
}


	return kvplist
}

// called once for each key generated by Map, with a list
// of that key's string value. should return a single
// output value for that key.
func Reduce(key string, values *list.List) string {

v := 0
for  i := values.Front(); i!= nil ;i = i.Next(){
	v = v+1
}

ans := strconv.Itoa(v)
return ans

}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master x.txt sequential)
// 2) Master (e.g., go run wc.go master x.txt localhost:7777)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	if len(os.Args) != 4 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	} else if os.Args[1] == "master" {
		if os.Args[3] == "sequential" {
			mapreduce.RunSingle(5, 3, os.Args[2], Map, Reduce)
		} else {
			mr := mapreduce.MakeMapReduce(5, 3, os.Args[2], os.Args[3])
			// Wait until MR is done
			<-mr.DoneChannel
		}
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], Map, Reduce, 100)
	}
}
