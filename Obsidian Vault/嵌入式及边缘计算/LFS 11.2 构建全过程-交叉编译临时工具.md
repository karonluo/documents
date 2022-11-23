#T800 #嵌入式 #物联网 #边缘计算 #Linux #LFS #操作系统 
## 重要提醒

以下所有步骤都在 `lfs` 用户进行，请保证目录为 `/mnt/lfs` 变量 `$LFS=/mnt/lfs` 是存在的。
切记 在 `lfs` 用户中激活 `~/.bashrc` 以及 `~/.bashrc_profile`

~~~Shell
su lfs
source ~/.bash_profile
source ~/.bashrc
~~~
保证以下结果正确
~~~Shell
echo $LFS
/mnt/lfs
echo $LFS_TGT
x86_64-lfs-linux-gnu 
~~~

## 安装 M4 1.4.19
### 概念

m4 是 POSIX 标准中的一部分，所有版本的 UNIX 下都可用。虽然这种语言可以单独使用，但大多数人需要 m4 仅仅是因为 GNU autoconf 中的 “configure” [脚本](https://baike.baidu.com/item/%E8%84%9A%E6%9C%AC?fromModule=lemma_inlink)依赖它。宏处理器（或预处理器）一般用作文本替换工具。最终用户经常会用它来处理要反复使用的文本模板，典型的是用于编程工具，还会用于文本编辑和文字处理工具。

### 准备工作
请确保当前目录为 `$LFS/sources`
~~~Shell
cd $LFS/sources
tar xf m4-1.4.19.tar.xz
cd m4-1.4.19
~~~
### 编译工作
**启动配置**
**注意：本次无需先创建 build 目录**
~~~Shell
./configure --prefix=/usr \
 --host=$LFS_TGT \
 --build=$(build-aux/config.guess)
~~~

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
~~~

## 安装 Ncurses 6.3

### 概念

`Ncurses` 提供字符终端处理库，包括面板和菜单。
`Ncurses` 依赖于: `Bash`, `Binutils`, `Coreutils`, `Diffutils`, `Gawk`, `GCC`, `Glibc`, `Grep`, `Make`, `Sed`。

### 准备工作

~~~Shell
cd $LFS/sources
tar xf ncurses-6.3.tar.gz
cd ncurses-6.3
sed -i s/mawk// configure
~~~
关于  `sed -i s/mawk// configure` 该 sed 命令用于 替换 配置中 `mawk` 为 空，保障编译时不使用 `mawk` 而是优先 `awk`
`mawk` 需要安装，`awk` 是 `Linux` 默认的命令
扩展阅读:  [Linux命令之awk介绍](Linux命令之awk介绍.md)


### 编译工作
**启动配置和编译**
~~~Shell
cd $LFS/sources/ncurses-6.3
mkdir build
pushd build
  ../configure
  make -C include
  make -C progs tic
popd
~~~

扩展阅读: [[Linux命令之pushd和popd详解]]
**启动配置**
~~~Shell
./configure --prefix=/usr \
 --host=$LFS_TGT \
 --build=$(./config.guess) \
 --mandir=/usr/share/man \
 --with-manpage-format=normal \
 --with-shared \
 --without-normal \
 --with-cxx-shared \
 --without-debug \
 --without-ada \
 --disable-stripping \
 --enable-widec
~~~

**命令参数解析**
_`--with-manpage-format=normal`_
这防止 Ncurses 安装压缩的手册页面，否则在宿主发行版使用压缩的手册页面时，Ncurses 可能这样做。
_`--with-shared`_
该选项使得 Ncurses 将 C 函数库构建并安装为共享库。
_`--without-normal`_
该选项禁止将 C 函数库构建和安装为静态库。
_`--without-debug`_
该选项禁止构建和安装用于调试的库。
_`--with-cxx-shared`_
该选项使得 Ncurses 将 C++ 绑定构建并安装为共享库，同时防止构建和安装静态的 C++ 绑定库。
_`--without-ada`_
这保证不构建 Ncurses 的 Ada 编译器支持，宿主环境可能有 Ada 编译器，但进入 **chroot** 环境后 Ada 编译器就不再可用。
_`--disable-stripping`_
该选项防止构建过程使用宿主系统的 **strip** 移除调试符号。对交叉编译产生的程序使用宿主工具可能导致构建失败。
_`--enable-widec`_
该选项使得宽字符库 (例如 `libncursesw.so.6.3`) 被构建，而不构建常规字符库 (例如 `libncurses.so.6.3`)。宽字符库在多字节和传统 8 位 locale 中都能工作，而常规字符库只能在 8 位 locale 中工作。宽字符库和普通字符库在源码层面是兼容的，但二进制不兼容。

**启动编译并安装**

~~~Shell
make -j4
~~~

本次安装有特殊选项，请安装以下命令进行
~~~Shell
make DESTDIR=$LFS TIC_PATH=$(pwd)/build/progs/tic install
~~~
`TIC_PATH=$(pwd)/build/progs/tic`
我们需要传递刚刚构建的，可以在宿主系统运行的 **tic** 程序的路径，以保证正确创建终端数据库。

~~~Shell
echo "INPUT(-lncursesw)" > $LFS/usr/lib/libncurses.so
~~~
`echo "INPUT(-lncursesw)" > $LFS/usr/lib/libncurses.so`
我们很快将会构建一些需要 `libncurses.so` 库的软件包。创建这个简短的链接脚本。

## 安装 Bash 5.1.16

### 准备工作
请确保当前目录为 `$LFS/sources`
~~~Shell
cd $LFS/sources
tar xf bash-5.1.16.tar.gz
cd bash-5.1.16
~~~

### 编译工作

**提醒本次不需要创建 build 目录，直接进行配置和编译即可**

**启动配置**
~~~Shell
./configure --prefix=/usr                   \
            --build=$(support/config.guess) \
            --host=$LFS_TGT                 \
            --without-bash-malloc
~~~
**命令参数解析**
_`--without-bash-malloc`_
该选项禁用 `Bash` 自己的内存分配 (`malloc`) 函数，因为已知它会导致段错误。这样，`Bash` 就会使用 `Glibc` 的更加稳定的 `malloc` 函数。

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
ln -sv bash $LFS/bin/sh
~~~
`ln -sv bash $LFS/bin/sh` 目的是为那些使用 **sh** 命令运行 shell 的程序考虑，创建一个链接

## 安装 Coreutils 9.1

### 概念

Coreutils 软件包包含用于显示和设定系统基本属性的工具。
Coreutils 是 GNU 下的一个软件包，包含 Linux 下的 `ls` 等常用命令。
这些命令的实现要依赖于 `shell` 程序。

### 准备工作
请确保当前目录为 `$LFS/sources`
~~~Shell
cd $LFS/sources
tar xfv coreutils-9.1.tar.xz
cd coreutils-9.1
~~~

### 编译工作

**提醒本次不需要创建 build 目录，直接进行配置和编译即可**

**启动配置**
~~~Shell
./configure --prefix=/usr                     \
            --host=$LFS_TGT                   \
            --build=$(build-aux/config.guess) \
            --enable-install-program=hostname \
            --enable-no-install-program=kill,uptime
~~~
**命令参数解析**
`--enable-install-program=hostname`
该选项表示构建 **hostname** 程序并安装它 —— 默认情况下它被禁用，但 Perl 测试套件需要它。

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
~~~

**后续处理**
~~~Shell
mv -v $LFS/usr/bin/chroot $LFS/usr/sbin
mkdir -pv $LFS/usr/share/man/man8
mv -v $LFS/usr/share/man/man1/chroot.1 $LFS/usr/share/man/man8/chroot.8
sed -i 's/"1"/"8"/' $LFS/usr/share/man/man8/chroot.8
~~~
将程序移动到它们最终安装时的正确位置。在临时环境中这看似不必要，但一些程序会硬编码它们的位置，因此必须进行这步操作。

## 安装 Diffutils 3.8

### 概念

Diffutils 软件包包含显示文件或目录之间差异的程序。
即 Linux 中的 `diff` 命令

### 准备工作
请确保当前目录为 `$LFS/sources`
~~~Shell
cd $LFS/sources
tar xfv diffutils-3.8.tar.xz
cd diffutils-3.8
~~~

### 编译工作

**提醒本次不需要创建 build 目录，直接进行配置和编译即可**

**启动配置**
~~~Shell
./configure --prefix=/usr --host=$LFS_TGT
~~~

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
~~~

