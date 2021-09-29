package structool_test

import (
	"fmt"
	"time"

	"github.com/RussellLuo/structool"
)

func Example_decode() {
	in := map[string]interface{}{
		"string":   "s",
		"bool":     true,
		"int":      1,
		"error":    "oops",
		"time":     "2021-09-29T00:00:00Z",
		"duration": "2s",
	}
	out := struct {
		String   string        `structool:"string"`
		Bool     bool          `structool:"bool"`
		Int      int           `structool:"int"`
		Error    error         `structool:"error"`
		Time     time.Time     `structool:"time"`
		Duration time.Duration `structool:"duration"`
	}{}

	codec := structool.New().DecodeHook(
		structool.DecodeStringToError,
		structool.DecodeStringToTime(time.RFC3339),
		structool.DecodeStringToDuration,
	)
	if err := codec.Decode(in, &out); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", out)

	// Output:
	// {String:s Bool:true Int:1 Error:oops Time:2021-09-29 00:00:00 +0000 UTC Duration:2s}
}

func Example_encode() {
	in := struct {
		String   string        `structool:"string"`
		Bool     bool          `structool:"bool"`
		Int      int           `structool:"int"`
		Error    error         `structool:"error"`
		Time     time.Time     `structool:"time"`
		Duration time.Duration `structool:"duration"`
	}{
		String:   "s",
		Bool:     true,
		Int:      1,
		Error:    fmt.Errorf("oops"),
		Time:     time.Date(2021, 9, 29, 0, 0, 0, 0, time.UTC),
		Duration: 2 * time.Second,
	}

	codec := structool.New().EncodeHook(
		structool.EncodeErrorToString,
		structool.EncodeTimeToString(time.RFC3339),
		structool.EncodeDurationToString,
	)
	out, err := codec.Encode(in)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", out)

	// Output:
	// map[string]interface {}{"bool":true, "duration":"2s", "error":"oops", "int":1, "string":"s", "time":"2021-09-29T00:00:00Z"}
}
