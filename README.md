# 假数据生成

魔法数字生成，各种假数据生成实现

## 例子
``` golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ideadsnow/fakegen"
)

func main() {
	// 随时间流逝随机步长单调递增的计数生成器，用于一些需要生成假数据的场景，如：总预约人数、总下载量等
	counter := fakegen.NewMICounter(10, fakegen.WithRate(5), fakegen.WithMaxStep(100))

	for i := 0; i < 10; i++ {
		t := rand.Int63n(5 * 1000 * 1000 * 1000) // [0-5)秒随机
		time.Sleep(time.Duration(t))

		n := counter.Next(time.Now())

		fmt.Println(i, n)
	}
}
```

Output
```shell
0 14
1 15
2 15
3 15
4 33
5 37
6 37
7 38
8 39
9 41
```

## MICounter

monotone increasing by timestamp counter  
随时间（秒级）和调用次数增加，随机且单调递增的计数生成器  
时间间隔越久，随机递增的可能步长越大

```
step = ((nowTS - prevTS)*rate) % maxStep

if step == 0:
  step = maxStep

f(nowTS) = prevCount + step
```

Tips：

1. 建议 New 的时候读存储传入 count 初始值，且获得新 count 值后写入存储，避免进程重启导致计数器回退
2. rate 用于扩大步长范围，结合业务场景合理使用可以令数据显得更加真实
3. maxStep 用于限制最大步长，避免计数器突变

