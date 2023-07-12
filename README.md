# Unique-Studio-Web-Summer-Camp

## 学习进度

### DAY0

重温了下 markdown 和 git ，发任务前大致学习了 go （菜鸟教程写的真的是依托答辩），后续大概会学习推荐的 go 语言圣经。非常神奇没有 U 盘，校内店都关了，第二天去校外的店买个吧。

### DAY1

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

### DAY2

电脑坏了。

开始把 linux 配置基本都搞好后关了一次机，电源适配器还连着，宿舍电突然停了又开，然后就开不了机了。。。开始以为是静电原因，用经典方法无果，后面拿去校内维修店告诉我主板烧了。。。

后面电脑得拿去保修了，现在是在手机浏览器上编辑。

有点郁闷，之后就结合鸟哥的书再学习了一些linux 知识。计划后面几天就纸上谈兵地学下 go 了。电脑一下修不好的话……希望能早点修好吧。

### DAY3

拜读了 Go 语言圣经，这本书不单单是讲语法，有很多拓展实战的内容，读得还是蛮开心的。再补些 http 知识和 Go 相关函数可以尝试去写点小程序吧。 Go 没有安排明确的任务，我就写点有意思的东西在那一块了。电脑保修最早预约在了十三号上午，祈祷 ing……

---

## Linux

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
---

## 文本处理

- [x] 学习简单的标记语言：**markdown**
- [x] 学习基本的⽂本处理命令（less/more/tail/cat）
- [x] ⼀些⾼级的⽂本处理命令(grep/awk/sed)
- [x] Vim


---

## 版本控制 Git

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

---

## Shell 和构建工具

---

## Go

- 变量的赋值和运算要求类型相同，常量为了方便设定成无类型。
- 取模结果的符号与被取模的数一致，除法取整向零方向，负数为算术右移。
- 存在长度为 0 但不等于 nil 的slice。
- 之前对于键值对给数组赋值的方法试验了下：如果在键值对赋值后只用值赋值，则默认键为前面的键加一，不能一次对同一位置赋值多次。
- Go 是因为没有继承所以搞了个匿名成员的东西来方便编程嘛……
- 闭包！老生常谈的东西了……

---

## Docker
