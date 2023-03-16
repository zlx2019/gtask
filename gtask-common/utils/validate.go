/**
  @author: Zero
  @date: 2023/3/16 14:47:47
  @desc: 参数校验工具

**/

package utils

import (
	"regexp"
	"zero.com/gtask-common/response/errors"
)

// VerifyMobile 校验手机号码是否符合规范
func VerifyMobile(mobile string) error {
	if mobile == "" {
		return errors.New("手机号码不能为空")
	}
	// 通过正则匹配
	must := regexp.MustCompile("^1[3|4|5|8|7|9][0-9]\\d{8}$")
	if !must.MatchString(mobile) {
		return errors.MobileNotLegalErr
	}
	return nil
}
