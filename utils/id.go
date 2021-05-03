package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"log"
	"os"
)

var currentNode *snowflake.Node
var nodeId int64

func init() {
	nodeId := int64(os.Getpid() % 1024)
	//节点id需要自己分配
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		//LogErr(err, zap.String("reason", "snowflake node error"))
		log.Fatal(err)
	}
	currentNode = node
}

//必须是0-1023
func InitNode(nodeId int64) {
	//节点id需要自己分配
	node, err := snowflake.NewNode(nodeId % 1024)
	if err != nil {
		//LogErr(err, zap.String("reason", "snowflake node error"))
		return
	}
	currentNode = node
}

//snowflake 分布式唯一
func NextId() int64 {
	return currentNode.Generate().Int64()
}

/**
完全随机, 几乎不可能相同
*/
func UuidBinary() []byte {
	data, err := uuid.New().MarshalBinary()
	if err != nil {
		LogErr(err)
		return nil
	}
	return data
}

func UuidText() string {
	data, err := uuid.New().MarshalText()
	if err != nil {
		LogErr(err)
		return ""
	}
	return string(data)
}

func Uuid() uint32 {
	return uuid.New().ID()
}

/**
 mysql n 节点, 每节点 递增n,
0, n, 2n
1, n+1. 2n+1
n-1, 2n-1, 3n-1

redis
snowflake
uuid
*/
