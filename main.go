package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hashicorp/hcl2/hclwrite"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

var (
	source = flag.String("source", "", "--source=./path/to/source.json")
)

func inputFromReader(r io.Reader) (map[string]interface{}, error) {

	var out map[string]interface{}

	dec := json.NewDecoder(r)
	for {
		err := dec.Decode(&out)
		if err == io.EOF {
			break
		}

		return out, err
	}

	return out, nil
}

func main() {

	flag.Parse()

	// default to read the data from stdin
	inputSource := os.Stdin

	if *source != "" {
		var err error
		inputSource, err = os.Open(*source)
		if err != nil {
			log.Fatal("Unable to open source file", err)
		}
	}

	input, err := inputFromReader(inputSource)
	if err != nil {
		log.Fatal("Unable to load source", err)
	}

	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	for k, v := range input {
		raw, _ := json.Marshal(v)

		simple := &ctyjson.SimpleJSONValue{}

		json.Unmarshal(raw, simple)
		rootBody.SetAttributeValue(k, simple.Value)
	}

	fmt.Fprintf(os.Stdout, "%s", hclwrite.Format(f.Bytes()))
}
