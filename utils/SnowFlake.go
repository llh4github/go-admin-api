package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// IDWorker 雪花算法工具
var IDWorker *SnowflakeIdWorker

// NextIDDefalut 默认配置下id生成
func NextIDDefalut() string {
	if IDWorker == nil {
		IDWorker, _ = CreateWorker(1, 2)
		return IDWorker.NextIDStr()
	}
	return IDWorker.NextIDStr()
}

const (
	//t := time.Date(2015, 1, 1, 00, 00, 00, 00, time.Local).UnixNano() / 1e6;//获取时间戳 毫秒
	//开始时间戳 2015-1-1
	epoch int64 = 1420041600000
	// 机器id所占的位数
	workerIdBits int64 = 5
	// 数据标识id所占的位数
	datacenterIdBits int64 = 5
	//支持的最大机器id，结果是31 (这个移位算法可以很快的计算出几位二进制数所能表示的最大十进制数)
	maxWorkerId int64 = -1 ^ (-1 << workerIdBits)
	// 支持的最大数据标识id，结果是31
	maxDatacenterId int64 = -1 ^ (-1 << datacenterIdBits)
	//序列在id中占的位数
	sequenceBits int64 = 12
	// 机器ID向左移12位
	workerIdShift int64 = sequenceBits
	// 数据标识id向左移17位(12+5)
	datacenterIdShift int64 = sequenceBits + workerIdBits
	// 时间截向左移22位(5+5+12)
	timestampLeftShift int64 = sequenceBits + workerIdBits + datacenterIdBits
	// 生成序列的掩码，这里为4095 (0b111111111111=0xfff=4095)
	sequenceMask int64 = -1 ^ (-1 << sequenceBits)
)

//  SnowflakeIdWorker 雪花算法结构体
type SnowflakeIdWorker struct {
	mutex         sync.Mutex // 添加互斥锁 确保并发安全
	lastTimestamp int64      // 上次生成ID的时间截
	workerId      int64      // 工作机器ID(0~31)
	datacenterId  int64      //数据中心ID(0~31)
	sequence      int64      // 毫秒内序列(0~4095)
}

// CreateWorker 创建SnowflakeIdWorker
// 	workerId 工作ID (0~31)
// 	datacenterId 数据中心ID (0~31)
func CreateWorker(wID int64, dID int64) (*SnowflakeIdWorker, error) {
	if wID < 0 || wID > maxWorkerId {
		return nil, errors.New("Worker ID excess of quantity")
	}
	if dID < 0 || dID > maxDatacenterId {
		return nil, errors.New("Datacenter ID excess of quantity")
	}
	// 生成一个新节点
	return &SnowflakeIdWorker{
		lastTimestamp: 0,
		workerId:      wID,
		datacenterId:  dID,
		sequence:      0,
	}, nil
}

// NextIDStr 生成字符串的id
func (w *SnowflakeIdWorker) NextIDStr() string {
	return strconv.FormatInt(w.nextId(), 10)
}

// nextId 生成id
func (w *SnowflakeIdWorker) nextId() int64 {
	// 保障线程安全 加锁
	w.mutex.Lock()
	// 生成完成后 解锁
	defer w.mutex.Unlock()
	// 获取生成时的时间戳 毫秒
	now := time.Now().UnixNano() / 1e6
	//如果当前时间小于上一次ID生成的时间戳，说明系统时钟回退过这个时候应当抛出异常
	if now < w.lastTimestamp {
		errors.New("Clock moved backwards")
		//根据需要自定义错误码
		return 3001
	}
	if w.lastTimestamp == now {
		w.sequence = (w.sequence + 1) & sequenceMask
		if w.sequence == 0 {
			// 阻塞到下一个毫秒，直到获得新的时间戳
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.sequence = 0
	}
	// 将机器上一次生成ID的时间更新为当前时间
	w.lastTimestamp = now
	ID := int64((now-epoch)<<timestampLeftShift | w.datacenterId<<datacenterIdShift | (w.workerId << workerIdShift) | w.sequence)
	return ID
}

/*
 * 将十进制数字转化为二进制字符串
 */
func convertToBin(num int64) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// 将数字强制性转化为字符串
		s = strconv.FormatInt(lsb, 10) + s
	}
	return s
}
