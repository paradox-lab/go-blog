# Linux

有以下分支:
## Ubuntu
> Ubuntu是一个以桌面应用为主的Linux操作系统，其名称来自非洲南部祖鲁语或豪萨语的“ubuntu"一词，意思是“人性”“我的存在是因为大家的存在"，是非洲传统的一种价值观。

**官网**: https://cn.ubuntu.com/

### apt-get更换国内源
> https://blog.csdn.net/qq_24326765/article/details/81916222

### 制作桌面快捷启动图标
> 在/usr/share/applications 新建xxx.desktop文件，填入类似以下内容
```text
[Desktop Entry]
Name=Python (v3.9)
Comment=Python Interpreter (v3.9)
Exec=/usr/bin/python3.9
Icon=/usr/share/pixmaps/python3.9.xpm
Terminal=true
Type=Application
Categories=Development;
StartupNotify=true
NoDisplay=true
```

> 给权限755，复制到桌面，右键点击允许启动

### 20.04.2
PASS

# 文件、目录与磁盘格式
## 压缩和打包

_常见的压缩文件和拓展名_
```text
*.Z        compress程序压缩的文件;
*.gz      gzip程序压缩的文件;
*.bz2    bzip2程序压缩的文件：
*.tar      tar程序打包的数据，并没有压缩过;
*.tar.gz tar程序打包的文件，其中经过gzip的压缩;
*.tar..bz2 tar程序打包的文件，其中经过bzip2的压缩;
```
> Linux上常见的压缩命令是gzip与bzip2


## 清空当前文件夹
```commandline
rm -rf *
```

# 常见问题：
- centos执行sudo命令提示找不到
> https://www.php.cn/centos/449396.html
- sudo避免`Permission denied’
> https://www.cnblogs.com/testlife007/p/6944136.html
```text
sudo 后面加sh -c命令， 它可以让 bash 将一个字串作为完整的命令来执行，这样就可以将 sudo 的影响范围扩展到整条命令。
```
```commandline
sudo sh -c "echo a > 1.txt"
```

- 加入环境变量
> https://www.cnblogs.com/chenqionghe/p/4295219.html