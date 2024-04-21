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
