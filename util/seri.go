package util

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
json -> []byte
[]byte -> json
proto -> []byte
[]byte -> proto
*/

func JsonDecode() string {
	bytes, err := json.Marshal(User{Uid: 2222, Name: "fewf", Password: "fewfwss"})
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%s\n", string(bytes))
		log.Printf("proto %s\n", string(bytes))
	}

	return string(bytes)
}

type User struct {
	Uid      uint32 `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"pwd"`
}
