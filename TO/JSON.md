# JSON

### JSON：Javascript对象表示方法。是储存和交换信息的语法类似XML

JSON具有自我描述性，文本格式在语法上与JavaScript对象代码相同

**一对JSON键/值**
```json
{"key":"value"}
```

**JSON对象实例**

```json
{"name":"username", "password":"password"}
```

**等价的Javascript：**

```js
var name = "username"
var password="password"
```

**JSON数组**

```json
[{"name":"username"}, {"password":"password"}]
```

**等价的Javascript：**

```js
var example = [{"name":"username"}, {"password":"password"}]
```

**PHP通过json_encode对变量进行编码**

```php
json_encode($value, $option, $depth)
```

```
参数：
value：待编译的vaule（必选）
option：可扩展定义，设置为JSON_FORCE_OBJECT时将非关联数组转换成类（可选）
depth：最大深度（可选）
```

在没有设置option为JSON_FORCE_OBJECT时关联数组会转换为JSON为数组，非关联数组会转换为JSON为对象

**php实例：**

```php
<?php
    //非关联数组
    $arr = array('a', 'b', 'c');
    $arr1 = array(0 => 'a', 1 => 'b', '2' => 'c');
    //关联数组
    $arr2 = array(1 => 'a', 2 => 'b', 3 => 'c');
    $arr3 = array('k' => 'a', 'k1' => 'b', 3 => 'c');

    echo 'arr: ' . json_encode($arr) . "\n";
    echo 'arr1: ' . json_encode($arr1) . "\n";
    echo 'arr2: ' . json_encode($arr2) . "\n";
    echo 'arr3: ' . json_encode($arr3) . "\n";
?>
```

**输出：**

```
arr: ["a","b","c"]
arr1: ["a","b","c"]
arr2: {"1":"a","2":"b","3":"c"}
arr3: {"k":"a","k1":"b","3":"c"}
```

添加option为JSON_FORCE_OBJECTS是转换为对象，返回：

```
arr: {"0":"a","1":"b","2":"c"}
arr1: {"0":"a","1":"b","2":"c"}
arr2: {"1":"a","2":"b","3":"c"}
arr3: {"k":"a","k1":"b","3":"c"}
arr4: {"1":1,"2":2,"3":3}
```