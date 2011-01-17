package main

// If you absolutely NEEDED to import the local version of the package rather
// than installing it as recommended, you can use the following import instead

// import "../../pkg/buffers/_obj/jnwhiteh.net/buffers

import "jnwhiteh.net/buffers"
import "fmt"

func main() {
	buf := buffers.NewByteBuffer()

	buf.WriteString("Hello ")
	buf.Write([]byte("World!"))

	fmt.Println(buf.GetString())
}
