package main

func main() {
// tag::example[]
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 is []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 is []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 is []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 is []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t is []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b is []byte{'b', 'a', 'r' }
// end::example[]
}
