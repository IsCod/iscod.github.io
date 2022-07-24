# MySQL

## 性能剖析

性能定义为完成某件任务所需要的时间度量，换句话说就，性能即响应时间。而捕捉应用程序和MySQL查询的响应时间就成了性能优化的第一步。

### 应用程序性能剖析

应用剖析需要结合你所使用的编程语言选择合适的分析工具如：

编程语言|分析工具
------|------
php|xhprof
go|pprof


### MySQL查询剖析

#### 慢日志

在MySQL当前版本中，慢日志是开销最低、精度最高的测量查询时间的工具, 在I/O密集型场景下的基准测试表明, 慢日志带来的开销基本可以忽略不计


* 查看慢日志是否开启

```sql
show variables like '%query%'
```

* 设置慢日志临时开启和慢查询时间临界点

```sql
set global slow_query_log = 1; # 开启慢日志
set long_query_time = 2; # 时间临界值2s,可以设置为0记录所有查询
```

通过查看慢日志我们就能捕捉到相应时间较慢的查询语句, 进而针对该语句进行优化。
不过有时候因为某些原因如权限不足等, 无法在服务器上记录查询日志。

#### 查询状态

* 使用`SHOW processlist`

对此我们可以使用不间断的查询`SHOW FULL processlist`来记录查询第一次出现和消失的时间, 很多时候这样的精度也可以发现问题

```bash
mysql -proot -e 'show processlist'
```

