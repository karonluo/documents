#T800 #嵌入式 #物联网 #边缘计算 #Linux #LFS #操作系统 

## 创建LFS目录

~~~Shell
export LFS=/mnt/lfs
mkdir -pv $LFS/{etc,var} $LFS/usr/{bin,lib,sbin}
for i in bin lib sbin; do
 ln -sv usr/$i $LFS/$i
done
case $(uname -m) in
 x86_64) mkdir -pv $LFS/lib64 ;;
esac
mkdir -pv $LFS/tools
~~~

> 注意：以下步骤均在 `/mnt/lfs` 目录中进行操作

## 创建用户组

+ **我们需要添加一个普通用户组到用户中，以免特权指令影响到新操作系统的安全**
~~~Shell
groupadd lfs
useradd -s /bin/bash -g lfs -m -k /dev/null lfs
passwd lfs
~~~

+ **授予 `lfs` 访问 `/mnt/lfs` 权限**
~~~Shell
chown -v lfs $LFS/{usr{,/*},lib,var,etc,bin,sbin,tools,sources}
case $(uname -m) in
 x86_64) chown -v lfs $LFS/lib64 ;;
esac
~~~
如下是执行结果，如有异常，可手动进行配置。
~~~Shell
changed ownership of '/mnt/lfs/usr' from root to lfs
changed ownership of '/mnt/lfs/usr/bin' from root to lfs
changed ownership of '/mnt/lfs/usr/lib' from root to lfs
changed ownership of '/mnt/lfs/usr/sbin' from root to lfs
ownership of '/mnt/lfs/lib' retained as lfs
changed ownership of '/mnt/lfs/var' from root to lfs
changed ownership of '/mnt/lfs/etc' from root to lfs
ownership of '/mnt/lfs/bin' retained as lfs
ownership of '/mnt/lfs/sbin' retained as lfs
changed ownership of '/mnt/lfs/tools' from root to lfs
changed ownership of '/mnt/lfs/sources' from root to lfs
root@cfb58bfe2343:/mnt/lfs# case $(uname -m) in
>  x86_64) chown -v lfs $LFS/lib64 ;;
> esac
changed ownership of '/mnt/lfs/lib64' from root to lfs
~~~

## 切换用户 lfs 进行后续工作

~~~Shell
su lfs
~~~

+ **对新账户设置环境**
设置 `.bash_profile`
~~~Shell
cat > ~/.bash_profile << "EOF"
exec env -i HOME=$HOME TERM=$TERM PS1='\u:\w\$ ' /bin/bash
EOF
~~~
设置 `.bashrc`
~~~Shell
cat > ~/.bashrc << "EOF"
set +h
umask 022
LFS=/mnt/lfs
LC_ALL=POSIX
LFS_TGT=$(uname -m)-lfs-linux-gnu
PATH=/usr/bin
if [ ! -L /bin ]; then PATH=/bin:$PATH; fi
PATH=$LFS/tools/bin:$PATH
CONFIG_SITE=$LFS/usr/share/config.site
export LFS LC_ALL LFS_TGT PATH CONFIG_SITE
EOF
~~~
启用环境
~~~Shell
source ~/.bash_profile
source ~/.bashrc
~~~


