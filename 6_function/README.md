### 函数/方法相关学习笔记
* 函数 <br/>
    - 函数定义<br/>
    函数和方法定义非常相似. 函数的定义声明没有接收者, 所以一般直接写字go文件里.类似java的静态方法
    - 函数的作用域<br/>
    函数的作用域按照函数名的首字母大小写来区分. 小写代表该函数的作用域只属于所声明的包内使用, 不能被其他包使用.
    如果函数名以大写开头, 该函数的作用域就可以被其他包调用.类型的属性名也类似.

* 方法<br/>
    - 方法定义<br/>
    方法的声明和函数类似, 他们的区别是: 方法在定义的时候会在func和方法名之间增加一个参数, 这个参数就是接收者, 这样我们定义的
    这个方法就和接收者绑定在了一起, 称之为这个接收者的方法.
