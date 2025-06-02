package cmd

import (
	"base_frame/internal"
	"base_frame/pkg/cmd"
	"base_frame/pkg/common/config"
	"base_frame/pkg/program"
	"context"
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

const DefaultConfigPath = "dev/config.yaml"

type ApiCmd struct {
	*cmd.RootCmd
	ctx    context.Context // 设置基本的上下文
	config config.Config   // 配置项
}

func NewApiCmd() *ApiCmd {
	var ret ApiCmd
	ret.RootCmd = cmd.NewRootCmd(program.GetProcessName())
	ret.ctx = context.WithValue(context.Background(), "version", "test_version")
	ret.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.runE()
	}
	return &ret
}

func (a *ApiCmd) runE() error {
	// 定义 -c 参数，接收配置文件路径
	configPath := flag.String("c", "", DefaultConfigPath)
	flag.Parse()

	if *configPath == "" {
		log.Fatal("请使用 -c 参数指定配置文件路径，例如：-c dev/config")
	}

	// 使用 Viper 加载配置文件
	viper.SetConfigFile(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	// 获取配置文件内容
	if err := viper.Unmarshal(&a.config); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	return internal.Start(a.ctx, &a.config)
}

func (a *ApiCmd) Exec() error {
	return a.Command.Execute()
}
