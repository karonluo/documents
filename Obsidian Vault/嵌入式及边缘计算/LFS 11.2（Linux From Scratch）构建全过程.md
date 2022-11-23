#T800 #嵌入式 #物联网 #边缘计算 #Linux #LFS #操作系统 

## 描述

人员知识要求：
1. 熟练 Linux 命令，包括不限于 sed、gwk、find、ln、正则表达式、grep 等相关知识
2. 熟练编译工作，包括不限于：`Python`、`C和C++` 、`CMake`、`gcc`、`g++`、`Go` 相关知识
3. 熟练虚拟机或容器相关知识，方便每个步骤进行提交备份

本文档不会对基础命令和语言语句进行解释，请参考不同的知识点
[[Linux常用SHELL命令大全]]
[[Linux命令之awk介绍]]
[[Linux命令之pushd和popd详解]]
[[C和C++语言]]
[[Python语言]]
[[Go语言]]

本文档，参考 LFS 相关文档逐步进行构建自己的 Linux 操作系统。
[LFS News (linuxfromscratch.org)](https://linuxfromscratch.org/news.html)
[Linux From Scratch - Version 11.2](https://linuxfromscratch.org/lfs/downloads/11.2/LFS-BOOK-11.2.pdf)
其中有必要解释的地方，我会进行一些翻译，有不当的地方，可以提出指正，我会尽量减少机翻带来的问题。

## 关键提醒

**构建过程应该在虚拟机或容器中进行，每个步骤完成后或过程中，尽可能的提交快照或新版本容器镜像，准备随时因为操作失误或者错误过程进行回滚。**

## 步骤

![[LFS 11.2 构建全过程-准备工作]]

![[LFS 11.2 构建全过程-下载软件包]]

![[LFS 11.2 构建全过程-最后的准备工作]]

![[LFS 11.2 构建全过程-交叉工具链的构建]]

![[LFS 11.2 构建全过程-交叉编译临时工具]]