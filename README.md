# HookoGee
learn from Gee

# 第一天 雏形
定义了一个web引擎，即一个名为engine且记录路由的struct，在这个结构体里面封装了handler类型函数( func(http.ResponseWriter, *http.Request) )。
更重要的是该engine实现了ServeHTTP(http.ResponseWriter, *http.Request)；

也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了

# 第二天 上下文

- 初衷1：所有的handler都是有两个参数，无非是根据请求*http.Request，构造响应http.ResponseWriter。这个参数粒度可能有点细了
- 初衷2：对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。

```go
// Context示例
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}
```


