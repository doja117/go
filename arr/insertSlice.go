package main
import "fmt"
import "strconv"
import "encoding/json"

// slice is the slice to which another slice will be added
// insertion is the slice that needs to be inserted.
// index is the position for insertion
func insertSlice(slice, insertion []string, index int) []string {
    ans:=make([]string,len(slice)+len(insertion))
    copy(ans,slice[0:index])
    copy(ans[index:],insertion)
    copy(ans[index+len(insertion):],slice[index:])
    return ans
}
