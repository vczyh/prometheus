package extend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//type RuleItem struct {
//	Name        string            //规则所属组得名称
//	Fn          string            //类别
//	Interval    int               //规则计算间隔
//	Alert       string            //告警名称
//	Expr        string            //规则表达式
//	For         string            //持续时间
//	Labels      map[string]string //规则维度信息
//	Annotations map[string]string //规则描述信息
//}

type RuleInfo struct {
	MonitorType string      `json:"monitor_type"`
	RuleGroups  []RuleGroup `json:"rule_groups"`
}

type RuleGroup struct {
	GroupName string         `json:"group_name"`
	Interval  time.Duration `json:"interval,omitempty"`
	Rules     []Rule         `json:"rules"`
}

type Rule struct {
	Record      string            `json:"record,omitempty"`
	Alert       string            `json:"alert,omitempty"`
	Expr        string            `json:"expr"`
	For         time.Duration    `json:"for,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func LoadBody() ([]RuleInfo, error) {
	resp, err := http.Get("http://10.0.7.141:8080/v1/rules")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ruleInfos []RuleInfo
	err = json.Unmarshal(body, &ruleInfos)
	if err != nil {
		return nil, err
	}

	return ruleInfos, nil
}
