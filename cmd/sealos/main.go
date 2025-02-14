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
	"github.com/containers/buildah"
	"github.com/labring/sealos/cmd/sealos/cmd"  //包被导入已执行初始化函数
)

func main() {
	if buildah.InitReexec() {   //这里为什么要层层引用呢？buildah.InitReexec又引用reexec的如果未指定任何参数，程序立即退出
		return
	}
	cmd.Execute()  //执行cmd模块root.go里的这个方法。安装时用的sealos run,去看cmd/run.go
}
