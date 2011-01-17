package main

import "jnwhiteh.net/buffers"
import "fmt"

func main() {
	buf := buffers.NewByteBuffer()

	buf.WriteString("Hello ")
	buf.Write([]byte("World!"))

	fmt.Println(buf.GetString())
}
