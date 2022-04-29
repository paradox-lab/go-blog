# Git
**编辑时间**:2021.5.19

**官网**: https://git-scm.com/

**Docs**: https://git-scm.com/doc

**Docs中文版**: https://git-scm.com/book/zh/v2

## 安装
CentOS
> 教程: https://www.cnblogs.com/jhxxb/p/10571227.html

# 第一部分 基础
## 工作区、暂存区与存储库
```text
在Git中，针对工作目录、暂存区以及存储库3个主要区域，可以通过不同的Git命令，
把文件移往不同的区域。
(1)git add命令可以把文件从工作目录移至暂存区。
(2)git commit命令可以把暂存区的内容移至存储库。

如果觉得先add再commit有点烦琐，所以在Commit时多加一个-a参数，缩短这个流程：
git commit -a -m "update content"

这样即使没有先add，也可以完成Commit。但要注意的是，这个-a参数只对已经存在的
Repository区域的文件有效，对新加入的文件是无效的。

```

## 用户设置
```commandline
git config --global user.name "luzhenxiong"
git config --global user.email "397132445@qq.com"
```

**检查设置**
```commandline
git config --list
```
> 所有Git相关的设置都默认保存在自己账号下的.gitconfig文件中。

## 基本使用
### 创建仓库
```commandline
git init
```

### add
```commandline
git add welcome.html
```
_增加多个文件_
```commandline
git add *.html
```
_增加全部文件_
```commandline
git add -all
```
或
```commandline
git add .
```

### commit
```commandline
git commit -m "init commit"
```
> 加上了--allow-empty参数，没有注释也可以commit
```commandline
git commit --allow-empty -m
```

#### 追加文件到最近一次的commit
> 方法一：使用git reset命令把最后一次的Commit删除，加入新文件后再重新Commit。

```text
方法二：使用--amend参数进行Commit。
例如，有一个名为cinderella.html的文件，想把它加入到最近一次的Commit中，假设
它状态是Untracked，先使用git add命令先把文件加到暂存区。
接着在Commit时加上--amend参数：
git commit --amend --no-edit
这样就可以把文件并入最近一次的Commit。最后面哪个--no-edit参数的意思是
"我不要编辑Commit信息"，这样不会跳出Vim编辑器窗口。
```

### 删除文件
```text
可以先执行rm命令删除，然后再执行git add命令加入暂缓区的两段式动作，也可以直接使用
git rm命令来完成：
git rm welcom.html

这样就直接在暂缓区了，不需要再add一次。

3.加上--cached参数
不论是执行rm命令，还是执行git rm命令，都会真的把这个文件从工作目录中删除。如果不是
真的想把这个文件删除，只是不想让这个文件再被Git控制了，可以加上--cached参数：
git rm welcom --cached

这样就不会真的把文件删了，而只是把文件从Git中移除了。
```

### 恢复被删除的文件
```text
例如，删除html文档
rm *.html

这时如果要把cinderella.html挽救回来，可以使用git checkout命令
git checkout cinderella.html

如果想把所有被删除的文件都挽救回来：
git checkout .
这样，刚刚被删除的文件都挽救回来了。


这个命令会把暂存区(Staging Area)中的内容或文件拿来覆盖工作目录中(Working Directory)
的内容或文件。因此，执行git checkout命令时，会把welcome.html文件或者
当前目录下的所有文件恢复到上一次Commit的状态。
如果在执行这个命令时多加了一个参数：
git checkout HEAD~2 welcome.html

那么距离现在两个版本以上的那个welcome.html文件就会被用来覆盖当前工作目录中的welcome.html文件，
但要注意，这同时也会更新暂存区的状态。
其中：
git checkout HEAD~2 .
这个命令的意思就是“用距离现在两个版本以上的文件来覆盖当前目录中的文件，同时更新暂缓
区中的状态”。
```

