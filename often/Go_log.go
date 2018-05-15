package main

import (
	log1 "github.com/sirupsen/logrus"
	log2 "github.com/cihub/seelog"
	"os"
	"fmt"
)

func TestLoggrus1()  {
	log1.WithFields(log1.Fields{"姓名":"懿"}).Info("中奖啦！！！")
}

func TestLoggrus2()  {
	//日志格式化为JSON而不是默认的ASCII
	log1.SetFormatter(&log1.JSONFormatter{})

	//输出stdout而不是默认的stderr，也可以是一个文件
	log1.SetOutput(os.Stdout)

	//只记录严重或者以上的警告
	log1.SetLevel(log1.WarnLevel)

	//此行不会打印
	log1.WithFields(log1.Fields{"姓名":"懿"}).Info("我居然没中奖！！！")

	log1.WithFields(log1.Fields{"姓名": "YI"}).Error("中奖是不可能中奖的，这辈子都是不可能的！")
}

func TestSeelog1()  {
	defer log2.Flush()
	log2.Info("Hello from seelog!")
}

func TestSeelog2()  {
	//xml 格式的字符串，xml 配置了如何输出日志
	testConfig :=`
	<seelog type="sync">
        // 配置输出项，本例表示同时输出到控制台和日志文件中
    	<outputs formatid="main">
    		<console/>
    		<file path="log.log"/>
    	</outputs>
    	<formats>
    	    // 日志格式，outputs中的输出项可以通过id引用该格式
    		<format id="main" format="[%LEVEL] [%Time] [%FuncShort @ %File.%Line] %Msg%n"/>
    	</formats>
    </seelog>`

    //根据配置信息生成logger(配置信息的对象表示)
    logger, err := log2.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}

	//配置日志组件加载配置信息，然后输出函数比如info在输出时，会加载改配置
	log2.ReplaceLogger(logger)
    log2.Info("Hello from Seelog")
}

func main() {
	//logrus 是用Go语言实现的一个日至系统，与标准库log完全兼容并且核心API很稳定，是Go语言目前最活跃的日志库
	TestLoggrus1()
	TestLoggrus2()

	//seelog 使用 Go 语言实现的一个日志系统，它提供了一些简单的函数来实现复杂的日志分配、过滤和格式化
	TestSeelog1()
	TestSeelog2()
}