![profiles](https://iscod.github.io/images/mysql_processlist.png)

`show processlist`中Command列表示当前语句执行的状态，

- Sleep

线程正在等待客户发送新的请求

- Query

线程正在执行查询或者正在将结果发送给客户端

- Locked

该线程正在等待锁表

- Analyzing and Statistics

线程正在收集存储引擎的统计信息, 并生成查询的执行计划

- Copying to tmp table [on disk]

线程正在执行查询, 并且将其结果集复制到一个临时表中。这种状态一般要么是在做`GROUP BY`操作, 要么是文件排序操作, 或者是`UNION`操作。
如果这个状态后面还有`on disk`标记, 那么表示MySQL正在讲一个内存临时表放到磁盘上。

- Sorting result

线程正在对结果集进行排序

- Sending data

线程在多个状态之间传送数据, 或者在生成结果集, 或者在向客户端返回数据

#### 查询响应时间

* 使用`SHOW PROFILES`

profiles默认是禁用的, 可以通过修改profiling开启

```sql
set profiling = 1 # 开启profiles记录
```

```sql
show variables like "pro%";
```

`profiles`会在一个查询提交给服务器时, 将剖析休息记录到一张临时表, 并给查询赋予一个从1开始的标识符。

```sql
show profiles;
```

![profiles](https://iscod.github.io/images/mysql_profiles_1.png)

可以看到`profiles`以很高的精度显示了查询的响应时间

## 数据类型优化

MySql支持的数据类型非常多, 选择正确的数据类型对获取高性能至关总要, 不管存储哪些类型的数据, 有几个简单的原则可以帮助你做出更好的选择

- 更小的通常更好:

一般情况下, 应该选择尽量使用可以正确存储数据的最小类型数据, 更小的数据类型通常更快, 因为它们占用的磁盘, 内存, CPU缓存更小, 并CPU处理时需要的周期也更短

- 简单就好

简单的数据类型的操作通常需要更少的CPU周期。例如, 整型比字符串操作代价更低, 因为字符集多校对规则（排序规则）使字符串比整型比较更复杂。

典型的两个例子是: 应该用```整型存储IP地址```和使用```MySQL内建类型存储日期和时间```而不是字符串

> IPV4实际是一个32位无符号整数, 而不是字符串。用小数点将其分成四段只是为了人们阅读容易。`inet_aton()` 和`inet_ntoa()`是MySql内置函数将IPV4处理成32位无符号的整数

- 尽量避免`NULL`

很多表都包含可未`NULL`（空值）的列，即使应用程序并不需要保存`NULL`也是如此, 这是因为`MULL`是列的默认值

如果查询中包含可能为`NULL`的列, 对MySQL来说更难优化, 因为可为`NULL`的列使用的索引、索引统计、和值都比较复杂。

> 可为`NULL`的列会使用更多的存储空间, 在MySQL里也需要特殊处理。当可为`NULL`的列被索引时, 每个索引记录都需要一个额外的字节, 在MyISAM里甚至还可能导致固定大小的索引变成可变大小的索引。

#### 整数类型

整型的几种类型有`TINYINT`、`SMALLINT`、`MEDIUMINT`、`INT`、`BIGINT`等。

整数类型可选`UNSIGNED`属性, 表示不允许出现负值。有符号和无符号类型使用相同的存储空间, 并且具有相同的性能, 因此可以根据实际情况选择合适的类型。

> `MySQL`可以为整数类型指定宽度, 例如`INT(11)`, `对大多数应用程序这是没有意义的`。它不会限制值的合法范围, 只是规定了MySQL一些交互工具显示字符的个数。对于存储和计算来说, `INT(1)`和`INT(11)`是相同的。

#### 实数类型

实数是带有小数部分的数字。而然它们不只是为了存储小数部分, 也可以使用`DECIMAL`存储比`BIGINT`更大的整数。
实数类型有精确小数的`DECIMAL`类型, 也有用于浮点计算的`FLOAT`, `DOUBLE`类型。

浮点类型和`DECIMAL`类型都可以指定精度, 对于`DECIMAL`列, 可以指定小数点前后所允许的最大位数, 但是会影响列的空间消耗, 因为小数点本身也占了1字节空间。浮点类型在存储相同范围的值时,通常比`DECIMAL`使用更少的空间。`FLOAT`使用4个字节存储, `DECIMAL`使用8个字节存储。

> 因为需要额外的空间和计算开销所以应该尽量只在对小数进行精确计算时才使用`DECIMAL`类型, 比如存储财务数据。
但在数据量较大时可以采用`BIGINT`进行替代。所以依然建议使用`整数类型`存储财务数据。

#### 字符串类型

字符串包含`CHAR`、`VARCHAR`、`TINYTEXT`、`TEXT`、`LONGTEXT`等

`CHAR`、`VARCHAR`是两个最最要的字符串类型, 但是这两个值的存储方式和存储引擎有很大关系, 不过最主要的`InnoDB`和`MyISAM`是类似, 下面的叙述基于该存储引擎。如果您使用的不是这两种类型存储引擎, 那么需要参考各引擎文档。

* CHAR

`CHAR`类型是定长的：MySQL总是根据定义的字符串长度分配足够的空间。

`CHAR`适合存储很短的字符串, 或者所有值都接近一个同一个长度。
例如`CHAR`非常适合存储密码的`MD5`值, 因为这是一个定长的值。对于一个经常变更的数据, `CHAR`也比`VARCHAR`更好, 因为定长的`CHAR`不容易产生碎片。

对与非常短的列`CHAR`比`VARCHAR`在存储空间上也更有效率, 例如`char(1)`用来存储单个字符, 只需要一个字符, 但是`varchar(1)`却需要两个字符, 因为它还需要一个字符用来记录长度。

* VARCHAR

`VARCHAR`用户存储可变长字符串。

大多数情况下`VARCHAR`比定长类型更节省空间, 因为它仅使用必要的空间（越短的字符串使用的空间越少）

`VARCHAR`需要使用1或2个额外字节记录字符串的长度, 如果列的最大长度小于或等于255, 则使用1个字节, 否则使用2个字节

`VARCHAR`适合以下几种情况:

1, 字符串列的最大长度比平均长度大的多

2, 列的更新很少(减少碎片)

3, 使用了`UTF-8`这样的复杂字符集, 每个字符都使用不同的字节数进行存储


#### 日期和时间类型

时间类型主要有`DATETIME`和`TIMESTAMP`, 两者都能很好的工作, 但是在某一些场景, 一个比另外一个会工作的更好。

* DATETIME

这个类型可以保存大范围的值, 从10001年到9999年, 精度为秒。它把日期和时间封装到格式为YYYYMMDDHHMMSS的整数中, 与时区无关。使用8字节的存储空间。

* TIMESTAMP

`TIMESTAMP`类型保存了从1970年1月1日午夜（UTC）以来的秒数。它和`Unix`时间戳相同。`TIMESTAMP`只使用4字节的存储空间, 因此它的存储范围比`DATETIME`小的多, 只能表示宠1970到2038年。

MySQL提供了`FROM_UNIXTIME()`函数把`Unix`时间戳转换为日期, 并提供了`UNIX_TIMESTAMP()`函数把日期转换为`Unix`时间戳


> 除特殊行为之外, 通常应该尽量使用`TIMESTAMP`, 因为它比`DATETIME`空间效率高。有时候人们会选择将`Unix`时间戳存储为整数, 但是这样并不能带来任何收益。用整数保存时间戳的格式不方便处理, 因此并不推荐这么做。

#### 位数据类型

#### 选择标识符号

## 索引优化

索引（MySql中也叫做`键(key)`）是存储引擎用于快速找到记录的一种数据结构。

理解索引是如何工作的, 一般都会类比一本书的“索引"部分: 如果想在一本书中找到某个特定主题, 一般会先看书的目录（既“索引"）, 找到对应的页码。

在`MySql`中, 存储引擎用类似的方法使用索引, 其先在索引中找到对应值, 然后根据匹配的索引记录找到对应的数据行。

索引可以包含一个或多个列的值。如果索引包含多个列，那么列的顺序十分重要, 因为`MySql`只能高效地使用索引的最左前缀列。

### 索引类型

索引有很多种类型, 可以为不同的场景提供更好的性能。在`MySql`中, 索引是在存储引擎层面而不是服务器层实现。所以, 并没有统一的索引标准: 不同的存储引擎索引的工作方式不同, 也不是所有的存储引擎都支持所有类型的索引。即使多个存储引擎支持同一类型的索引, 其底层的实现也可能不同。

#### B-Tree索引

当人们讨论索引时如果没有特别指明类型, 多半说的是B-Tree索引, 它使用[B-tree](https://iscod.github.io/#/data_struct/tree?id=_2-3%e6%a0%91%e5%92%8c2-3-4%e6%a0%91%e7%ad%89)数据结构来存储数据。大多数`MySql`引擎都支持这种索引。

存储引擎以不同的方式使用B-Tree索引, 性能也各有不同。例如：`MyISAM`使用前缀压缩技术使索引更小, 但`InnoDB`则按照原数据格式进行存储。

B-Tree索引适用于全键值、键值范围和键前缀查找。其中键前缀查找只适用于最左前缀的查询。

假设有如下数据表: 

```sql
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL DEFAULT '0',
  `last_name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓',
  `first_name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `dob` date NOT NULL,
  PRIMARY KEY (`id`),
  KEY `last_first_index` (`last_name`,`first_name`, `dob`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
  ```

- 全值查询

全值匹配值的是和索引中的所有列(`last_name`,`first_name`, `dob`)进行匹配, 例如:
```sql
select * from user where last_name = "li" and first_name = "hua" and dob = "1990-01-01"
```

- 匹配最左前缀

前面提到的索引可用于查询所有姓为 "li" 的人，既只使用索引的第一列(`last_name`), 例如:
```sql
select * from user where last_name = "li"
```

- 匹配列前缀

也可以只匹配某一列的值开头部分, 例如可以查询`last_name`列中"l“开头的姓的人:

```sql
select * from user where last_name = "l%"
```
- 匹配范围值

前面提到的索引可以查询姓名在 "li" 和 "ning" 之间的人，这里也只使用索引的第一列(`last_name`)

```sql
select * from user where last_name  between "li" and "ning"
```

- 精确匹配某一列并范围查询匹配另外一列

前面提到的索引也可用于查询姓全为 "li" 的人, 且名字是字母开头为 "h" 的人。
既第一列(`last_name`)全匹配, 第二列(`first_name`)范围匹配

```sql
select * from user where last_name = "li" and first_name like "h%"
```

- 只访问索引的查询

B-Tree通常可以支持"只访问索引的查询", 既查询只需要访问索引, 而无需访问数据行。这是覆盖索引的部分

--------

关于B-Tree索引的限制:

- 如果不是按照索引的最左列开始查找, 则无法使用索引。例如上面例子中的索引无法查找名字为"hua"的人, 也无法查找某个特定生日的人, 因为这两列都不是最左数据列。类似地，也无法查找姓氏以某个字母结尾的人。例如：

- 不能跳过索引中的列, 也就是说, 前面所诉的索引无法用于查找姓氏为"li"并且在特定日期出生的人。如果不指定名(first_name), 则MySQL只能使用索引的第一列

- 如果查询中有某列的范围查询, 则其右边所有的列都无法使用索引优化查找。
例如：

```sql
select * from user where last_name = "li" and first_name like "h%" and dob = "1990-01-01"
```
这个查询只能使用索引的前两列, 因为`LIKE`是一个范围条件。如果范围查询列值的数量有限, 可以使用多个等于条件代替范围条件。


#### 哈希索引

哈希索引(hash index)基于哈希表实现, 只有精确匹配索引所有列的查询才有效。
对于每一行数据, 存储引擎都会对所有的索引列计算一个哈希码(hash code), 哈希码是一个较小的值, 并且不同键值的行计算出来的哈希码也不一样。
哈希索引将所有的哈希码存储在索引中, 同时在哈希表中保存指向每个数据行的指针。

在`MySQL`中只有`Memory`引擎显式支持哈希索引。值得一提的是`Memory`引擎是支持非唯一哈希索引的, 这在数据库世界里是比较与众不同的。如果多个列的哈希值相同, 索引会以链表的方式存放多个记录指针到同一个哈希条目中。

哈希索引也有它的限制：

- 哈希索引只包含哈希值和行指针, 而不存在字段值, 所以不能使用索引中的值来避免读取行。不过, 访问内存中的行的速度很快, 所有大部分情况下这一点对性能的影响并不明显

- 哈希索引数据并不是按照索引值顺序存储的, 所以无法用于排序

- 哈希索引也不支持部分索引列匹配查找, 因为哈希索引始终是使用索引列的全部内容来计算哈希值。例如, 在数据列(A,B)上建立哈希索引, 如果查询只有数据列A, 则无法使用索引

- 哈希索引只支持等值比较查询, 包括 =、IN()、<=>。但不支持任何范围查询, 例如 ```where price > 100```

- 访问哈希索引的数据非常快, 除非有很多哈希冲突（不同的索引列值却有相同的哈希值）。
当出现哈希冲突的时候, 存储引擎必须遍历链表中所有的行指针, 逐行进行比较, 直到找到所有符合条件的行

- 如果哈希冲突很多的话, 一些索引维护操作的代价也会很高。例如, 如果在某个选择性很低（哈希冲突很多）的列上建立哈希索引, 那么从表中删除一行时, 存储引擎需要遍历对应哈希值的链表中的每一行, 找到并删除对应行的引用, 冲突越多, 代价越大

#### 全文索引

全文索引是一种特殊类型的索引, 它查找的是文本中关键词, 而不是直接比较索引中的值。
全文搜索和其他几个索引的匹配方式完全不一样。它有许多需要注意的细节, 比如停用词、词干和复数, 布尔搜索等。全文索引类似与搜索引擎做的事情, 而不是简单的`WHERE`条件匹配。

#### 其他索引类别

在`MySQL`还有很多第三方存储引擎使用不同类型的数据结构来存储索引。
例如: `MyISAM`支持空间数据索引(R-TREE), TokuDB使用分形树索引(fractal tree index), ScaleDB使用 Patricia tries

### 索引策略

正确的创建和使用索引是实现高性能查询的基础。

#### 独立的列

如果查询中的列不是独立的, 则`MySQL`就不会使用索引。`独立的列`是指索引列不能是表达式的一部分, 也不能是函数的参数

例如：
```sql
select * from user where id + 1 = 3;
```
凭肉眼可以看出`where`条件与`id = 2`等价, 但是`MySQL`无法自动解析方程式。这完全是用户行为。我们应该简化`where`条件的习惯, 始终将索引列单独放到比较符号的一侧

#### 前缀索引和索引选择性

有时候索引的字符列很长, 这会使索引变的大且慢, 一种策略是创建一个模拟哈希索引。但有时候这样做还不够。

通常我们可以索引列开始的部分字符, 这样可以大大节约索引空间, 从而提高索引效率, 但这样也会降低索引的选择性。

索引的选择性是指: 不重复的索引值（也称为基数）和数据表的记录总数(#T)的比值, 范围从 1/#T 到 1 之间。索引的选择性越高则查询效率越高, 因为选择性高的索引可以让`MySQL`在查找时过滤掉更多的行, 唯一索引的选择性是 1, 这是最好的索引选择性, 性能也是最好的。

对于`BLOB`, `TEXT`或者很长的`varchar`类型的列, 必须使用前缀索引, 因为MySQL不允许索引这些列的完整长度。
而诀窍在于要选择足够长的前缀以保证较高的选择性。

```sql
SELECT COUNT(*) cnt, LEFT(address, 3) FROM address WHERE 1 GROUP BY address ORDER BY cnt LIMIT 10;
```

通过测试不同的前缀长度, 找到最接近完整列的选择性

```sql
ALTER TABLE address ADD KEY (address(3));
```

前缀索引是一种能使索引更小、更快的有效方法, 但另一方面也有缺点：MySQL无法使用前缀索引做`ORDER BY`和`GROUP BY`, 也无法使用前缀索引做覆盖扫描。

> 有时候后缀索引(suffix index)也很有用途(例如找到某个域名下的所有邮件地址)。MySQL原生并不支持后缀索引, 但是可以把字符串反转后进行存储, 并基于此建立前缀索引。可以通过触发器结合`REVERSE()`来维护这种索引。

#### 多列索引

很多人对多列索引的理解不够, 一种常见的错误是, 为每个列创建独立的索引, 或者按照错误的顺序创建多列索引。

有一些专家会建议"把 WHERE 条件里面的列都建立索引", 实际上这这个建议是错误的。这样一来最好的情况下也只能是 "一星" 索引, 其性能比起真正最优的索引可能差几个数量级。例如：

```sql
select * from user where last_name = "lihua" or first_name = "hua"
```

#### 选择合适的索引顺序

- 不考虑排序和分组

在一个多列`B-Tree`索引中, 索引列的顺序意味着索引首先按照最左列进行排序, 其次是第二列, 等等。

当不需要考虑排序和分组时，将选择性最高的列放在最左侧通常是很好的，这时候索引的作用只是用于优化 WHERE 条件的查找。

对于如下查询为例:

```sql
select * from payment where staff_id = 2 and customer_id = 584;
```

是应该创建 (staff_id, customer_id) 索引还是应该颠倒一下信息？可以通过查询来确定这个表中值的分布情况，并确定那个列的选择性更高

```sql
select count(distinct staff_id)/count(*) as staff, count(distinct customer_id)/count(*) as customer from payment;
+--------+----------+
| staff  | customer |
+--------+----------+
| 0.0885 |  0.9358  |
+--------+----------+
1 row in set (0.00 sec)

```

`customer_id`的选择性更高, 所以答案是将其最为索引列的第一列：
```sql
ALTER TABLE `payment` ADD INDEX customer_staff(customer_id, staff_id);
```

> 经验法则: 将选择性最高的列放在前面通常是很好的


- 如果有分组和范围查询时？

试想一下如果对于下列查询：

```sql
select * from payment where staff_id = 2 and customer_id > 584 and customer_id < 600;
```

对比创建`(customer_id, staff_id)`和`(staff_id,customer_id)`两种索引顺序时的查询信息：

```sql
explain select * from payment force index(customer_staff) where staff_id = 2 and customer_id > 584 and customer_id < 600\G;
*************************** 1. row ***************************
           id: 1
  select_type: SIMPLE
        table: payment
   partitions: NULL
         type: range
possible_keys: customer_staff
          key: customer_staff
      key_len: 6
          ref: NULL
         rows: 54
     filtered: 2.50
        Extra: Using index condition
1 row in set, 1 warning (0.00 sec)
```

```sql
explain select * from payment force index(staff_customer) where staff_id = 2 and customer_id > 584 and customer_id < 600\G;
*************************** 1. row ***************************
           id: 1
  select_type: SIMPLE
        table: payment
   partitions: NULL
         type: range
possible_keys: staff_customer
          key: staff_customer
      key_len: 11
          ref: NULL
         rows: 2
     filtered: 100.00
        Extra: Using index condition
1 row in set, 1 warning (0.01 sec)
```

虽然`customer_id`列拥有更高的选择性，但是使用`customer_staff`索引时，扫描的行数(rows)很多。
这是由于索引扫描围造成了两种偏差。详细了解[扫描索引的范围](https://use-the-index-luke.com/sql/where-clause/searching-for-ranges/greater-less-between-tuning-sql-access-filter-predicates)。

> 经验法则: 首先索引相等，然后索引范围。

#### 覆盖索引

如果索引的叶子节点已经包含查询的数据，那么还设有什么必要进行回表查询？如果一个索引包含（或者说覆盖）所有需要查询的字段值，我们就称之为`覆盖索引`



## 查询优化

查询性能低下最基本的原因是访问的数据太多, 某些查询可能不可避免地需要筛选大量数据, 但这并不常见。

大部分性能低下的查询都可以通过减少访问数据量的方式进行优化。

对于低效的性能查询, 我们可以通过下面两个步骤来分析:

- 确认应用程序是否检索大量超过需要的数据, 这通常意味着访问了太多的行, 但有时候也可能是访问了太多的列
- 确认MySQL服务器层是否在分析大量超过需要的数据行

这里有一些经典案例:

- 查询了不需要的记录
- 多表关联时返回了全部列
- 总是取出全部的列
- 重复查询相同的数据

`explain`命令是我们查询优化的利器, `explain`可以告诉我们查询语句的很多相关信息, 比如扫描的行数, 使用的索引等

![profiles](https://iscod.github.io/images/mysql_explain_1.png)

`explain` 返回的相关列：

列名|含义|描述
----|-------|---
id| 选择标识符| |
select_type| | |
table|表名| |
partitions|匹配的分区| |
type|访问类型|All:全表扫描,index:索引扫描, range:范围扫描, ref:非唯一性索引或者唯一索引前缀扫描, eq_ref:唯一性索引const:常数引用等|
possible_keys|查询可能用到的索引| |
key|查询实际使用的索引| |
key_len|索引字段的长度| |
ref|列与索引的比较| |
rows|预计扫描的行数| |
filtered|按表条件过滤的行百分比| |
Extra|执行情况描述说明| |

## 常用命令

- help

```sql
HELP CONTENTS;
```

- 查看当前运行的进程

```sql
show processlist;
show full processlist;
 ```

 - 杀死进程

```sql
kill 9
```

- exlpin

```sql
explain select * from user where id in (10);
```

- 整理文件碎片

```sql
optimize table table_name;
```

* 参考
    * [use-the-index-luke](https://use-the-index-luke.com/)