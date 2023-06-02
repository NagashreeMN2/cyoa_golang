package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/NagashreeMN2/cyoa"
)

func main() {

	fileName := flag.String("json", "gopher.json", "JSON file for parsing")
	flag.Parse()

	jsonFile, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	var story cyoa.Story
	_ = json.NewDecoder(jsonFile).Decode(&story)
	fmt.Printf("%+v\n", story)
}
