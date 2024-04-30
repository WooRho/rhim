package middleware

import (
	"github.com/bwmarrin/snowflake"
	"rhim/tools"
	"sync"
)

// UniqueID 唯一标识
type UniqueID snowflake.ID

func (v UniqueID) UInt64() uint64 {
	return uint64(v)
}
func (v UniqueID) UInt() uint {
	return uint(v)
}
func (v UniqueID) Int64() int64 {
	return int64(v)
}

// NodeIface SnowFlakeNode
type NodeIface interface {
	// Init 初始化
	Init()
	// GenerateID 生成UniqueID
	GenerateID() UniqueID
}

// NewCustomNode 自定义
// nodeNum use consts best
// dm := NewCustomNode()
// 实际开发中需要init它 作为全局服务实例
// In actual development, it needs to be init as a global service instance
func NewCustomNode(nodeNum ...int64) NodeIface {
	var num int64
	if len(nodeNum) > 0 {
		num = nodeNum[0]
	}
	node := newCustomNode(num)
	node.Init()
	return node
}

var _ NodeIface = &customNode{}

const (
	defaultNodeBits uint8 = 4 // 节点数，最大 10 位(1-10)，可以有 1024 个节点
	dufaultStepBits uint8 = 8 // 计数序列码，最大 12 位(1-12)，每毫秒产生 4096 个 ID
)

type customNode struct {
	node     *snowflake.Node
	nodeOnce sync.Once
	nodeErr  error
	nodeNum  int64
}

func newCustomNode(nodeNum int64) *customNode {
	return &customNode{
		nodeNum: nodeNum,
	}
}

func (n *customNode) Init() {
	n.nodeOnce.Do(func() {
		snowflake.NodeBits = defaultNodeBits // 4位 最多16个节点
		snowflake.StepBits = dufaultStepBits // 8位 每毫秒产生 255 个 ID（为了兼容js）

		//节点
		if n.nodeNum == 0 {
			n.nodeNum = 1
		}
		n.node, n.nodeErr = snowflake.NewNode(n.nodeNum)
		if n.nodeErr != nil {
			panic(n.nodeErr)
		}
	})
}

func (n *customNode) GenerateID() UniqueID {
	if n.node == nil {
		n.Init()
	}

	id := n.node.Generate()
	n.nodeOnce.Do(func() {
		tools.Persistence("snow", id)
	})

	return UniqueID(id)
}
