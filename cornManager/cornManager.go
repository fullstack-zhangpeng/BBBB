package cornManager

import (
	"Service/service-server/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	corn "github.com/robfig/cron"
)

var c *corn.Cron

func InitCorn() {
	c = corn.New()
	c.AddFunc("0 0 9 * * *", func() {
		sendDutyInfo()
	})
	c.Start()
}

func StartPushDuty() {
	c.Start()
}

func StopPushDuty() {
	c.Stop()
}

func getTodayDutyInfo() []interface{} {
	var duty map[string]interface{}
	file := "./resource/duty.json"
	err := json.Unmarshal(util.GetDataFromeFile(file), &duty)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	today := time.Now().Format("20060102")

	dutyList, ok := duty[today].([]interface{})

	if ok {
		return dutyList
	}
	return nil
}
func sendDutyInfo() {
	dutyList := getTodayDutyInfo()

	content := ""

	prefix := "今日值班人员如下：\n"

	content += prefix

	for i, val := range dutyList {
		dutyInfo, ok := val.(map[string]interface{})
		if ok {
			content += fmt.Sprintf("%s, %s", dutyInfo["position"], dutyInfo["name"])
			if i != (len(dutyList) - 1) {
				content += "\n"
			}
		}
	}
	content += "\n"

	suffix := "请大家注意关注群内问题！"
	content += suffix

	fmt.Println(content)
	// content := "我就是我, 是不一样的烟火@ Test"
	formt := `
	{
		"msgtype": "text",
		"text": {
			"content": "%s"
		}
	}`
	body := fmt.Sprintf(formt, content)
	jsonValue := []byte(body)
	dingdingURL := "https://oapi.dingtalk.com/robot/send?access_token=16730cf5887fdd912115d3a729c2fddbd074bbe9b7bbc182740d47250696bcdb"
	resp, _ := http.Post(dingdingURL, "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	// 	return err
	// }
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(result))
	fmt.Print("-------------")
}

// func Start() {
// 	c := cron.New()
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllTag...")
// 	})
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllArticle...")
// 	})

// 	c.Start()

// 	t1 := time.NewTimer(time.Second * 10)
// 	for {
// 		select {
// 		case <-t1.C:
// 			t1.Reset(time.Second * 10)
// 		}
// 	}
// }
