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
