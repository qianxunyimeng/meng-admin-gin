// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/5 19:40:00
// @Desc 路由的辅助信息 例如：是否keepalive,是否固定等

package vo

type RouterHandleVo struct {
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Component string `json:"component"`
}

func (r *RouterHandleVo) SetTitle(title string) {
	r.Title = title
}

func (r *RouterHandleVo) SetIcon(icon string) {
	r.Icon = icon
}

func (r *RouterHandleVo) SetComponent(component string) {
	r.Component = component
}
