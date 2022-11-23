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

## 安装 Binutils-2.39

### 准备工作

`cd` 到 `sources` 目录中，使用以下命令解压并创建和进入编译目录
~~~Shell
cd $LFS/sources
tar xf binutils-2.39.tar.gz
cd binutils-2.39
make build
cd build
~~~

### 编译工作
+ **启动配置**
~~~Shell
../configure --prefix=$LFS/tools \
 --with-sysroot=$LFS \
 --target=$LFS_TGT \
 --disable-nls \
 --enable-gprofng=no \
 --disable-werror
~~~
+ **启动编译并安装**
~~~Shell
make -j4
make install
~~~

安装完成后，应该可以在
`/mnt/lfs/tools/x86_64-lfs-linux-gnu/bin` 
`/mnt/lfs/tools/x86_64-lfs-linux-gnu/lib `
目录中看到相关文件

## 安装 gcc-12.2.0 - 步骤 1

### 准备工作

**注意：gcc 是最最重要的软件包，它用于系统编译的，是 c 语言的编译工具包，因此需要严格按照步骤进行，在这之前请提交快照或容器镜像，以便于错误回滚。**

**这里有大坑，我已经验证了该问题，gcc-12.2.0 目录中并不包含 gmp mpfr mpc 这仨是专门的组件包，需要独立解包并复制到 gcc-12.2.0 目录中** 

**后续: 新的 LFS-BOOK 11.2 版本已经修改这部分内容，作为提醒**

请确保当前目录是 `/mnt/lfs/sources`
~~~Shell
cd $LFS/sources
tar xf gcc-12.2.0.tar.xz
tar xf mpfr-4.1.0.tar.xz
mv mpfr-4.1.0 mpfr
mv mpfr gcc-12.2.0
tar xf gmp-6.2.1.tar.xz
mv gmp-6.2.1 gmp
mv gmp gcc-12.2.0
tar xf mpc-1.2.1.tar.gz
mv mpc-1.2.1 mpc
mv mpc gcc-12.2.0
cd gcc-12.2.0
~~~

注意：`x86_64` 需要执行设置64位库的默认目录名为 `lib`  下面的代码用于修改配置文件，配置文件：
`/mnt/lfs/sources/gcc-12.2.0/gcc/config/i386/t-linux64`
先备份一个文件 `/mnt/lfs/sources/gcc-12.2.0/gcc/config/i386/t-linux64.orig`
再修改替换 `m64=../lib64 为 m64=../lib`

~~~Shell
case $(uname -m) in
 x86_64)
 sed -e '/m64=/s/lib64/lib/' \
 -i.orig gcc/config/i386/t-linux64
 ;;
esac
~~~


创建并进入 `build` 目录 `/mnt/lfs/sources/gcc-12.2.0/build`
~~~Shell
mkdir build
cd build
~~~

### 编译工作
请确保当前目录为: `/mnt/lfs/sources/gcc-12.2.0/build`

**启动配置**

~~~Shell
../configure \
 --target=$LFS_TGT \
 --prefix=$LFS/tools \
 --with-glibc-version=2.36 \
 --with-sysroot=$LFS \
 --with-newlib \
 --without-headers \
 --disable-nls \
 --disable-shared \
 --disable-multilib \
 --disable-decimal-float \
 --disable-threads \
 --disable-libatomic \
 --disable-libgomp \
 --disable-libquadmath \
 --disable-libssp \
 --disable-libvtv \
 --disable-libstdcxx \
 --enable-languages=c,c++
 ~~~

**命令参数解析**
+ `--prefix` 安装路径前缀为 `/mnt/lfs/tools`
+ `--target` 安装路径结合前缀，指定在  `/mnt/lfs/tools/x86_64-lfs-linux-gnu`
+ `--with-glibc-version` 指定 `glibc` c语言编译 `libc` 库版本为 `2.36` 这个选项指定了将在目标上使用的glibc的版本。它与主机的libc没有关系，因为所有由 pass1 gcc 编译的文件都将在 `chroot` 环境中运行，而 `chroot` 环境与主机发行版的 `libc` 的环境中运行。
+ `--with-sysroot` 指定系统根目录为 `/mnt/lfs` 
+ `--with-newlib` 由于目前还没有可用的C语言库，这就确保了在构建时定义了 `inhibit_libc` 常量 `libgcc`。这可以防止编译任何需要 `libc` 支持的代码。
+ `--without-headers` 指定不使用 `headers` 在创建一个完整的交叉编译器时，GCC需要与目标系统兼容的标准头文件。对于我们的目的是不需要这些头文件。这个开关可以防止GCC寻找它们。
+ `--disable-libatomic` 
+ `--disable-libgomp`
+ `--disable-libquadmath`
+ `--disable-libssp`
+ `--disable-libvtv`
+ `--disable-libstdcxx`
+ `--disable-decimal-float`
+ `--disable-threads`
以上开关关闭了针对 decimal-float、threads、libatomic、libgomp, libquadmath、libssp、libvtv，以及 C++标准库(libstdcxx)。这些功能在构建交叉编译器时将无法编译，构建交叉编译器时，这些特性将无法编译，而且对于交叉编译临时libc的任务来说也不是必须的。
+ `--disable-shared` 这个开关迫使 `GCC` 静态地链接其内部库。我们需要这样做，因为共享库需要 `glibc`，而目标系统上还没有安装它。
+ `--disable-multilib` 在 `x86_64 `上，LFS 不支持 `multilib` 配置。这个开关对 `x86` 来说是无害的。
+ `--enable-languages` 这个选项确保只构建 C 和 C++ 编译器。这些是现在唯一需要的语言。

