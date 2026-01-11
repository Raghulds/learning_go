package gofundamentals

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CallGithub() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	fmt.Println(userInfo(ctx, "raghulds"))
}

func userInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad status. Error: ", err)
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

// io.Reader - Any input stream
func parseResponse(r io.Reader) (string, int, error) {
	// anonymous struct
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(r)
	// Decode needs a pointer to modify it's value
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("Error while decoding: ", err)
		return "", 0, err
	}

	return reply.Name, reply.NumRepos, nil
}

// JSON <-> Go
/*
string <> string
true/false <> bool
number <> float64, float32, int, int8 ... int64, uint, uint8 ...
array <> []T, []any
object <> map[string]any, struct

encoding/json API
JSON -> []byte -> Go: Unmarshal
Go -> []byte -> JSON: Marshal
JSON -> io.Reader -> Go: Decoder
Go -> io.Writer -> JSON: Encoder
*/

/*
Inside Decode, Go uses reflection to
* read JSON fields,
* match with struct,
* write into memory of struct
*/

/*
Reflection
------------
Ideas:
* Service/Endpoint reflection (API discovery)
	Exposing service metadata for grpcurl, IDEs..
* Message serialization (protobuf encoding/decoding)

Protobuf uses code generation & avoids reflection
*/
