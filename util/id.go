package util

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var node *snowflake.Node

func init() {
	var err error
	//节点id需要自己分配
	node, err = snowflake.NewNode(1)
	if err != nil {
		LogErr(err, zap.String("reason", "snowflake node error"))
		return
	}
}

//snowflake 分布式唯一
func NextId() int64 {
	return node.Generate().Int64()
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

func UuidText() []byte {
	data, err := uuid.New().MarshalText()
	if err != nil {
		LogErr(err)
		return nil
	}
	return data
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
