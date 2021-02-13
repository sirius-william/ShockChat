// @Title  command
// @Description 处理命令
// @Author  SiriusWilliam  2021-02-12 22:25
// @Update  2021-02-12 22:25
package terminal

var CommandType map[string]func(key []string)

// 添加一个命令，及其处理函数
func AddCommandFunc(command string, dealingFunc func(key []string)) {
	_, ok := CommandType[command]
	if !ok {
		CommandType[command] = dealingFunc
	}
}

func InitTerminal() {

}
