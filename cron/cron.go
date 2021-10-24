package cronTask

import (
	"github.com/robfig/cron/v3"
	"log"
)

func CronTaskInit() {

	// 在线生成cron表达式 https://cron.qqe2.com/
	c := cron.New()
	c.AddFunc("@every 2s", func() {
		log.Println("从0秒开始，每2秒执行一次")
	})
	c.AddFunc("* 0/1 * * *", func() {
		log.Println("从0分开始，每1分钟执行一次")
	})

	c.Start()

}