### 回滚commit
```text
如果想要拆掉最后一次的Commit，可以采用"相对"或"绝对"的做法。"相对"的做法是这样的：
git reset eb08594^
最后的那个"^"符号代表的是"前一次"，所以eb08594^是指eb08594这个Commit的前一次；
如果是eb08594^^，则是往前两次······以此类推。不过要倒退5此，通常是改写成:
git reset master ^
或
git reset HEAD^

这两个命令会得到一样的结果。

如果你很清楚要把当前的状态退回到哪个Commit，可以直接指明：
git reset ad15454

它就会切换到ad15454de这个Commit的状态。

·Reset模式
git reset命令可以搭配参数使用。常用参数有3个，分别是--mixed、--soft以及--hard。搭配
不同的参数，执行结果会有些微差别。
(1)mixed模式。
--mixed是默认的参数，如果没有另外加参数，git reset命令将使用--mixed模式。该模式会把
暂存区的文件删除，但不会影响工作目录的文件。也就是说，Commit拆出来的文件会留在工作
目录，但不会留在暂存区。

(2)soft模式。
这种模式下的Reset，其工作目录与暂存区的文件都不会删除，所以看起来就只有HEAD的移动
而已。因此，Commit拆出来的文件会直接放在暂存区。

(3)hard模式。
在这种模式下，无论是工作目录还是暂存区的文件，都会被删除。
```

## 分支
### Git Flow
> 根据Git Flow的建议，分支主要分为Master分支、Develop分支、Hotifx分支、
Release分支以及Feature分支，各分支负责不同的功能。其中Master分支和
Develop分支又被称为长期分支，因为它们会一直存在于整个Git Flow中，而其他
的分支大多会因任务结束而被删除。
#### Master
> Master分支。Master分支主要是用来存放稳定、随时可上线的项目版本。这个
分支的来源只能是从别的分支合并过来的，开发者不会直接Commit到这个分支。因为
是稳定版本，所以通常会在这个分支的Commit上大商版本号标签。
#### Develop
> Develop分支。Develop分支是所有开发分支中的基础分支，当要新增功能时，所有的
Feature分支都是从这个分支划分出去的。而Feature分支的功能完成后，也会合并到这个
分支。
#### Hotfix
> Hotfix分支。当在线产品发生紧急问题时，会从Master分支划分出一个Hotfx分支进行
修复。Hotfix分支修复完成之后，会合并回Master分支，同时合并到一份Develop分支。
#### Release
> Release分支。当认为Develop分支足够成熟时，就可以把Develop分支合并到Release分支，
在其中进行上线前的最后测试。测试完成后，Release分支将同时合并到Master及Develop
这两个分支中。
#### Feature
> Feature分支。如果要新增功能，就要使用Feature分支了。Feature分支都是从Develop分支
划分出来的，完成之后回合并回Develop分支。

### 使用分支
#### 查看
```commandline
git branch
```
#### 创建
```commandline
git branch dev
```
#### 更名
```commandline
git branch -m cat tiger
```
> 把cat分支改成tiger分支

#### 删除
```commandline
git branch -d dog
```
> 如果要删除的分支还没有被完全合并，Git会有贴心小提示，这时只需改用-D参数
```commandline
git branch -D dog
```

#### 切换
```commandline
git checkout cat
```
> 创建并切换
```commandline
git checkout -b sister
```

#### 合并
> 如果想要用master分支来合并dev分支，就要先切换回master分支：

```commandline
git checkout master
git merge dev
```
> 如果出现合并冲突，修改冲突文件，然后add&commit，重新执行合并操作

**遇到冲突时，忽略合并**
```text
例如master分支和qa分支的setting.py配置文件不一致，合并时这个文件保持原样，使用.gitattributes文件
```
_.gitattributes_
```text
aje\config\setting.py merge=ours
```

第一次使用先执行`git config --global merge.ours.driver true`

> 参考资料 https://git-scm.com/book/zh/v2/%E8%87%AA%E5%AE%9A%E4%B9%89-Git-Git-%E5%B1%9E%E6%80%A7的合并策略

#### Reabase
```text
假设有cat、dog及master这3个分支，并且切换至cat分支上。这时如果执行以下命令：
git merge dog
则会产生一个额外的commit来接合两边的分支。

在Git中还有一个命令叫作git rebase，也可以用来做与git merge命令类似的事情。中文含义
大致是重新定义分支的参考基准。
以上面这个例子来说，cat与dog两个分支的base都是master。接着使用git rebase命令
来"组合"cat和dog这两个分支：
git rebase dog

上述命令的含义大致就是现在重新定义我的参考基准，并且将使用dog分支作为我新的参考
基准。
```

## 标签
### 列出标签
```commandline
git tag
```
_如果只对特定的标签感兴趣，可以使用模式匹配查找，如：_
```commandline
git tag -l 'v1.8.5*'
```
> 在git的标签中，有两种类型的标签：轻量标签(lightweight)和附注标签(annotated)，它们的区别如下：

