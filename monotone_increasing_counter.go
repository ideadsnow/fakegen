package fakegen

import (
	"math/rand"
	"time"
)

/*
	monotone increasing by t counter
	随时间（秒级）和调用次数增加，随机且单调递增的计数生成器
	时间间隔越久，随机递增的可能步长越大

	f(nowTS) = prevCount + max( ((nowTS - prevTS)*rate)%maxStep, maxStep )
	Tips：
		1. 建议 New 的时候读存储传入 count 初始值，且获得新 count 值后写入存储，避免进程重启导致计数器回退
		2. rate 用于扩大步长范围，结合业务场景合理使用可以令数据显得更加真实
		3. maxStep 用于限制最大步长，避免计数器突变
*/

type MICounter struct {
	options

	count     int64 // 预载
	timestamp int64 // 上一次的时间戳
	r         *rand.Rand
}

// Next 入参：UNIX t
func (mi *MICounter) Next(t time.Time) int64 {
	timestamp := t.Unix() // 计数器变化是秒级
	delta := timestamp - mi.timestamp
	mi.timestamp = timestamp
	if delta <= 0 {
		return mi.count
	}
	mi.incr(delta)
	return mi.count
}

// 理想情况下增长量符合泊松分布，这里求简单直接线性实现，一般业务要求也没那么高
func (mi *MICounter) incr(delta int64) {
	step := delta * int64(mi.rate)
	if mi.maxStep > 0 {
		step = step % int64(mi.maxStep)
	}
	if step <= 0 {
		step = int64(mi.maxStep)
	}
	mi.count += mi.r.Int63n(step)
}

func NewMICounter(count int64, opts ...Option) *MICounter {
	options := options{
		rate:    DefaultRate,
		maxStep: DefaultMaxStep,
		t:       time.Now(),
	}

	for _, o := range opts {
		o.apply(&options)
	}

	timestamp := options.t.Unix() // 计数器变化是秒级
	counter := &MICounter{
		options: options,

		count:     count,
		timestamp: timestamp,
		r:         rand.New(rand.NewSource(timestamp)),
	}

	return counter
}
