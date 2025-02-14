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

package cmd

import (
	"fmt"
	"os"

	"github.com/labring/sealos/pkg/utils/contants"
	"github.com/labring/sealos/pkg/utils/file"
	"github.com/labring/sealos/pkg/utils/logger"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	debug bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sealos",
	Short: "simplest way install kubernetes tools.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(onBootOnDie)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logger")   //这个持久标志,必须设置的
}

func onBootOnDie() {
	var rootDirs = []string{
		contants.LogPath(),        //获取日志、数据路径及工作目录，定义见contants包。Logpath()及DataPaht()见contans/data.go
		contants.DataPath(),
		contants.Workdir(),      //Wrokdir()见contants/worker.go，为家目录的.sealos下
	}
	if err := file.MkDirs(rootDirs...); err != nil {  //创建上述目录，见file/file_v2.go，为什么不用系统自带的呢？
		logger.Error(err)   //logger这里引用了另一个很好地处理日志的模块，报错并退出
		panic(1)
	}
	logger.CfgAndFile(debug, contants.LogPath(), "sealos", false)    //使用这个模块，设置日志，将日志输出到日志路径
	setupLogrus()
}

const logLevel = "warn"

func setupLogrus() {
	logrusLvl, _ := logrus.ParseLevel(logLevel)   //设置日志级别
	logrus.SetLevel(logrusLvl)
}
