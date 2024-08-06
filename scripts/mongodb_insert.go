package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 数据库连接信息
const (
	// mongoURI     = "mongodb://root:abcd1234@10.40.138.88:27017,10.40.25.174:27017" // MongoDB连接URI
	// mongoURI     = "mongodb://root:abcd1234@10.40.138.88:27017,10.40.25.174:27017/?replicaSet=udb-qcyfsvo2zyp" // MongoDB连接URI
	mongoURI     = "mongodb://root:k22BVbYEriueRA2Z@10.9.65.132:27017,10.9.29.176:27017,10.9.177.129:27017/?replicaSet=umongodb-rs-10s65j61gnh7" // MongoDB连接URI
	database     = "walter"                                                                                                                      // 数据库名称
	collection   = "chen"                                                                                                                        // 集合名称
	numDocuments = 10000                                                                                                                         // 要插入的文档数量
)

func main() {
	text := `阮一峰的网络日志 » 首页 » 档案
搜索
上一篇：科技爱好者周刊：第 7
下一篇：科技爱好者周刊：第 7
分类： 开发者手册
Tmux 使用教程
作者： 阮一峰

日期： 2019年10月21日

Tmux 是一个终端复用器（terminal multiplexer），非常有用，属于常用的开发工具。

本文介绍如何使用 Tmux。



一、Tmux 是什么？
1.1 会话与进程
命令行的典型使用方式是，打开一个终端窗口（terminal window，以下简称"窗口"），在里面输入命令。用户与计算机的这种临时的交互，称为一次"会话"（session） 。

会话的一个重要特点是，窗口与其中启动的进程是连在一起的。打开窗口，会话开始；关闭窗口，会话结束，会话内部的进程也会随之终止，不管有没有运行完。

一个典型的例子就是，SSH 登录远程计算机，打开一个远程窗口执行命令。这时，网络突然断线，再次登录的时候，是找不回上一次执行的命令的。因为上一次 SSH 会话已经终止了，里面的进程也随之消失了。

为了解决这个问题，会话与窗口可以"解绑"：窗口关闭时，会话并不终止，而是继续运行，等到以后需要的时候，再让会话"绑定"其他窗口。

1.2 Tmux 的作用
Tmux 就是会话与窗口的"解绑"工具，将它们彻底分离。

（1）它允许在单个窗口中，同时访问多个会话。这对于同时运行多个命令行程序很有用。

（2） 它可以让新窗口"接入"已经存在的会话。

（3）它允许每个会话有多个连接窗口，因此可以多人实时共享会话。

（4）它还支持窗口任意的垂直和水平拆分。

类似的终端复用器还有 GNU Screen。Tmux 与它功能相似，但是更易用，也更强大。

二、基本用法
2.1 安装
Tmux 一般需要自己安装。


# Ubuntu 或 Debian
$ sudo apt-get install tmux

# CentOS 或 Fedora
$ sudo yum install tmux

# Mac
$ brew install tmux
2.2 启动与退出
安装完成后，键入tmux命令，就进入了 Tmux 窗口。


$ tmux
上面命令会启动 Tmux 窗口，底部有一个状态栏。状态栏的左侧是窗口信息（编号和名称），右侧是系统信息。



按下Ctrl+d或者显式输入exit命令，就可以退出 Tmux 窗口。


$ exit
2.3 前缀键
Tmux 窗口有大量的快捷键。所有快捷键都要通过前缀键唤起。默认的前缀键是Ctrl+b，即先按下Ctrl+b，快捷键才会生效。

举例来说，帮助命令的快捷键是Ctrl+b ?。它的用法是，在 Tmux 窗口中，先按下Ctrl+b，再按下?，就会显示帮助信息。

然后，按下 ESC 键或q键，就可以退出帮助。

三、会话管理
3.1 新建会话
第一个启动的 Tmux 窗口，编号是0，第二个窗口的编号是1，以此类推。这些窗口对应的会话，就是 0 号会话、1 号会话。

使用编号区分会话，不太直观，更好的方法是为会话起名。


$ tmux new -s <session-name>
上面命令新建一个指定名称的会话。

3.2 分离会话
在 Tmux 窗口中，按下Ctrl+b d或者输入tmux detach命令，就会将当前会话与窗口分离。


$ tmux detach
上面命令执行后，就会退出当前 Tmux 窗口，但是会话和里面的进程仍然在后台运行。

tmux ls命令可以查看当前所有的 Tmux 会话。


$ tmux ls
# or
$ tmux list-session
3.3 接入会话
tmux attach命令用于重新接入某个已存在的会话。


# 使用会话编号
$ tmux attach -t 0

# 使用会话名称
$ tmux attach -t <session-name>
3.4 杀死会话
tmux kill-session命令用于杀死某个会话。


# 使用会话编号
$ tmux kill-session -t 0

# 使用会话名称
$ tmux kill-session -t <session-name>
3.5 切换会话
tmux switch命令用于切换会话。


# 使用会话编号
$ tmux switch -t 0

# 使用会话名称
$ tmux switch -t <session-name>
3.6 重命名会话
tmux rename-session命令用于重命名会话。


$ tmux rename-session -t 0 <new-name>
上面命令将0号会话重命名。

3.7 会话快捷键
下面是一些会话相关的快捷键。

Ctrl+b d：分离当前会话。
Ctrl+b s：列出所有会话。
Ctrl+b $：重命名当前会话。
四、最简操作流程
综上所述，以下是 Tmux 的最简操作流程。

新建会话tmux new -s my_session。
在 Tmux 窗口运行所需的程序。
按下快捷键Ctrl+b d将会话分离。
下次使用时，重新连接到会话tmux attach-session -t my_session。
五、窗格操作
Tmux 可以将窗口分成多个窗格（pane），每个窗格运行不同的命令。以下命令都是在 Tmux 窗口中执行。

5.1 划分窗格
tmux split-window命令用来划分窗格。


# 划分上下两个窗格
$ tmux split-window

# 划分左右两个窗格
$ tmux split-window -h


5.2 移动光标
tmux select-pane命令用来移动光标位置。


# 光标切换到上方窗格
$ tmux select-pane -U

# 光标切换到下方窗格
$ tmux select-pane -D

# 光标切换到左边窗格
$ tmux select-pane -L

# 光标切换到右边窗格
$ tmux select-pane -R
5.3 交换窗格位置
tmux swap-pane命令用来交换窗格位置。


# 当前窗格上移
$ tmux swap-pane -U

# 当前窗格下移
$ tmux swap-pane -D
5.4 窗格快捷键
下面是一些窗格操作的快捷键。

Ctrl+b %：划分左右两个窗格。
Ctrl+b "：划分上下两个窗格。
Ctrl+b <arrow key>：光标切换到其他窗格。<arrow key>是指向要切换到的窗格的方向键，比如切换到下方窗格，就按方向键↓。
Ctrl+b ;：光标切换到上一个窗格。
Ctrl+b o：光标切换到下一个窗格。
Ctrl+b {：当前窗格与上一个窗格交换位置。
Ctrl+b }：当前窗格与下一个窗格交换位置。
Ctrl+b Ctrl+o：所有窗格向前移动一个位置，第一个窗格变成最后一个窗格。
Ctrl+b Alt+o：所有窗格向后移动一个位置，最后一个窗格变成第一个窗格。
Ctrl+b x：关闭当前窗格。
Ctrl+b !：将当前窗格拆分为一个独立窗口。
Ctrl+b z：当前窗格全屏显示，再使用一次会变回原来大小。
Ctrl+b Ctrl+<arrow key>：按箭头方向调整窗格大小。
Ctrl+b q：显示窗格编号。
六、窗口管理
除了将一个窗口划分成多个窗格，Tmux 也允许新建多个窗口。

6.1 新建窗口
tmux new-window命令用来创建新窗口。


$ tmux new-window

# 新建一个指定名称的窗口
$ tmux new-window -n <window-name>
6.2 切换窗口
tmux select-window命令用来切换窗口。


# 切换到指定编号的窗口
$ tmux select-window -t <window-number>

# 切换到指定名称的窗口
$ tmux select-window -t <window-name>
6.3 重命名窗口
tmux rename-window命令用于为当前窗口起名（或重命名）。


$ tmux rename-window <new-name>
6.4 窗口快捷键
下面是一些窗口操作的快捷键。

Ctrl+b c：创建一个新窗口，状态栏会显示多个窗口的信息。
Ctrl+b p：切换到上一个窗口（按照状态栏上的顺序）。
Ctrl+b n：切换到下一个窗口。
Ctrl+b <number>：切换到指定编号的窗口，其中的<number>是状态栏上的窗口编号。
Ctrl+b w：从列表中选择窗口。
Ctrl+b ,：窗口重命名。
七、其他命令
下面是一些其他命令。


# 列出所有快捷键，及其对应的 Tmux 命令
$ tmux list-keys

# 列出所有 Tmux 命令及其参数
$ tmux list-commands

# 列出当前所有 Tmux 会话的信息
$ tmux info

# 重新加载当前的 Tmux 配置
$ tmux source-file ~/.tmux.conf
八、参考链接
A Quick and Easy Guide to tmux
Tactical tmux: The 10 Most Important Commands
Getting started with Tmux
（完）
`

	// text = "西游记中的孙悟空在火烧赤壁中并没有接参与战斗，但他发挥了重要的作用。\n\n在小说中，孙悟空为了得到金箍棒，曾经大闹天宫，最终被玉皇大帝招安封为弼马温。得知职位低微后，孙悟空怒回花果山，并战胜李天王和哪吒三太子的讨伐，迫使玉帝封其为齐天大圣。孙悟空在天宫建有府邸，并奉旨管理蟠桃园。\n\n由于孙悟空的神通广大，他在火烧赤壁中发挥了重要的作用。在小说中，赤壁之战中，曹操率领大军攻打东吴，东吴军在赤壁之战惨败。在这场战争中，孙悟空曾经化身为火德星君，率领火部众神围攻曹操的军队，但由于曹操军队的弓箭手众多，孙悟空的火部众神被击败。\n\n虽然孙悟空没有直接参与火烧赤壁的战斗，但他在这场战争中发挥了重要的作用，为东吴军的胜利做出了贡献。"

	// 创建MongoDB客户端
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// 连接MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// 获取集合句柄
	coll := client.Database(database).Collection(collection)

	rand.Seed(time.Now().UnixNano())

	// 插入大量数据
	fail := 0
	success := 0
	for i := 1; i <= numDocuments; i++ {
		time.Sleep(200 * time.Millisecond)
		// time.Sleep(5 * time.Second)
		document := bson.M{
			"name":       uuid.New().String(),
			"age":        rand.Intn(1000),
			"shard":      rand.Intn(50),
			"text":       text,
			"created_at": fmt.Sprintf("%v", time.Now()),
			"timestamp":  time.Now().UnixMilli(),
		}

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		_, err = coll.InsertOne(ctx, document)
		if err != nil {
			fmt.Printf("time: %v insert: %d fail, err: %v\n", time.Now(), i, err)
			fail++
			cancel()
			continue
		}
		success++
		fmt.Printf("time: %v insert: %d success\n", time.Now(), i)
		cancel()
		if i%100 == 0 {
			fmt.Printf("-----------插入了 %d 条文档, 失败了 %d 条文档-----------\n", success, fail)
		}
	}
}

// GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main *.go
// go run *.go

// ```bash
// # 使用分片功能
// use walter
// db.chen.insert({name:"kant", age:32})
// db.chen.createIndex({name: 1}, {background: true})
// sh.enableSharding("walter")
// sh.shardCollection("walter.chen", { name: 1 })
// ```

// ```bash
// # 查看数据分布
// db.chen.getShardDistribution()
// ```

// ```bash
// db.chen.find().sort({_id: -1}).limit(1)
// db.chen.find({}, { age: 1, shard: 1, created_at: 1, timestamp:1 }).sort({_id: -1}).limit(3)

// db.chen.find({"timestamp": { $lte: 1718786607973 }}).count()
// db.chen.countDocuments({"timestamp": { $lte: 1718786607973 }})
// ```
