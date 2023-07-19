# Unique-Studio-Web-Summer-Camp

<details>
  <summary><h2>TASK 0 : Previous Knowledge</h2></summary> 
  
<details>
  <summary><h3>Daily Report</h3></summary>

#### DAY0

重温了下 markdown 和 git ，发任务前大致学习了 go （菜鸟教程写的真的是依托答辩），后续大概会学习推荐的 go 语言圣经。非常神奇没有 U 盘，校内店都关了，第二天去校外的店买个吧。

#### DAY1

被 grub 干碎的一天。

我也是和吴迪学长一样拥有 4 次 archlinux 安装经验的人辣~

电脑空间太小了最多只能压缩出 50G 安装了。。。

第一次下意识把根分区挂载到 ```/boot``` 上了，部署 grub 的时候报错。又自作聪明地把 EFI 分区挂载到新装系统的 ```/mnt/boot``` 上，之后又自作聪明地退回 LiveCD 模式瞎搞，结果后面弄得一团糟。（一切的源头就是这个小小的错误）

第二次成功辣！~~虽然 grub 配置文件没有 windows 入口~~ 图形界面选择了 Xorg+Gnome，KDE 貌似有点显示问题。再搞了一些配置然后想到，虽然 grub 没有 windows 入口，但是我可以在 BIOS 里改开机优先级呀，结果果然可以！然后。。。回到 linux 开机时，又报错 ```invalid cluster 0``` 应该是 windows 开机影响到了grub，后面找了许多方法去 liveCD 里修复 grub 无用，遂重装系统。

```
//大多数是这种方法修复
grub-install /dev/sdx // /dev/sdx为系统所在的设备，重新安装grub
grub-update //更新grub
```