### 创建标签

- 轻量标签： 像是一个不会改变的标签，它只是特定提交的引用
- 附注标签： 是存储在git数据库中的一个完整对象，因此它可以包含打标签者的信息（名字和电子邮件地址）、打标签时间以及标签信息。
此外，它还可以被校验，且使用GNU Privacy Guard (GPG)签名与验证
  
> 至于使用哪一种，官方建议是使用附注标签，因为它可以拥有更多的信息。不过如果不想要保存那些信息，那么轻量标签也是OK的

#### 打轻量标签
> 打轻量标签很简单，只需要一个标签名字，而不需要任何额外的信息，即：git tag <TagName>，如下：

```commandline
git tag v1.4-lw

git show v1.4-lw
```

#### 打附注标签
> 创建附注标签，则需要增加-a参数表明它是一个附注标签，还需要增加-m参数指定信息。如下：

```commandline
git tag -a v1.4 -m 'version 1.4'
```
> 当应用git show时，完整的信息就会被展示出来了：

### 后期打标签
> 如果不想对当前提交节点打标签，而是想对过去提交节点打标签，那么就可以使用后期打标签功能。

> 我们可以记录下它的commit-hash,然后，使用如下这种语法：
```commandline
git tag -a v1.2 9fceb02d0ae598e95dc970b74767f19372d61af8
```

### 删除标签
> 如果要删除一个标签，则可以使用命令git tag -d <TagName>，如：
```commandline
git tag -d v1.4-lw
```

> 但是，这不会删除远程服务器中的标签，如果希望这种删除同步到远程服务器中，则需要使用git push <remote> :refs/tags/<TagName>这种语法，示例如下：

```commandline
git push origin :ref/tags/v1.4-lw
```

### 检出标签
> 我们可以使用git checkout来检出特定的标签。但是需要注意的是，这会导致头指针处于游离状态（detachthed HEAD），在这种状态下，如果做了某些更改并提交，那么标签不会发生任何变化。并且由于提交不属于任何分支，所以也无法访问这些提交。
因此，git checkout检出标签适用于对旧版本的只读状态，如果需要在旧版本上修复问题，则应该在检出之后，拉取新的分支，从而就不会有头指针游离状态带来的问题了

## 远程仓库
### 本地已有项目关联远程新建仓库
- 新建远程仓库，不要初始化README.ME和.ignore
- 本地设置远程节点
- git push 远程节点 master
- 如果推送失败，要求git pull，可以强制推送: git push --force --set-upstream origin master

### 设置远程节点
```commandline
git remote add origin https://gitee.com/luzhenxiong/Package.git
```

### 删除远程节点
```commandline
git remote rm origin
```

### 查看远程分支
```commandline
git remote -v
```
> 使用如下命令可以在本地新建一个temp分支，并将远程origin仓库的master分支代码下载到本地temp分支

### 删除远程分支
```commandline
git push origin --delete dev
```

### 拉取
#### fetch
```commandline
git fetch origin master:temp
```
> 使用如下命令来比较本地代码与刚刚从远程下载下来的代码的区别：
```commandline
git diff temp
```
#### pull
> git pull的作用是，从远程库中获取某个分支的更新，再与本地指定的分支进行自动merge。完整格式是：
> git pull <远程库名> <远程分支名>:<本地分支名>

```commandline
git pull origin develop:develop
```
```commandline
git pull origin
```
> git pull 命令等同于先做了git fetch ，再做了git merge。即：
```commandline
git fetch origin develop
git checkout develop
git merge origin/develop
```
或者
```commandline
git fetch origin master:tmp
git diff tmp 
git merge tmp
git branch -d tmp
```
### 推送
```commandline
git push origin master
```

> 如果首次推送出现错误：failed to push some refs to git  ，则需要先将远程仓库的中的README.md文件pull到本地，执行：
```commandline
git pull --rebase origin master
git push origin master
```

### 更换远程仓库地址
```commandline
git remote set-url origin [new-repo-url]
```

### 远程标签
> 默认情况下，本地打的标签是包含推送到远程服务器上的，如果需要推送到远程服务器上，则需要使用：

```commandline
git push origin <TagName>
```
> 如：git push origin v1.5，如果希望以此推送多个标签，则可以使用：

```commandline
git push --tags
```
> 这样子它就会将远程服务器上不存在的本地标签都推送到远程服务器中

