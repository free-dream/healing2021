package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

/**【未测试】
 * @Description 通过用户 id 返回该用户的所有信息
 * @Param 用户 id
 * @return 含有该用户的所有信息的结构体，判断数据库操作是否成功的ok(true说明成功)
 **/
func GetUserById(Id int) (statements.User, bool) {
	MysqlDB := setting.MysqlConn()
	var OneUser statements.User
	if err := MysqlDB.Where("id=?", Id).First(&OneUser); err != nil {
		return OneUser, false
	}
	return OneUser, true
}
