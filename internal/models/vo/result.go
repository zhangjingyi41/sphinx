package vo

type common struct {
	code int
	msg  string
	data interface{}
}

type failed struct {
	code int
	msg  string
	data interface{}
}
type success struct {
	code int
	msg  string
	data interface{}
}

// 只有Result有对外访问权限
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 构建基础的响应结构体
func Builder() *common {
	return &common{}
}

// 响应成功
func (c *common) Ok() *success {
	r := &success{}
	r.code = 0 // 默认成功状态码为0
	return r
}

// 响应信息
func (c *success) Msg(msg string) *success {
	c.msg = msg
	return c
}

// 整合响应数据结构体
func (c *failed) Finished() *Result {
	r := new(Result)
	r.Code = c.code // int的零值是0，所以不需要判断是否为空，如果是响应成功，用默认零值即可
	// 检查msg是否为空
	if c.msg == "" {
		c.msg = "success"
	}
	r.Msg = c.msg
	r.Data = c.data
	return r
}

// 整合响应数据结构体
func (c *success) Finished() *Result {
	r := new(Result)
	r.Code = c.code // int的零值是0，所以不需要判断是否为空，如果是响应成功，用默认零值即可
	r.Msg = c.msg
	r.Data = c.data
	return r
}

// 响应失败
func (c *common) Interrupted() *failed {
	r := &failed{}
	r.code = -1 // 默认错误状态码为-1
	return r
}

// 状态码
func (c *failed) Code(code int) *failed {
	c.code = code
	return c
}

// 响应失败信息
func (c *failed) Reason(msg string) *failed {
	c.msg = msg
	return c
}

// 响应失败数据
func (c *success) Data(data interface{}) *success {
	c.data = data
	return c
}
