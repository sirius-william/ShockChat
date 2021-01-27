// @Title  Global.go
// @Description	定义了一些全局变量
// @Author  SiriusWilliam  2021-01-22 16:54
// @Update  2021-01-22 16:54
package utils

/*
 * 用于存储用户id和connId的映射关系，用户发消息时使用
 */
var ConnectionIdReflectorZinxConnID map[int32]int32

func InitVars() {
	// 初始化Global中，用户id和connId的对于关系的容器
	ConnectionIdReflectorZinxConnID = make(map[int32]int32)
}
