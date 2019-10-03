package util

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
)

func TestId(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 10; i < 10000; i++ {
		// Generate a snowflake ID.
		id := node.Generate()

		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
		//fmt.Printf("String ID: %s\n", id)
		//fmt.Printf("Base2  ID: %s\n", id.Base2())
		//fmt.Printf("Base64 ID: %s\n", id.Base64())
	}
}

func TestUUID(t *testing.T) {
	for i := 1; i < 1000; i++ {
		id := Uuid()
		t.Logf("%v\n", id)
	}
}

func TestUUIDBINARY(t *testing.T) {
	for i := 1; i < 1000; i++ {
		id := UuidBinary()
		t.Logf("%v %d\n", id, len(id))
	}
}

func TestUUIDText(t *testing.T) {
	for i := 1; i < 1000; i++ {
		id := UuidText()
		t.Logf("%v %s\n", id, id)
	}
}

func TestJsonDecode(t *testing.T) {
	user1 := User{Uid: 10002, Name: "fwefw", Password: "wfs12345"}

	bytes, err := json.Marshal(&user1)
	fmt.Println(string(bytes))
	fmt.Printf("%s, err=%v \n", string(bytes), err)

	var user2 User
	err = json.Unmarshal([]byte(bytes), &user2)

	fmt.Printf("%+v, err=%v \n", user2, err)
}

func TestJsonDecode2(t *testing.T) {
	JsonDecode()
}

func TestRandIntBetween(t *testing.T) {
	n := 10000000
	for i := -n; i < n; i++ {
		rnd1 := -i
		rnd2 := i
		rnd := RandIntBetween(rnd1, rnd1+rnd2)

		if rnd < rnd1 || rnd > rnd1+rnd2 {
			t.Fatalf("rand error %d not in [%d, %d]", rnd, rnd1, rnd2)
		}
	}
}

///////////////////////////////////////Benchmark

func main() {

}
