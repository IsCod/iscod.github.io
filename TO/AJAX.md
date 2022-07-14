# Ajax

Ajax是使用现有标准的新方法，而不是新标准

## Ajax = 异步的Javascript和XML
  
Ajax是一种不需要加载整个页面就可以更新部分页面的技术

Ajax通过在后台与服务器进行少量数据交换，AJAX 可以使网页实现异步更新。这意味着可以在不重新加载整个网页的情况下，对网页的某部分进行更新。

## XMLHttpRequest对象

XMLHttpRequest对象是Ajax应用的基础

所有现代浏览器均支持 XMLHttpRequest 对象（IE5 和 IE6 使用 ActiveXObject）。

XMLHttpRequest 用于在后台与服务器交换数据。这意味着可以在不重新加载整个网页的情况下，对网页的某部分进行更新。

XMLHttpRequest对象:

```javascript
var xmlhttp;
if (window.XMLHttpRequest) {
    xmlhttp = new XMLHttpRequest()
} else {
    //IE6,IE5
    xmlhttp = new ActiveXObject("Microsof.XMLHTTP");
}
xmlhttp.open('GET', '/index.php', false);
xmlhttp.send();
var innerHTML = xmlhttp.responseText;
```

原生的XML包含了大量的代码和对浏览器兼容性的检测，下面使用更简洁的jQuery的方法进行Ajax请求

## jQuery的Ajax方法

**jQuery.ajax()**

jQuery.ajax()方法通过HTTP请求加载远程数据。

```javascript
jQuery.ajax([setting])
```

?> _string_ 可选参数，可通过$.ajaxSetup()设置默认参数

**$.ajaxSetup可设置的常用参数：**
```
- url：发送请求的地址（可选）
- type：请求方法，默认为GET，可设置为POST、PUT、DELETE等（可选）
- data：发送到服务器的数据，默认使用key/value的对象形式，默认转换成相适应的格式，如GET转换为urlencode格式，processData参数可设置转换（可选）
- success：请求成功后回调函数（可选）
- error：请求出错回调函数（可选）
- dataType：返回数据类型，如果没有设置根据HTTP协议中的MIME判断，可能与预期不符（可选）
```

**其它参数：**
```
- cache: 是否缓存页面默认为true,dataType为jsonp,script时默认为false
- content-Type发送信息的内容编码默认为www-x-from-urlencoded，username、password用于HTTP访问认证请求时的用户名和密码等等
```

示例：

```javascript
$.ajax({
  url: '/index.php',
  type: 'POST',
  data: {'name':'myname'},
  success: function(json) {
    alert(json.result);
  },
  error: function() {
    alert('网络错误');
  },
  dataType:'json'
})
```

## 使用封装的get()或post()

jQuery.get() 使用HTTP的GET方法向服务器载入数据

jQuery.post() 使用HTTP的POST方法向服务器载入数据

```javascript
jQuery.get(url, data, success(), dataType);
jQuery.post(url, data, success(), dataType);
```

**参数说明：**

```
- url：发送请求的地址（必须）
- type：请求方法，默认为GET，可设置为POST、PUT、DELETE等（可选）
- data：发送到服务器的数据，默认使用key/value的对象形式，默认转换成相适应的格式，如GET转换为urlencode格式，processData参数可设置转换（可选）
- success：请求成功后回调函数（可选）
- dataType：返回数据类型，如果没有设置根据HTTP协议中的MIME判断，可能与预期不符（可选）
```

该函数就是简写的jQuery.Ajax()函数，等价于：

```javascript
jQuery.ajax({url:url, tyep:type data:data, success:success(), dataType:dataType})
```

## jQuery.get()、jQuery.post()错误处理

**get()、post()的错误处理使用.error()对象完成**

```javascript
$.get().error(function(){});
$.post().error(function(){});
```

**jQuery.get实例**

```javascript
$.get(
  '/index.php',
  {name:'myname'},
  function(json){
    alert(json.result);
  },
  json
).error(function(){
  alert('网络错误');
})
```

**jQuery.posts实例方法**

（该示例与ajax的示例是等价的）：

```javascript
$.post(
  '/index.php',
  {name:'myname'},
  function (json) {
    alert(json.result);
  },
  json
).error(function() {
  alert('网络错误');
})
```