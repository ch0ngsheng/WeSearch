package mq

import (
	"log"
	"time"
	
	"github.com/zeromicro/go-zero/core/logx"
)

// RetryWith 重试， max=-1 means forever
func RetryWith(name string, interval time.Duration, max int, do func() error) {
	var err error
	ttl := max
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		if ttl == 0 {
			log.Fatalf("%s, retry exceed max %d times, err: %v", name, max, err)
		}
		ttl -= 1

		if err = do(); err == nil {
			return
		}

		ticker.Reset(interval)
		logx.Errorf("%s failed, ttr: %d, err: %+v", name, ttl, err)
		select {
		case <-ticker.C:
		}
	}
}
