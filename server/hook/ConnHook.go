// @Title  ConnHook.go
// @Description
// @Author  SiriusWilliam  2021-01-22 17:12
// @Update  2021-01-22 17:12
package hook

import (
	"ShockChatServer/utils"
	"ShockChatServer/utils/redis"
	"fmt"
	"github.com/aceld/zinx/ziface"
)

/*
 * 在连接断开销毁时，将id的映射关系从容器中删除
 */
func AfterConnectionStopped(conn ziface.IConnection) {
	pro, err := conn.GetProperty("id")
	if err != nil {
		fmt.Println(err)
		return
	}
	id := pro.(int32)
	delete(utils.ConnectionIdReflectorZinxConnID, id)
	_, _ = redis.Exec("hdel", "token", pro)
}
