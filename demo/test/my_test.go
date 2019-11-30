package my_test

import (
	"encoding/json"
	"fmt"
	docxt "go-docx-templates"
	"testing"
	"time"
)

type Plan struct {
	Start   string
	End     string
	Company string
	Contact string
	Mobile  string
	Total   int64
	Place   string
	Goal    string
	Stu     string
}

func TestWord(t *testing.T) {
	str := `{"examin":[{"Exdate":1561061258391,"Val":23,"Typ":"纪律规范"},{"Exdate":1561061258391,"Val":18,"Typ":"学习情况"},{"Exdate":1561061258391,"Val":6,"Typ":"综合素养"},{"Exdate":1561061258391,"Val":8,"Typ":"奖励加分"}],"flunk":[{"Term":1,"No":"330722199812092611","Grade":0,"Name":"必备外的技能证书学分","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":0,"Name":"各类比赛获奖学分","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":5,"Name":"音乐","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":8,"Name":"特色社会主义","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":8,"Name":"企业管理","Passgrade":60}],"interview":[{"Company":"永康威力科技股份有限公司","Post":"数控铣岗","Treatment":"额温柔温柔","Isemployed":1}],"pass":[],"plan":[{"Post":"数控铣岗","Start":1561173764000,"End":1561173766000,"Company":"永康威力科技股份有限公司"},{"Post":"钳工","Start":1559664000000,"End":1560355200000,"Company":"浙江铁牛汽车车身有限公司"},{"Post":"数控铣岗","Start":1559944089000,"End":1560441600000,"Company":"浙江铁牛汽车车身有限公司"}],"punish":[{"RecordAt":1561478400000,"Detail":"了模糊葡萄XP你也是","Rule":"讲脏话"}],"reward":[],"Stu":{"enroll_year":2014,"gender":"男","mobile":"13967936015","Name":"徐 最","no":"330722199812092611","org":"数控141","org_id":"aade804e5d8811e9b1ed8c1645d34af8/37fca1335ba611e9b8558c1645d34af8"}}`
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		fmt.Println(err)
	}

	fmt.Println(m)

	template, err := docxt.OpenTemplate("plan1.docx")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := template.RenderTemplate(m); err != nil {
		fmt.Println(err)
		return
	}
	if err := template.Save("result.docx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func TestTime(t *testing.T) {
	timeLayout := "2006-01-02"
	dataTimeStr := time.Unix(1559899190000/1000, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}
