package xsnowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/util/grand"
)

var node *snowflake.Node
var err error

func init() {
	node, err = snowflake.NewNode(int64(grand.N(1, 1023)))
	if err != nil {
		panic(err)
	}
}

func Int64() int64 {
	return node.Generate().Int64()
}