第三次配置 grub 的时候，又报错说某个文件系统时只读的，是之前就装过 grub 的原因，所以找删除  grub 的方法。Linux 删除的方法和上述的大同小异，于是去找windows 平台的方法。一种要进入类似 LiveCD 的修复模式修复 MBR，还要再搞 windows11 的安装介质？？一种进入 diskpart 模式给 EFI 盘符去操作删除 grub 文件参考[这篇](http://www.manongjc.com/detail/63-ykqveghovllfpmf.html)。但是人家删的是 Ubuntu 的，直接给你建了一整个文件夹，但我的情况是各种文件分布在各个文件夹，虽然可以通过文件修改时间来判断，但是万一 grub 修改了 windows 的开机文件呢？中间又试了各种不奏效的方法。最后精神有点失常还是把 EFI 区的一些文件删了，真的幸亏电脑还能开机。

第四次挺顺利，大概三四十分钟就装好图形界面了。 grub 已经被我当成病毒一样的存在了，还特意新建了一个新的引导分区挂载。（后知后觉地发现直接新建也许不用删原来分区上的 grub 吧）两个系统试着切换了下没问题。

啥事没干，光倒腾装系统了。呵呵，Arch 永远的神！

#### DAY2

电脑坏了。

开始把 linux 配置基本都搞好后关了一次机，电源适配器还连着，宿舍电突然停了又开，然后就开不了机了。。。开始以为是静电原因，用经典方法无果，后面拿去校内维修店告诉我主板烧了。。。

后面电脑得拿去保修了，现在是在手机浏览器上编辑。

有点郁闷，之后就结合鸟哥的书再学习了一些linux 知识。计划后面几天就纸上谈兵地学下 go 了。电脑一下修不好的话……希望能早点修好吧。

#### DAY3

拜读了 Go 语言圣经，这本书不单单是讲语法，有很多拓展实战的内容，读得还是蛮开心的。再补些 http 知识和 Go 相关函数可以尝试去写点小程序吧。 Go 没有安排明确的任务，我就写点有意思的东西在那一块了。电脑保修最早预约在了十三号上午，祈祷 ing……

#### DAY4

受不鸟啦！电脑修不好只能晚上去网吧通宵了，去阿里云申请了个 ECS 又开始配环境。。。然后大致学了下 Shell 和 构建工具，又是摸鱼的一天呢~

#### DAY5

世界上还有我这么蠢的人吗。。。早上离开网吧的时候就关了机没结账结果就把我剩下六十多元全扣掉了。。。

看了下 docker，感觉没啥好说的www

#### DAY6

作息已经完全颠倒了，现在是九点睡三点醒。。。下午通知我去取电脑，谢天谢地总算修好了。~~在网吧待了一晚衣服好多地方有黑点，完全洗不掉，还有浓重的烟味，幸好不用再去了 QwQ~~又再看了下 Go，想写一个爬 P 站图片的爬虫，写的过程有些曲折，本来还想再写一个 web 服务的，但没时间了。。。

#### DAY7

按现在的作息 DAYn 是指第 n 天的下午到第 n+1 天的早晨。。。晚上开完会就开始摸鱼然后早点睡觉调整作息吧。。。

</details>

---

<details>
<summary><h3>Linux</h3></summary>

- [x] 安装 archlinux

个人感受：
- 不管之前引导分区是否存在，最好新建一个。
- 安装图形界面前联网推荐手机usb共享网络。
- 永远尊重 grub

- [x] 了解发⾏版的包管理器的使⽤(pacman/yay)
- [x] 能使⽤命令⾏进⾏对⽂件或⽂件夹的创建、复制、删除、搜索、移动、查看等操作(ls/cd/mkdir/touch/cp/rm/find/mv)
- [x] 了解 Linux ⽂件权限，会修改⽂件权限、⽂件所有者
- [x] 了解 SUID/SGID/SBIT 和隐藏权限(chattr/lsattr)
- [x] 了解进程
- [x] 会查看、结束进程(ps/top/kill/killall/pkill)
- [x] 了解端口
- [x] 学习如何查看系统中端⼝占⽤的情况
- [x] 了解守护进程的概念及其管理⽅式
- [x] 配置 Shell
- [x] 了解 Linux 的防⽕墙机制

</details>

---

<details>
  <summary><h3>文本处理</h3></summary>

- [x] 学习简单的标记语言：**markdown**
- [x] 学习基本的⽂本处理命令（less/more/tail/cat）
- [x] ⼀些⾼级的⽂本处理命令(grep/awk/sed)
- [x] Vim

</details>

---

<details>
  <summary><h3>版本控制 Git</h3></summary>

- [x] Git 是什么，有什么用？
> Git 是一种分布式版本控制系统（Version Control System），能够管理跟踪计算机文件的版本和变化，并协调多人对同一代码库的开发。

- [x] 拥有一个 Github 账号
- [x] 学习简单的 git 操作，如 add, commit, branch, status 等
- [x] 学习如何回退版本（了解三种不同的回退模式）
```
git reset --soft //仅回退 HEAD 指针，workspace 和 index 不变
git reset --mixed //回退 HEAD 指针和 index，即回到 git commit 之前的状态
git reset --hard //回退 HEAD 指针，index 和 workspace，即回退到 git add 之前的状态
```

- [x] 查看 git 日志？
```
git log //查看被回退修正过的日志
git reflog //查看全部日志，包括回退操作
git log --graph --pretty=oneline //可查看分支合并情况
```
- [x] 如何修改 git commit 信息? 
```
git commit --amend //修改最近一次提交的信息

git rebase -i HEAD~n //修改倒数 n 条的信息，将需要修改的提交的 pick 改为 edit
git commit --amend //修改该次提交的信息
git rebase --continue //退出 rebase 交互
```

- [x] 学习如何合并分支
- [x] 学习如何暂存工作区
```
git stash save "message" //保存当前修改，并添加一个描述信息。
git stash list //列出所有保存的修改。
git stash apply stash@{n} //将指定的保存的修改应用到当前分支,默认最近一次
git stash drop stash@{n} //删除指定的保存的修改，默认最近一次
git stash pop //将最近一次修改应用到分支并删除
```
- [x] 考虑多个上游的管理
```
git remote add <name> <url> //添加一个名为 <name> 的远程仓库，并指定其 URL。
git remote remote <name> //删除名为 <name> 的远程仓库。
git remote -v //列出当前仓库中已经存在的远程仓库，并显示其 URL。
git remote show <name>//查看远程仓库的详细信息
```
- [ ] 看 Pro Git 深入了解⼀下 git 

</details>

---

<details>
  <summary><h3>Shell 和构建工具</h3></summary>

- [x] 能熟悉使用管道，I/O 重定向等 Shell 内置功能。

```
ls | head -n X | tail -n Y //输出文件夹前 X-Y+1 到 X 个文件的文件名
ls -t | head -n X //输出文件夹最新的 X 个文件
ls -l --time-style="+%Y-%m-%d-%H:%M:%S" | sed '1d' | awk '{print $6 " " $7}' | sort -t ' ' -k 1 -r | head -n X | awk '{print $2}' //我真是有够无聊的 

```
- [x] 常用构建工具的使用(Makefile)

```
NAME=main

.PHONY=build
build:
      go build -o ${NAME} main.go

.PHONY=run
run:
      ./${NAME} ${ARGS}

.PHONY=clean
clean:
      go clean

.PHONY=start
start:build run

```
</details>

---

<details>
  <summary><h3>Go</h3></summary>

- 变量的赋值和运算要求类型相同，常量为了方便设定成无类型。
- 取模结果的符号与被取模的数一致，除法取整向零方向，负数为算术右移。
- 存在长度为 0 但不等于 nil 的slice。
- 之前对于键值对给数组赋值的方法试验了下：如果在键值对赋值后只用值赋值，则默认键为前面的键加一，不能一次对同一位置赋值多次。
- Go 是因为没有继承所以搞了个匿名成员的东西来方便编程嘛……
- 闭包！老生常谈的东西了……
- 在实现接口的方面，一个类型的值不拥有其指针的方法，但一个类型的指针拥有其值的方法。
- 接口值为 nil 当且仅当其类型和值都为 nil。

</details>

---

<details>
  <summary><h3>Docker</h3></summary>

- [x]  了解 docker

> Docker是一种开源的容器化平台，允许用户将应用程序打包成一个独立的、可移植的容器，然后在任何环境中运行，提供了一种简单、快速、可靠和可移植的方式来打包、部署和运行应用程序。

- [x]  镜像/容器

> 镜像是一个静态的模板，容器是镜像的可运行实例，类似类和实例。但感觉镜像抽象程度也没类那么高（

- [x] 构建镜像/启动容器

> 可以编写 Dockerfile 文件构建或者从容器导出镜像

```
docker export container | docker import - image //从容器导出镜像

docker run //启动容器
```

- [x] 简单的 Dockerfile 编写

```
#配置 golang，并在创建时运行 main.go
FROM centos:7
RUN yum install -y wget \
      && wget https://golang.google.cn/dl/go1.14.4.linux-amd64.tar.gz \
      && tar -zxf go1.14.4.linux-amd64.tar.gz -C /usr/local
ENV GOROOT=/usr/local/go 
ENV PATH=$PATH:$GOROOT/bin
COPY main.go /root/Go/
ENTRYPOINT ["go","run","/root/Go/main.go"]
CMD [""]
```

- [x] 查看现有容器的状态

```
docker ps -a
```

- [x] 如何进入一个容器

```
docker exec -i -t name/id /bin/bash
```

- [x] 如何停止一个容器

```
docker stop name/id
```

- [x] docker-compose
i
> Docker Compose是一个用于定义和运行多个Docker容器的工具，可以通过一个单一的YAML文件来描述容器之间的关系、配置和依赖。基本就是 docker 版的 Makefile 了。

- [x] 简单的 docker-compose.yml 的编写

```
#目前对于各种服务理解不深，只知道大致写法不知道实际该怎么编写。

#docker-compose版本
version:'3'

#各个服务
services:
#服务名称
service_name:
  #容器生成方式有两种
  image: #指定镜像
  build: #用指定目录的 Dockerfile 生成
  
  container_name: #容器名
  restart: #重启策略
  volumes: #挂载路径设置
    - "path1:path2:rw/ro" #读写/只读
    - ...
    - ...
  depends_on: #容器依赖
    - service_name #服务名
    - ...
    - ...
  enviroments: #环境变量
    - key:value
    - ...
    - ...
  links: #连接的服务
    - service_name #服务名
  ports: #映射到宿主机的端口
    - "port" #随机映射
    - "port1:port2" #指定映射
    - ...
  expose: #暴露端口，被连接的服务访问
    - port
    - ...
    - ...
```

- [x] docker network

> Docker网络是Docker引擎提供的一种功能，它允许Docker容器之间进行通信和连接，并提供一种隔离和安全的网络环境。


</details>

---

</details>



<details>
  <summary><h2>TASK 1 : Anonymous Forum Web Application Backend</h2></summary>

<details>
  <summary><h3>Daily Report</h3></summary>

#### DAY1

大致看了下任务和提供的链接，有了整体的了解。后面计划：

1. 设计数据库，学习用 Gorm 框架来与数据库交互。
2. 根据要求设计 API 接口。
3. 学习用 Gin 框架具体实现各种 API 接口。
4. 用 HTTPie 测试，空余时间多的话为了交互方便可能会再写个客户端。
5. 用 docker-compose 构建服务，加个 web 和 数据库的网络通信应该就差不多了。

emmmm 感觉好多啊。。。今天还又摸了 QwQ 

希望顺利吧。

#### DAY2

学了下 Gorm，开始码代码。

关于帖子的表的设计有点犹豫，一种是将帖子和评论看成一种东西存在一起，但是帖子还是有标题，链接之类额外的性质，空间会有很多浪费，但是大体比较方便；一种是将帖子和评论分离，api 就不是很统一。对于举报操作，因为帖子和评论都没区别，所以设计成前者会方便一点。对于匿名操作，因为是根据帖子获得匿名，所以设计成后者会方便点。最终还是选择了后者。

#### DAY3

数据库部分代码应该差不多了，后面开始学习 Gin 框架。

</details>

</details>

