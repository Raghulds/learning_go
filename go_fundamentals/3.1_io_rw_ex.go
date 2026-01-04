package gofundamentals

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Reader  →  gives bytes
// Writer  →  takes bytes

// | Package / Type        | **Read does**       | **Reads FROM (source)**              | **Write does**  | **Writes TO (destination)**          |
// | --------------------- | ------------------- | ------------------------------------ | --------------- | ------------------------------------ |
// | `strings.Reader`      | Copies bytes        | Go string (memory, immutable)        | ❌               | ❌                                    |
// | `bytes.Buffer`        | Copies bytes        | Internal byte slice (memory)         | Appends bytes   | Internal byte slice (memory)         |
// | `bytes.Reader`        | Copies bytes        | Fixed byte slice (memory)            | ❌               | ❌                                    |
// | `os.File`             | Syscall `read`      | File on disk (via OS kernel)         | Syscall `write` | File on disk (via OS kernel)         |
// | `net.Conn`            | Socket receive      | Kernel socket buffer (network)       | Socket send     | Kernel socket buffer (network)       |
// | `http.Request.Body`   | Stream read         | TCP request body (client → server)   | ❌               | ❌                                    |
// | `http.ResponseWriter` | ❌                   | ❌                                    | Stream write    | TCP response (server → client)       |
// | `gzip.Reader`         | Decompresses        | Underlying Reader (compressed bytes) | ❌               | ❌                                    |
// | `gzip.Writer`         | ❌                   | ❌                                    | Compresses      | Underlying Writer (compressed bytes) |
// | `json.Decoder`        | Tokenizes & decodes | Underlying Reader (JSON bytes)       | ❌               | ❌                                    |
// | `json.Encoder`        | ❌                   | ❌                                    | Encodes JSON    | Underlying Writer                    |
// | `os.Stdout`           | ❌                   | ❌                                    | Writes bytes    | Terminal / stdout stream             |
// | `os.Stderr`           | ❌                   | ❌                                    | Writes bytes    | Terminal / stderr stream             |

// Read from a string
func ReadFromStr(readString string) {
	r := strings.NewReader(readString)
	buf := make([]byte, 10)

	for {
		n, err := r.Read(buf)
		fmt.Println("n=%d buf=%q err=%w\n", n, buf[:n], err)

		if err == io.EOF {
			break
		}
	}
}

// Write to a buffer
func WriteToBuffer(text string) {
	var buffer bytes.Buffer // Implements Writer

	n, err := buffer.Write([]byte(text))
	if err != nil {
		fmt.Errorf("Error while writing - %w", err)
	}

	fmt.Println("Written %d", n)
	buffer.WriteString(" written")

	fmt.Println(buffer.String())
}

func CopyData(textToCopy string) {
	src := strings.NewReader(textToCopy)
	dest := &bytes.Buffer{}

	n, _ := io.Copy(dest, src)
	fmt.Println("Bytes copied %d", n)
	fmt.Println("Dest - ", dest.String())
}

// Custom Reader
type MyReader struct{}

func (r *MyReader) Read(data []byte) (int, error) {
	dataToCopy := "hi"
	copy(data, dataToCopy)
	return len(dataToCopy), io.EOF
}

func ReadWithCustomReader() {
	// io.Copy uses MyReader's Read
	io.Copy(os.Stdout, &MyReader{})
}

// Custom Writer
type LogWriter struct{}

func (w *LogWriter) Write(data []byte) (int, error) {
	fmt.Println("writing:", string(data))
	return len(data), nil
}

func WriteWithCustomWriter() {
	// io.Copy uses LogWriter's Write
	io.Copy(&LogWriter{}, strings.NewReader("writing <>"))
}
