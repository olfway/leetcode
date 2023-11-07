import (
    "bytes"
    "unsafe"
)

var buffer1 = make([]byte, 0, 2000)
var buffer2 = make([]byte, 0, 2000)

func gcdOfStrings(str1 string, str2 string) string {

    a, b := len(str1), len(str2)
    ab := a+b

    slice1 := unsafe.Slice(unsafe.StringData(str1), a)
    slice2 := unsafe.Slice(unsafe.StringData(str2), b)

    buffer1 = buffer1[:ab]
    copy(buffer1, slice1)
    copy(buffer1[a:], slice2)

    buffer2 = buffer2[:ab]
    copy(buffer2, slice2)
    copy(buffer2[b:], slice1)

    if bytes.Equal(buffer1, buffer2) != true {
        return ""
    }

    for b != 0 {
        a, b = b, a % b
    }

    return str1[:a]
}
