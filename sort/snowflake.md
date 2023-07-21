# 雪花算法

雪花算法是有[twitter-snowflake](https://github.com/twitter-archive/snowflake)提出的分布式ID解决方案

雪花算法默认结构：

```
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+
```

1. 高并发
1. 高可用, 每台机器每毫秒最高提供4096个ID，可使用
1. 高扩展，将分布式服务包装为web或者RPC服务, 结合k8s进行扩展

### 原始雪花算法问题

1. 序列超过容量(4096):

    * 采用程序阻塞1ms, 系统提供下一毫秒的算法ID

1. 时钟回拨问题：

    1. 时钟回拨时间较短，例如100ms内，可以采用线程阻塞等待
    1. 时钟回拨时间时钟，例如 100ms - 1s 秒内，可以采用内存维护每毫秒内 max 进行缓存
    1. 时钟回拨时间较长，例如 1s-5s内, 可以不提供服务，在客户端进行兼容调用其它机器进行获取
    1. 时钟回拨时间很长，例如 > 5s，可以采用容器探针进行pod的重启更新

```go
// 构建雪花算法package
// https://github.com/iscod/iscod.github.io/tree/master/example/snowflake
package snowflake

import (
    "github.com/pkg/errors"
    "time"
)

var ErrClockBack error = errors.New("Clock moving backwards")

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

type SnowFlake interface {
    NextId() (ID, error)
}

// New returns an SnowFlake interface
func NewSnowflake(workerId int64) SnowFlake {
    s := &snowflake{
        //epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
        epoch:     time.UnixMilli(1288834974657),
        workerId:  workerId,
        seqBit:    12,
        workIdBit: 10,
        //sequenceMask: 4095,
    }
    s.sequenceMask = -1 ^ (-1 << 12) // 4095

    return s
}

type snowflake struct {
    workIdBit     int
    seqBit        int
    epoch         time.Time
    lastTimestamp int64
    sequence      int64
    workerId      int64
    sequenceMask  int64
}

func (s *snowflake) NextId() (ID, error) {
    if s == nil {
        return 0, errors.New("snowflake is nil")
    }

    timestamp := time.Since(s.epoch).Milliseconds()

    if timestamp > s.lastTimestamp {
        s.sequence = 0
    } else if timestamp < s.lastTimestamp { //时钟回拨
        if offset := s.lastTimestamp - timestamp; offset <= 5 {
            time.Sleep(time.Millisecond * time.Duration(offset))
            timestamp = time.Since(s.epoch).Milliseconds()
            if timestamp < s.lastTimestamp {
                return 0, ErrClockBack
            }
        } else {
            return 0, ErrClockBack
        }
    } else if timestamp == s.lastTimestamp {
        if s.sequence = (s.sequence + 1) & s.sequenceMask; s.sequence == 0 {
            //seq 为0的时候表示该毫秒已达到4096个，等待下一毫秒时间开始对seq做随机
            for timestamp == s.lastTimestamp {
                timestamp = time.Since(s.epoch).Milliseconds()
            }
        }
    }

    s.lastTimestamp = timestamp
    id := ID((timestamp)<<(s.workIdBit+s.seqBit) | (s.workerId << s.seqBit) | s.sequence)
    return id, nil
}
```

* 参考
    * [twitter-snowflake](https://github.com/twitter-archive/snowflake)
    * [Meituan-Leaf](https://github.com/Meituan-Dianping/Leaf)
    * [bwmarrin-snowflake](https://github.com/bwmarrin/snowflake)


