#### go语言的嵌入类型
这是一种可以把已有的类型声明在新的类型里的一种方式.<br/>
在其他语言中(如java), 他们使用类之间的继承来做这些事情, 但是在go语言中, 没有继承的概念. go提倡使用组合的方式进行代码的复用.

##### go嵌入类型的初始化
go类型的嵌入类型的初始化可参考main.go.

##### go嵌入类型方法覆盖
示例详见funcOverride.go.<br/>
示例中, Admin内嵌入了User. User类有一个方法SayHello, 同时Admin实现了自己的SayHello方法.此时就形成了方法的覆盖(Admin.SayHello()
覆盖了User.SayHello()).

##### go嵌入类型实现接口
示例代码内声明了Hello接口. User类实现了Hello接口的方法, Admin类内嵌入了User类, 此时可以认为Admin也同样实现了Hello接口(Admin没有
声明自己其它特殊的方法), 同事Admin也可以显示实现自己的逻辑.