**启动编译并安装**

~~~Shell
make -j4
make install
~~~

这次编译的 `GCC` 安装了几个内部系统头文件。
通常情况下，其中的 `limit.h`，会反过来包括相应的系统 `limit.h` 头文件，在本例中是 `$LFS/usr/include/limits.h`, 但目前并不存在而是由多个独立文件组成。
多个独立文件并不包括系统头的扩展特性，这对于构建 `glibc` 没问题，但以后需要完整的内部头文件，因此我们需要使用以下命令创建完整版本的内部头文件生成到 $LFS/tools/lib/gcc/x86_64-lfs-linux-gnu/12.2.0/install-tools/include/limits.h 这样就与 `GCC` 的构建系统在正常情况下一样了。

~~~Shell
cd ..
cat gcc/limitx.h gcc/glimits.h gcc/limity.h > \
 `dirname $($LFS_TGT-gcc -print-libgcc-file-name)`/install-tools/include/limits.h
~~~

## 安装  Linux-5.19.2 API Headers

### 准备工作

**注意： Linux-5.19.2 API Headers 既是 Linux 内核头 包含所有的应用程序接口（API），因此需要严格按照步骤进行，在这之前请提交快照或容器镜像，以便于错误回滚。**

请确保当前目录为: `/mnt/lfs/sources`
~~~Shell
cd $LFS/sources
tar xf linux-5.19.2.tar.xz
cd linux-5.19.2
~~~

### 编译工作
请确保当前目录为: `/mnt/lfs/sources/linux-5.19.2`

~~~Shell
make mrproper 
~~~
该命令确保内核源代码树绝对干净，`Linux 内核开发组` 建议在每次编译内核前运行该命令

~~~Shell
make headers
~~~
该命令编译所有的内核头文件

~~~Shell
find usr/include -type f ! -name '*.h' -delete
~~~
该命令通过 `find` 命令找到所有非头文件并删除

~~~Shell
cp -rv usr/include $LFS/usr
~~~
该命令将 `/mnt/lfs/sources/linux-5.19.2/usr/include` 中的所有文件复制到 `/mnt/lfs/usr` 目录中

附表
|头文件路径|API 解释|
|:---|:---|
|`/usr/include/asm/*.h`|The Linux API ASM Headers|
|`/usr/include/asm-generic/*.h`|The Linux API ASM Generic Headers|
|`/usr/include/drm/*.h`|The Linux API DRM Headers|
|`/usr/include/linux/*.h`|The Linux API Linux Headers|
|`/usr/include/misc/*.h`|The Linux API Miscellaneous Headers|
|`/usr/include/mtd/*.h`|The Linux API MTD Headers|
|`/usr/include/rdma/*.h`|The Linux API RDMA Headers|
|`/usr/include/scsi/*.h`|The Linux API SCSI Headers|
|`/usr/include/sound/*.h`|The Linux API Sound Headers|
|`/usr/include/video/*.h`|The Linux API Video Headers|
|`/usr/include/xen/*.h`|The Linux API Xen Headers|

## 安装 glibc-2.36 

### 准备工作

**注意：`glibc` 软件包包含主 `C` 库。这个库提供了分配内存的基本例程，搜索
目录、打开和关闭文件、读写文件、字符串处理、模式匹配、数学公式等等。因此需要严格按照步骤进行，在这之前请提交快照或容器镜像，以便于错误回滚。**

请确保当前目录为: `/mnt/lfs/sources`
~~~Shell
cd $LFS/sources
tar xf glibc-2.36.tag.xz
cd glibc-2.36
~~~
以上命令是解压 `glibc-2.36` 并进入 `glibc-2.36` 目录

~~~Shell
case $(uname -m) in
 i?86) ln -sfv ld-linux.so.2 $LFS/lib/ld-lsb.so.3
 ;;
 x86_64) ln -sfv ../lib/ld-linux-x86-64.so.2 $LFS/lib64
 ln -sfv ../lib/ld-linux-x86-64.so.2 $LFS/lib64/ld-lsb-x86-64.so.3
 ;;
esac
~~~
以上命令是为了根据系统（i386\[x86\]或x86_64) 进行配置链接 `ln -sfv` 将 `ld-linux` 不同版本的 `so` 链接到 `LFS` 系统中的文件中
如果你的环境是 `x86_64` 则会有以下输出结果:

