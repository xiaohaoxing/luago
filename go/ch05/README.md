## LUA 运算符相关

### 1. 算数运算符
| 运算符 | 含义      |
|-----|---------|
| +   | 加法      |
| -   | 减法、一元取反 |
| *   | 乘法      |
| /   | 除法      |
| //  | 整除      |
| %   | 取模      |
| ^   | 乘方      |


一些注意点：

1. 二元运算符，如果都是整数则结果也为整数，否则为浮点数。
2. LUA 中整除运算向下取整，在 Java/Go 中整除是将结果截断，两者不同。
3. 取模/取余可以用整除得到。
4. 乘方计算与其他算数运算符不同，具有右结合性，与字符串拼接相同。


### 2. 位运算符

| 运算符 | 含义          |
|-----|-------------|
| &   | 按位与         |
| \|  | 按位或         |
| ~   | 二元异或，一元按位取反 |
| <<  | 左移          |
| \>> | 右移          |


### 3. 比较运算符

| 运算符 | 含义   |
|-----|------|
| ==  | 等于   |
| ~=  | 不等于  |
| \>  | 大于   |
| \>= | 大于等于 |
| <   | 小于   |
| <=  | 小于等于 |


### 4. 逻辑运算符

| 运算符 | 含义  |
|-----|-----|
| and | 逻辑与 |
| or  | 逻辑或 |
| not | 逻辑非 |

> 一般的语言会给分配专用的运算符（&、|、~），在 lua 中是给了 3 个关键词
 

### 5. 长度运算符 "#"

``` lua
#"hello" => 5
#{1,2,3} => 3
```

### 6. 字符串拼接运算符 ".."

```lua
"a" .. "b" => "ab"
1 .. 2 .. 3 => 123
```