package validate

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type BaseValidate struct {
}

func init()  {
	msg := map[string]string{
		"Required":     "不能为空",
		"Min":          "最小为 %d 位",
		"Max":          "最大为 %d 位",
		"Range":        "范围在 %d 至 %d 位",
		"MinSize":      "最小长度为 %d 位",
		"MaxSize":      "最大长度为 %d 位",
		"Length":       "长度必须是 %d 位",
		"Alpha":        "必须是有效的字母字符",
		"Numeric":      "必须是有效的数字字符",
		"AlphaNumeric": "必须是有效的字母或数字字符",
		"Match":        "必须匹配格式 %s",
		"NoMatch":      "必须不匹配格式 %s",
		"AlphaDash":    "必须是有效的字母或数字或破折号(-_)字符",
		"Email":        "必须是有效的邮件地址",
		"IP":           "必须是有效的IP地址",
		"Base64":       "必须是有效的base64字符",
		"Mobile":       "必须是有效手机号码",
		"Tel":          "必须是有效电话号码",
		"Phone":        "必须是有效的电话号码或者手机号码",
		"ZipCode":      "必须是有效的邮政编码",
	}
	validation.SetDefaultMessage(msg)
}

func (*BaseValidate) Validate(param interface{},types interface{}) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(param)
	if !b {
		st := reflect.TypeOf(types)
		for _, err2 := range valid.Errors {
			fmt.Println(err2.Name)
			fmt.Println(err2.Value)
			if !(err2.Name != "Required" && err2.Value == "") {
				field,_ := st.FieldByName(err2.Field)
				var alias = field.Tag.Get("alias")
				msg := strings.Replace(err2.Message," ","",-1)
				if alias!="" {
					msg = strings.Replace(msg,err2.Field,alias,1)
				}
				return errors.New(msg)
			}
		}
	}
	return nil
}
