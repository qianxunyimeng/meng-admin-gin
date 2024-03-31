// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/31 23:13:00
// @Desc
package service

import "meng-admin-gin/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
