// Copyright © 2021 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/fanux/sealos/cmd"     //包被导入已执行初始化函数
)

func main() {
	cmd.Execute()   //执行cmd模块root.go里的这个方法。安装时用的sealos init,去看cmd/init.go
}
             /*这里使用了corba这个命令行模块，其命令与文件对应，例如sealos init对应init.go,sealos add对应add.go */
             /*sealos是根命令，对应root.go，先看看root.go*/
