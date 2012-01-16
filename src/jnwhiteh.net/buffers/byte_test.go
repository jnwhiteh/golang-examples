package buffers

import "testing"

func TestByteBuffer(t *testing.T) {
	var buf *ByteBuffer = NewByteBuffer()

	buf.WriteString("Hello ")
	buf.WriteString("World")
	buf.WriteString("!")

	res := buf.GetString()
	expected := "Hello World!"

	if expected != res {
		t.Errorf("Expected '%s' and got '%s'", expected, res)
	}
}
