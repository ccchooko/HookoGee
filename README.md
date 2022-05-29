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

# 第三天 route & trie树

1. 梳理总体架构

   - gee.New() 新建一个web框架引擎，该方法完成：

     - 定义一个Engine{}
     - new一个Router{}, 作为Engine{}的参数，Router{}：
         - 包含一个trim树
         - 包含一个制定path的handler的map
   - 执行Engine{}.GET/POST(), 在这个方法执行route的注册，即trie树的插入&handler方法的添加

     - engine.addRoute()  ---->  route.addRouter()  ------> trim.insert()
     - map[method + "-" + pattern] = handler

   - 定义Engine{}.ServeHTTP(), 实现了这个接口，才能自定义自己框架的handler，这个方法完成了：
    
     - 框架context的创建，包含一条请求的所有信息
     - 执行这个请求的handler方法，即type HandlerFunc func(c *Context)

   - 做好上述准备工作之后，Engine{}.Run()，就是实现了http.ListenAndServe(addr, **engine**)，这个方法的执行，必须一览Engine{}.ServeHTTP()的实现

2. trim 树实现
  代码里面关键的部分已经做了注释

# 第四天 分组控制

教程里面这块实现的有些怪怪的，比如，Group嵌入了Engine，Engine又嵌入了Group。参考了评论，改写了下，改写点说明如下：

- Group直接继承了*router，即也同时继承了route的所有方法
- Engine直接继承了*Group，也就是说Engine能像没有加入分组功能时，直接调用Group实现的方法，如GET、POST、addRoute等
- group单独拆出一个文件，个人感觉这样实现更清晰，有空可以在参考参考gin这块的实现