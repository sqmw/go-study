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