## 查看历史记录
```text
1.使用Git命令
git log

可以看到以下内容：
1.Commit的作者是谁。
2.什么时候Commit的。
3.每次的Commit大概做了些什么事。

git log --oneline --graph
加上这个参数，可以看到不一样的输出格式。

想要找某个人或某些人的Commit
git log --online --author="Sherly"
此外，使用"\|可以查询两人的Commit记录"
git log --oneline --author="Sherly\|Eddie"

想要找Commit信息中是否含有某些关键字
使用--grep参数，可以从Commit信息中搜索符合关键字的内容，如搜索"LOL"：
git log --oneline --grep="LOL"

怎样在Commit文件中找到Ruby
使用-S参数，可以在所有的Commit文件中进行搜索，找到那些符合特定条件的内容。
git log -S "Ruby"

怎样查找某一段时间内的Commit
在查看历史记录时，可以搭配--since和--until参数查询：
git log --oneline --since="9am" --until="12am"
这样就可以找出"今天早上9点到12点之间所有的Commit"。还可以再加一个after：
git log --oneline --since="9am" --until="12am" --after="2017-01"
这样就可以找到"从2017年1月之后，每天早上9点到12点之间的Commit"
```

### 查看特定文件的commit记录
```commandline
git log welcome.html
```
> 这样就能看到这个文件Commit的历史记录。如果想查看这个文件每次的Commit做了什么
改动，可以再给它加上一个-p参数：

### 查看代码是谁写的
> 想要知道某个文件的某一行代码是谁写的，在Git中可使用git blame命令帮你找出来：
```commandline
git blame index.html
```
> 如果文件太大，也可以加上-L参数，只显示指定行数的内容：
```commandline
git blame -L 5,10 index.html
```
> 这样就会只显示5~10行的信息。

## stash
> 能够将所有未提交的修改（工作区和暂存区）保存至堆栈中，用于后续恢复当前工作目录。

应用场景：
```text
1 当正在dev分支上开发某个项目，这时项目中出现一个bug，需要紧急修复，但是正在开发的内容只是完成一半，还不想提交，这时可以用git stash命令将修改的内容保存至堆栈区，
然后顺利切换到hotfix分支进行bug修复，修复完成后，再次切回到dev分支，从堆栈中恢复刚刚保存的内容。
2 由于疏忽，本应该在dev分支开发的内容，却在master上进行了开发，需要重新切回到dev分支上进行开发，可以用git stash将内容保存至堆栈中，切回到dev分支后，再次恢复内容即可。
总的来说，git stash命令的作用就是将目前还不想提交的但是已经修改的内容进行保存至堆栈中，后续可以在某个分支上恢复出堆栈中的内容。这也就是说，
stash中的内容不仅仅可以恢复到原先开发的分支，也可以恢复到其他任意指定的分支上。git stash作用的范围包括工作区和暂存区中的内容，也就是说没有提交的内容都会保存至堆栈中。
```

