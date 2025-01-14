package generate

type Access struct {
	Access map[string]*AccessInfo //授权组
}

type AccessInfo struct {
	Name       string    `json:"name"`       // 请求名称
	Code       string    `json:"code"`       // 请求名称
	Group      string    `json:"group"`      // 授权组
	IsNeedAuth int64     `json:"isNeedAuth"` // 是否需要认证（1是 2否）
	Desc       string    `json:"desc"`       // 备注
	Apis       []ApiInfo `json:"apis"`       //授权组下的接口
	AuthType   string    `json:"authType"`   // 1(all) 全部人可以操作 2(admin) 默认授予租户管理员权限 3(superAdmin,supper) default租户才可以操作(超管是跨租户的)
}

type ApiInfo struct {
	AccessCode   string `json:"accessCode"`   // 范围编码
	Method       string `json:"method"`       // 请求方式（1 GET 2 POST 3 HEAD 4 OPTIONS 5 PUT 6 DELETE 7 TRACE 8 CONNECT 9 其它）
	Route        string `json:"route"`        // 路由
	Name         string `json:"name"`         // 请求名称
	IsNeedAuth   int64  `json:"isNeedAuth"`   // 是否需要认证（1是 2否）
	BusinessType string `json:"businessType"` // 业务类型（1(add)新增 2修改(modify) 3删除(delete) 4查询(find) 5其它(other)
	Desc         string `json:"desc"`         // 备注
	AuthType     string `json:"authType"`     // 1(all) 全部人可以操作 2(admin) 默认授予租户管理员权限 3(superAdmin,supper) default租户才可以操作(超管是跨租户的)
}
