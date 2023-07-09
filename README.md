# Unique-Studio-Web-Summer-Camp

## 学习进度

### DAY0
重温了下 markdown 和 git ，发任务前大致学习了 go （菜鸟教程写的真的是依托答辩），后续大概会学习推荐的 go 语言圣经。非常神奇没有 U 盘，校内店都关了，第二天去校外的店买个吧。

---

## Linux

---

## 文本处理

- [x] 学习简单的标记语言：**markdown**
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

- [x] 查看git日志？
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

---

## Docker
