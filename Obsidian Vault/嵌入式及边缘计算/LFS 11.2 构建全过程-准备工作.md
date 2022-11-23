#T800 #嵌入式 #物联网 #边缘计算 #Linux #LFS #操作系统 

在 `LFS-BOOK-11.2` 中，提到了需要进行一系列的插件的版本准备
需要进行检查的插件及其版本如下： ^19ba7b
~~~Text
• Bash-3.2 (/bin/sh should be a symbolic or hard link to bash)
• Binutils-2.13.1 (Versions greater than 2.39 are not recommended as they have not been tested)
• Bison-2.7 (/usr/bin/yacc should be a link to bison or small script that executes bison)
• Coreutils-6.9
• Diffutils-2.8.1
• Findutils-4.2.31
• Gawk-4.0.1 (/usr/bin/awk should be a link to gawk)
• GCC-4.8 including the C++ compiler, g++ (Versions greater than 12.2.0 are not recommended as they have not
been tested). C and C++ standard libraries (with headers) must also be present so the C++ compiler can build
hosted programs
• Grep-2.5.1a
• Gzip-1.3.12
• Linux Kernel-3.2
The reason for the kernel version requirement is that we specify that version when building glibc in Chapter 5 and
Chapter 8, at the recommendation of the developers. It is also required by udev.
If the host kernel is earlier than 3.2 you will need to replace the kernel with a more up to date version. There
are two ways you can go about this. First, see if your Linux vendor provides a 3.2 or later kernel package. If so,
you may wish to install it. If your vendor doesn't offer an acceptable kernel package, or you would prefer not to
install it, you can compile a kernel yourself. Instructions for compiling the kernel and configuring the boot loader
(assuming the host uses GRUB) are located in Chapter 10.
• M4-1.4.10
• Make-4.0
• Patch-2.5.4
• Perl-5.8.8
• Python-3.4
• Sed-4.1.5
• Tar-1.22
• Texinfo-4.7
• Xz-5.0.0
~~~
在LFS-BOOK中，它贴心地准备了一个用于版本检查的bash脚本
我们打开终端，直接在终端中粘贴下面的脚本，按下回车即可自动执行
~~~Shell
cat > version-check.sh << "EOF"
#!/bin/bash
# Simple script to list version numbers of critical development tools
export LC_ALL=C
bash --version | head -n1 | cut -d" " -f2-4
MYSH=$(readlink -f /bin/sh)
echo "/bin/sh -> $MYSH"
echo $MYSH | grep -q bash || echo "ERROR: /bin/sh does not point to bash"
unset MYSH
echo -n "Binutils: "; ld --version | head -n1 | cut -d" " -f3-
bison --version | head -n1
if [ -h /usr/bin/yacc ]; then
 echo "/usr/bin/yacc -> `readlink -f /usr/bin/yacc`";
elif [ -x /usr/bin/yacc ]; then
 echo yacc is `/usr/bin/yacc --version | head -n1`
else
 echo "yacc not found"
fi
echo -n "Coreutils: "; chown --version | head -n1 | cut -d")" -f2
diff --version | head -n1
find --version | head -n1
gawk --version | head -n1
if [ -h /usr/bin/awk ]; then
 echo "/usr/bin/awk -> `readlink -f /usr/bin/awk`";
elif [ -x /usr/bin/awk ]; then
 echo awk is `/usr/bin/awk --version | head -n1`
else
 echo "awk not found"
fi
gcc --version | head -n1
g++ --version | head -n1
grep --version | head -n1
gzip --version | head -n1
cat /proc/version
m4 --version | head -n1
make --version | head -n1
patch --version | head -n1
echo Perl `perl -V:version`
python3 --version
sed --version | head -n1
tar --version | head -n1
makeinfo --version | head -n1 # texinfo version
xz --version | head -n1
echo 'int main(){}' > dummy.c && g++ -o dummy dummy.c
if [ -x dummy ]
 then echo "g++ compilation OK";
 else echo "g++ compilation failed"; fi
rm -f dummy.c dummy
EOF
bash version-check.sh
~~~
运行后，我的输出如下：
~~~Shell
alphainf@ubuntu:~$ bash version-check.sh
bash, version 4.3.48(1)-release
/bin/sh -> /bin/dash
ERROR: /bin/sh does not point to bash
Binutils: (GNU Binutils for Ubuntu) 2.26.1
version-check.sh: line 10: bison: command not found
yacc not found
Coreutils:  8.25
diff (GNU diffutils) 3.3
find (GNU findutils) 4.7.0-git
version-check.sh: line 21: gawk: command not found
/usr/bin/awk -> /usr/bin/mawk
gcc (Ubuntu 5.4.0-6ubuntu1~16.04.12) 5.4.0 20160609
g++ (Ubuntu 5.4.0-6ubuntu1~16.04.12) 5.4.0 20160609
grep (GNU grep) 2.25
gzip 1.6
Linux version 4.15.0-112-generic (buildd@lcy01-amd64-021) (gcc version 5.4.0 20160609 (Ubuntu 5.4.0-6ubuntu1~16.04.12)) #113~16.04.1-Ubuntu SMP Fri Jul 10 04:37:08 UTC 2020
version-check.sh: line 34: m4: command not found
GNU Make 4.1
GNU patch 2.7.5
Perl version='5.22.1';
Python 3.5.2
sed (GNU sed) 4.2.2
tar (GNU tar) 1.28
version-check.sh: line 41: makeinfo: command not found
xz (XZ Utils) 5.1.0alpha
g++ compilation OK
~~~
经过对比，我们发现存在以下的问题
 1、 `/bin/sh -> /bin/dash     ERROR: /bin/sh does not point to bash`
shell脚本未指向bash而是指向dash
`sudo ln -sf bash /bin/sh`

2、bison: command not found（bison是属于 GNU 项目的一个语法分析器生成器）
`sudo apt-get install bison`
注意：在安装bison期间，m4会自动被完成安装
3、gawk not found （linux下查找替换文本工具）
`sudo apt-get install gawk`
4、makeinfo:command not found
`sudo apt-get install texinfo`
完成上述修改后，我们运行一下指令进行检查
`bash version-check.sh`
输出如下：
~~~Shell
bash, version 4.3.48(1)-release
/bin/sh -> /bin/bash
Binutils: (GNU Binutils for Ubuntu) 2.26.1
bison (GNU Bison) 3.0.4
/usr/bin/yacc -> /usr/bin/bison.yacc
Coreutils:  8.25
diff (GNU diffutils) 3.3
find (GNU findutils) 4.7.0-git
GNU Awk 4.1.3, API: 1.1 (GNU MPFR 3.1.4, GNU MP 6.1.0)
/usr/bin/awk -> /usr/bin/gawk
gcc (Ubuntu 5.4.0-6ubuntu1~16.04.12) 5.4.0 20160609
g++ (Ubuntu 5.4.0-6ubuntu1~16.04.12) 5.4.0 20160609
grep (GNU grep) 2.25
gzip 1.6
Linux version 4.15.0-112-generic (buildd@lcy01-amd64-021) (gcc version 5.4.0 20160609 (Ubuntu 5.4.0-6ubuntu1~16.04.12)) #113~16.04.1-Ubuntu SMP Fri Jul 10 04:37:08 UTC 2020
m4 (GNU M4) 1.4.17
GNU Make 4.1
GNU patch 2.7.5
Perl version='5.22.1';
Python 3.5.2
sed (GNU sed) 4.2.2
tar (GNU tar) 1.28
texi2any (GNU texinfo) 6.1
xz (XZ Utils) 5.1.0alpha
g++ compilation OK
~~~
经确认，软件版本无误