~~~Shell
'/mnt/lfs/lib64/ld-linux-x86-64.so.2' -> '../lib/ld-linux-x86-64.so.2'
'/mnt/lfs/lib64/ld-lsb-x86-64.so.3' -> '../lib/ld-linux-x86-64.so.2'
~~~

`LFS-BOOK` 原文：
~~~Text

创建一个符号链接以符合LSB标准。此外，对于x86_64，创建一个兼容性符号链接，这对于动态库加载器的正常运行是必要的，以保证动态库加载器的正常运行
~~~

~~~Shell
patch -Np1 -i ../glibc-2.36-fhs-1.patch
~~~
以上命令为 `glibc-2.36` 打一个补丁，返回如下结果：
~~~Text
patching file Makeconfig
Hunk #1 succeeded at 246 (offset -4 lines).
patching file nscd/nscd.h
Hunk #1 succeeded at 160 (offset 48 lines).
patching file nss/db-Makefile
Hunk #1 succeeded at 21 (offset -1 lines).
patching file sysdeps/generic/paths.h
patching file sysdeps/unix/sysv/linux/paths.h
~~~

创建并进入 `build` 目录
~~~Shell
mkdir build
cd build
~~~

生成配置文件 `configparms` 保证 根目录指定在 `/usr/sbin` 中
~~~Shell
echo "rootsbindir=/usr/sbin" > configparms
~~~

### 编译工作

请确保当前目录为： `/mnt/lfs/sources/glibc-2.36/build`

**启动配置**
~~~Shell
../configure \
 --prefix=/usr \
 --host=$LFS_TGT \
 --build=$(../scripts/config.guess) \
 --enable-kernel=3.2 \
 --with-headers=$LFS/usr/include \
 libc_cv_slibdir=/usr/lib
 ~~~

**命令参数解析**
+ `--host=$LFS_TGT, --build=$(../scripts/config.guess)`
这些开关的综合效果是，glibc的构建系统将自己配置为交叉编译，使用$LFS/tools中的交叉链接器和交叉编译器。
+ `--enable-kernel=3.2`
让 `glibc` 在编译库时支持3.2及以后的Linux内核。旧内核的变通方法不被启用。
+ `--with-headers=$LFS/usr/include`
让 `glibc` 根据最近安装到 `$LFS/usr/include` 目录中的头文件来编译自己，这样它就知道内核有哪些功能，并能相应地优化自己。
+ `libc_cv_slibdir=/usr/lib`
这可以确保库被安装在/usr/lib中，而不是64位机器上默认的/lib64。

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
~~~
该命令中 `DESTDIR` 是指定将文件安装到哪个目录，如果不填写则会安装到 `/usr/bin` 。(lfs 账号并没有权限，因此就算忘记填写也不会真正覆盖原有文件。)

## 验证安装状态

~~~Shell
echo 'int main(){}' | gcc -xc -
readelf -l a.out | grep ld-linux
~~~
以上命令，是创建一个 c 代码编写的 a.out 中间文件 （使用 gcc -xc 编译生成 a.out)

返回以下信息，则代表正确。
~~~Text
[Requesting program interpreter: /lib64/ld-linux-x86-64.so.2]
~~~

删除 a.out
~~~Shell
rm -rf a.out
~~~

如果这部份存在错误，则说明之前构建的 `Binutils`，`gcc` 或 `glibc` 中的一个存在问题，需要回退到前面几个步骤，重新编译。

~~~Shell
$LFS/tools/libexec/gcc/$LFS_TGT/12.2.0/install-tools/mkheaders
~~~
以上命令，运行一次 `mkheader` 完成最后的 limits.h 头文件安装。

## 从 gcc-12.2.0 中安装 libstdc++

由于 `libstdc++` 在 `gcc-12.2.0` 目录中，因此我们需要回到 gcc-12.2.0 目录。
**注意：在这之前请提交快照或容器镜像，以便于错误回滚。**

### 准备工作
请先确保在 `/mnt/lfs/sources` 目录
先删除 gcc-12.2.0 包，避免非空目录有其他文件影响编译 libstdc++，重新解包 gcc-12.2.0.tar.xz ,建立 build 目录并进入后配置再编译
~~~Shell
rm -rf gcc-12.2.0
tar xf gcc-12.2.0.tar.xz
cd gcc-12.2.0
mkdir build
cd build
~~~

### 编译工作

**启动配置**
~~~Shell
../libstdc++-v3/configure \
 --host=$LFS_TGT \
 --build=$(../config.guess) \
 --prefix=/usr \
 --disable-multilib \
 --disable-nls \
 --disable-libstdcxx-pch \
 --with-gxx-include-dir=/tools/$LFS_TGT/include/c++/12.2.0
~~~

**启动编译并安装**
~~~Shell
make -j4
make DESTDIR=$LFS install
~~~

编译完成后需要将 libtool 文件删除，否则它会影响交叉编译
~~~Shell
rm -rf $LFS/usr/lib/lib{stdc++,stdc++fs,supc++}.la
~~~~
包括：`libstdc++.la`、`stdc++fs.la`、`supc++.la `文件

