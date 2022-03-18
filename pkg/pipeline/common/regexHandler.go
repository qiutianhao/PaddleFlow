/*
Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserve.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"fmt"
	"regexp"
)

type VariableChecker struct {
}


func (variableChecker *VariableChecker) CheckVarName(varName string) error {
	// 校验字符串是一个合格变量名，只能由字母数字下划线组成，且以字母下划线开头
	pattern := `^[a-zA-Z_][a-zA-Z_0-9]*$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(varName) {
		err := fmt.Errorf("format of variable name[%s] invalid, should be in ^[a-zA-Z_$][a-zA-Z_$0-9]*$", varName)
		return err
	}
	return nil
}

func (variableChecker *VariableChecker) CheckRefUpstreamStep(varValue string) error {
	// 匹配引用上游节点参数的字符串
	pattern := `^\{\{(\s)*[a-zA-Z0-9_]+\.[a-zA-Z0-9_]+(\s)*\}\}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(varValue) {
		err := fmt.Errorf("format of value[%s] invalid, should be like {{XXX.XXX}}", varValue)
		return err
	}
	return nil
}