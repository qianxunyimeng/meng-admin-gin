// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/5/5 19:39:00
// @Desc

package vo

type RouterVo struct {
	Path     string         `json:"path"`
	Handle   RouterHandleVo `json:"handle"`
	Children []RouterVo     `json:"children"`
}

func (r *RouterVo) SetPath(path string) {
	r.Path = path
}

func (r *RouterVo) SetHandle(handle RouterHandleVo) {
	r.Handle = handle
}

func (r *RouterVo) SetChildren(children []RouterVo) {
	r.Children = children
}
