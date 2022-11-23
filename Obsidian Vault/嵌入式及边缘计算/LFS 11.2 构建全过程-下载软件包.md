#T800 #嵌入式 #物联网 #边缘计算 #Linux #LFS #操作系统 

## 准备下载的路径

~~~Shell
export LFS="/mnt/lfs"
mkdir -p $LFS/sources
cd $LFS
~~~

> 注意：以下步骤均在 `/mnt/lfs` 目录中进行操作

## 准备下载软件包列表
~~~Shell
wget https://www.linuxfromscratch.org/lfs/downloads/stable/wget-list
~~~

## 开始下载软件包

~~~Shell
wget --input-file=wget-list --continue --directory-prefix=$LFS/sources
~~~

> 注意1：wget-list 中的软件包可能找不到，请替换成可用的版本，往往通过 wget-list 中的地址，找到相关的其他版本的文件，尽量使用更高版本的软件包。
> 注意2：有些软件包需要互联网，国内局域网需要科学上网才能正常下载。

### 下载 md5 检查文件用于检查文件包的完整性
~~~Shell
wget https://www.linuxfromscratch.org/lfs/downloads/stable/md5sums --continue --directory-prefix=$LFS/sources
~~~

### 执行 md5 检查
~~~shell
md5sum -c md5sums
~~~

拿到结果后，可以根据结果，单独下载不争取的软件包。

