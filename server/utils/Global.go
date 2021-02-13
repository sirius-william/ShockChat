// @Title  Global.go
// @Description	定义了一些全局变量
// @Author  SiriusWilliam  2021-01-22 16:54
// @Update  2021-01-22 16:54
package utils

import (
	"ShockChatServer/logger"
	"github.com/aceld/zinx/ziface"
)

// 连接池地址池长度，在此基础上进行链表拉链展开
const count int32 = 1000

/*
 * 用于存储用户id和connId的映射关系，用户发消息时使用
 */
// 哈希表中的链表节点
type userNode struct {
	userid         int32
	userConnection ziface.IConnection
	next           *userNode
	pre            *userNode
}

// 哈希表的地址池
type userPool struct {
	count int32
	pool  []*userNode
}

var ConnectionIdReflectorZinxConnID *userPool = &userPool{0, make([]*userNode, count)}

func InitVars() {
	// 初始化Global中，用户id和connId的对于关系的容器
	logger.Log.INFO("Init ConnectionPool HashTable.")
	pool := make([]*userNode, count)
	for i := int32(0); i < 1000; i++ {
		pool[i] = &userNode{i, nil, nil, nil}
	}
	ConnectionIdReflectorZinxConnID.pool = pool
}

func (u *userPool) Add(userid int32, conn ziface.IConnection) {
	var index int32 = hash(userid)
	temp := u.pool[index]
	for temp != nil {
		if temp.next == nil {
			temp.next = &userNode{userid: userid, next: nil, pre: temp, userConnection: conn}
			u.count += 1
			break
		}
		temp = temp.next
	}
}

//func (u *userPool) Print() {
//	temp := ConnectionIdReflectorZinxConnID.pool[0]
//	for temp != nil {
//		fmt.Printf("%p\tnext:%p\tpre:%p\n", temp, temp.next, temp.pre)
//		temp = temp.next
//	}
//}

func (u *userPool) Delete(userid int32) {
	var index int32 = hash(userid)
	temp := u.pool[index]
	for temp != nil {
		if temp.userid == userid {
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			u.count -= 1
			break
		}
		temp = temp.next
	}
}

func (u *userPool) GetConnection(userid int32) ziface.IConnection {
	var index int32 = hash(userid)
	temp := u.pool[index]
	for temp != nil {
		if temp.userid == userid {
			return temp.userConnection
		}
	}
	return nil
}

// 一个简单的哈希算法
func hash(userid int32) int32 {
	return userid % count
}

func (u *userPool) IsLogin(userid int32) bool {
	var index int32 = hash(userid)
	temp := u.pool[index]
	for temp != nil {
		if temp.userid == userid {
			return true
		}
	}
	return false
}

func (u *userPool) Count() int32 {
	return u.count
}