```text
$ git stash
Saved working directory and index state WIP on master: b2f489c second

$ git status
On branch master
nothing to commit, working tree clean

2 git stash save
作用等同于git stash，区别是可以加一些注释，如下：
git stash的效果：

stash@{0}: WIP on master: b2f489c second
git stash save “test1”的效果：

stash@{0}: On master: test1
3 git stash list
查看当前stash中的内容

4 git stash pop
将当前stash中的内容弹出，并应用到当前分支对应的工作目录上。
注：该命令将堆栈中最近保存的内容删除（栈是先进后出）
顺序执行git stash save “test1”和git stash save “test2”命令，效果如下：

$ git stash list
stash@{0}: On master: test2
stash@{1}: On master: test1

$ git stash pop
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   src/main/java/com/wy/StringTest.java

no changes added to commit (use "git add" and/or "git commit -a")
Dropped refs/stash@{0} (afc530377eacd4e80552d7ab1dad7234edf0145d)

$ git stash list
stash@{0}: On master: test1
可见，test2的stash是首先pop出来的。
如果从stash中恢复的内容和当前目录中的内容发生了冲突，也就是说，恢复的内容和当前目录修改了同一行的数据，那么会提示报错，需要解决冲突，可以通过创建新的分支来解决冲突。

5 git stash apply
将堆栈中的内容应用到当前目录，不同于git stash pop，该命令不会将内容从堆栈中删除，也就说该命令能够将堆栈的内容多次应用到工作目录中，适应于多个分支的情况。

$ git stash apply
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   src/main/java/com/wy/StringTest.java

no changes added to commit (use "git add" and/or "git commit -a")

$ git stash list
stash@{0}: On master: test2
stash@{1}: On master: test1

堆栈中的内容并没有删除。
可以使用git stash apply + stash名字（如stash@{1}）指定恢复哪个stash到当前的工作目录。

6 git stash drop + 名称
从堆栈中移除某个指定的stash

7 git stash clear
清除堆栈中的所有 内容

8 git stash show
查看堆栈中最新保存的stash和当前目录的差异。

$ git stash show
 src/main/java/com/wy/StringTest.java | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

git stash show stash@{1}查看指定的stash和当前目录差异。
通过 git stash show -p 查看详细的不同：

$ git stash show -p
diff --git a/src/main/java/com/wy/CacheTest.java b/src/main/java/com/wy/CacheTest.java
index 6e90837..de0e47b 100644
--- a/src/main/java/com/wy/CacheTest.java
+++ b/src/main/java/com/wy/CacheTest.java
@@ -7,6 +7,6 @@ package com.wy;
  */
 public class CacheTest {
     public static void main(String[] args) {
-        System.out.println("git stash test");
+        System.out.println("git stash test1");
     }
 }
diff --git a/src/main/java/com/wy/StringTest.java b/src/main/java/com/wy/StringTest.java
index a7e146c..711d63f 100644
--- a/src/main/java/com/wy/StringTest.java
+++ b/src/main/java/com/wy/StringTest.java
@@ -12,7 +12,7 @@ public class StringTest {

     @Test
     public void test1() {
-        System.out.println("=================");
+        System.out.println("git stash test1");
         System.out.println(Strings.isNullOrEmpty(""));//true
         System.out.println(Strings.isNullOrEmpty(" "));//false
         System.out.println(Strings.nullToEmpty(null));//""

同样，通过git stash show stash@{1} -p查看指定的stash的差异内容。

9 git stash branch
从最新的stash创建分支。
应用场景：当储藏了部分工作，暂时不去理会，继续在当前分支进行开发，后续想将stash中的内容恢复到当前工作目录时，如果是针对同一个文件的修改（即便不是同行数据），那么可能会发生冲突，恢复失败，这里通过创建新的分支来解决。可以用于解决stash中的内容和当前目录的内容发生冲突的情景。
发生冲突时，需手动解决冲突。
```

## .gitignore
```text
有些文件不想放在Git中

有些比较机密的文件不想放在Git中一起备份，如数据库的存取密码或AWS服务器的存取金钥。

除了比较机密的文件，对于一些程序编译的中间文件或暂存文件，同样不想将其放在Git中。

1.忽略这个文件
只需在项目目录中放一个.gitignore文件，并且设置想要忽略的规则即可。如果这个文件不存在，
就手动新增它：
touch .gitignore

然后编辑这个文件的内容：
#文件名称 .gitignore

#忽略secret.yml文件
secret.yml

#忽略config目录下的database.yml文件
config/database.yml

#忽略db目录下所有后缀是.sqlite3的文件
/db/*.sqlite3

#忽略所有后缀是.tmp的文件
.tmp

#如果想要忽略.gitignore这个文件也可以，只是通常不会这么作
#.gitignore

只要这个文件存在，即使这个文件没有被Commit或Push上Git服务器，也会有效果。但
通常建议将这个文件Commit进项目并且push上Git服务器，以扁让一起开发项目的所有人
可以共享相同的文件。

虽然.gitignore文件列出了一些忽略的规则，但其实这些忽略的规则也是可以被忽略的。
只需在执行git add命令时加上-f参数：
git add -f 文件名称

.gitignore文件设置的规则只对那些在规则设置之后存入的文件有效，那些已经存在的文件
，这些规则对它是无效的。

如果想套用.gitignore的规则，就必须先使用git rm --cached命令，然后它们就会被忽略了。

清除忽略的文件
可以使用git clean命令并配合-X参数：
git clean -fx
哪个额外加上的-f参数是指强制删除。这样一来，就可以清除那些被忽略的文件了。
```

# 第二部分 fatal汇总
 - **refusing to merge unrelated histories**
   - 原因：本地仓库和远程仓库是两个独立不同的仓库
   - 解决：在pull命令后紧接着使用--allow-unrelated-history
   - cmd:git pull origin master --allow-unrelated-histories
