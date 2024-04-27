# go-study

## 命名

- 变量(一般变量、函数)命名
    - 函数大写
    - 一般变量，小驼峰

## 关键字

- var
    - 用来在任意位置申明变量，没有赋值就是 0
    - 用来在一行申明多个变量的时候，只能申明同类的变量(这样符合逻辑)
    - 用法
        - `var variable_name variable_type`
        - `var variable_name [variable_type] = value`
- :=
    - 用来在函数体里面申明并且赋值


- import
    - 用来引入外部 package
    - 用法(下面的两种方法是一样的)
        ```
        import "math"
        import "fmt"
        ```
        ```
        import (
            "fmt"
            "math"
        )
        ```

- const(还需要看看类型转化的go的规则)
    - 使用 const 操作数字的时候，如果没有申明具体的类型
    - 那么使用的时候会自己推断，或者我们可以进行类型强转

- range
    - 有点和 Python 里面的 enumerate 相似
    - 在 go 里面使用 range 来循环只是一个简单的补充而已，和Python不一样
- chan
    - 表示 channel 需要指定 channel 里面传递的数据的类型
    - `val, more := <- chanA`
        - more == False <==> channel 关闭&&为空，意思是表示的是可不可能有更多的数据
    - `close(chanA)`
        - 表示关闭 channel
        - close 之后使用 range 遍历的时候就不会报错了(应该是使用了close之后，range就找到了结束的标识位，而不是盲目range了)
        - channel 的关闭仅仅是 send 而已，对于已经在 channel 里面的数据，是可以进行接收的

## packages

### package time

#### NewTimer()

- 仅仅执行一次
- <- t.C 是一个管道，我们接收到的值是想在的时间

#### NewTicker()

- 可以执行多次
- <- t.C 是一个管道，我们接收到的值是想在的时间
