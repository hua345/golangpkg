### 参考

- [Golang Context分析](https://www.jianshu.com/p/e5df3cd0708b)

#### Context适用场景

比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，
这些goroutine又可能会开启其他的goroutine。这样的话， 我们就可以通过Context，
来跟踪这些goroutine，并且通过Context来控制他们的目的，
这就是Go语言为我们提供的Context，中文可以称之为“上下文”。

另外一个实际例子是，在Go服务器程序中，每个请求都会有一个goroutine去处理。
然而，处理程序往往还需要创建额外的goroutine去访问后端资源，比如数据库、RPC服务等。
由于这些goroutine都是在处理同一个请求，所以它们往往需要访问一些共享的资源，
比如用户身份信息、认证token、请求截止时间等。而且如果请求超时或者被取消后，
所有的goroutine都应该马上退出并且释放相关的资源。
这种情况也需要用Context来为我们取消掉所有goroutine

#### Context 使用原则 和 技巧

- 不要把Context放在结构体中，要以参数的方式传递，parent Context一般为Background
- 应该要把Context作为第一个参数传递给入口请求和出口请求链路上的每一个函数，放在第一位，变量名建议都统一，如ctx。
- Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
- Context是线程安全的，可以放心的在多个goroutine中传递
- 可以把一个 Context 对象传递给任意个数的 `gorotuine`，对它执行`取消`操作时，所有 goroutine 都会接收到取消信号。
