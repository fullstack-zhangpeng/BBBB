package timing

import (
	"Service/service-server/db/model"
	"Service/service-server/dingtalk"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

var c *cron.Cron

func Init() {
	c = cron.New()
	c.AddFunc("0 10 9 * * *", func() {
		sendDutyInfo()
	})
	//c.AddFunc("* * * * * *", func() {
	//	log.Println("test timing service...")
	//	sendDutyInfo()
	//})
}

func Start() {
	c.Start()
}

func Stop() {
	c.Stop()
}

func getTodayDutyInfo() []model.Duty {
	today := time.Now().Format("20060102")
	p := &model.Duty{
		Date: today,
	}
	dutyList := p.NewRetrieve()

	return dutyList
}
func sendDutyInfo() {
	var dutyList []model.Duty
	dutyList = getTodayDutyInfo()
	if len(dutyList) <= 0 {
		return
	}

	content := ""
	prefix := "今日值班人员如下：\n"
	content += prefix
	for i, duty := range dutyList {
		content += fmt.Sprintf("%s, %s", duty.Position, duty.Name)
		if i != (len(dutyList) - 1) {
			content += "\n"
		}
	}
	content += "\n"

	suffix := "请大家注意关注群内问题！"
	content += suffix

	//移动端
	//dt := dingtalk.NewDingTalk("2062dd60bdc9bae2412e289c83d3e3613ecf4205f30df9c93e294cc5e00a5791")
	//大前端
	dt := dingtalk.NewDingTalk("16730cf5887fdd912115d3a729c2fddbd074bbe9b7bbc182740d47250696bcdb")
	message := dingtalk.NewTextMessage(content)
	res, err := dt.Send(message)
	fmt.Println(res)
	fmt.Println(err)
}
