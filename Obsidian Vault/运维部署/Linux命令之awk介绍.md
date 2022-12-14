#T200 #Linux #操作系统 #DEV #OPS #SHELL

## 概述

awk 是一种编程语言，用于在linux/unix下对文本和数据进行处理。数据可以来自标准输入、一个或多个文件，或其它命令的输出（即管道）。它支持用户自定义函数和 动态正则表达式等先进功能，是linux/unix下的一个强大编程工具。它在命令行中使用，但更多是作为脚本来使用。

awk的处理文本和数据的方式是这 样的，它逐行扫描文件，从第一行到最后一行，寻找匹配的特定模式的行，并在这些行上进行你想要的操作。如果没有指定处理动作，则把匹配的行显示到标准输出 (屏幕),即默认处理动作是print；如果没有指定模式，则所有被操作所指定的行都被处理，即默认指定模式是全部。awk分别代表其作者姓氏的第一个字母。因为它的作者是三个人，分别是Alfred Aho、Brian Kernighan、Peter Weinberger。gawk是awk的GNU版本，它提供了Bell实验室和GNU的一些扩展。

像 shell 一样，awk 也有好几种，常见的如 awk、nawk、mawk、gawk，其中

## awk 常见类型

### awk
最初在1977年完成，1985年发表了一个新版本的awk，它的功能比旧版本增强了不少,awk 能够用很短的程序对文档里的资料做修改、比较、提取、打印等处理,如果使用 C 或 Pascal 等语言编写程序完成上述的任务会十分不方便而且很花费时间，所写的程序也会很大;

### nawk
在 20 世纪 80 年代中期，对 awk语言进行了更新，并不同程度地使用一种称为 nawk(new awk) 的增强版本对其进行了替换。许多系统中仍然存在着旧的awk 解释器，但通常将其安装为 oawk (old awk) 命令，而 nawk 解释器则安装为主要的 awk 命令，也可以使用 nawk 命令。Dr. Kernighan 仍然在对 nawk 进行维护，与 gawk 一样，它也是开放源代码的，并且可以免费获得;

### mawk
mawk 是 awk 编程语言的解释器。awk语言在多媒体数据文件以及文本的检索和处理，算法的原型设计和试验都有广泛的使用。mawk带给awk新的概念，它实现了在《The AWK Programming Language》（Aho, Kernighan and Weinberger, The AWK Programming Language, Addison-Wesley Publishing, 1988.被认为是 AWK 手册。）中定义的 awk语言。mawk遵循 POSIX 1003.2 （草案 11.3）定义的 AWK 语言，包含了一些没有在AWK 手册中提到的特色，同时 mawk 提供一小部分扩展,另外据说mawk是实现最快的awk；

### gawk
是 GNU Project 的awk解释器的开放源代码实现。尽管早期的 GAWK 发行版是旧的 AWK 的替代程序，但不断地对其进行了更新，以包含 NAWK 的特性;

目前，大家都比较倾向于使用 awk 和 gawk, 本文中要介绍的awk是以GUN的gawk为例的。Ubuntu系统中的各种awk的选项设置，可以通过 sudo update-alternatives --config awk 来完成，实际上你通过手动修改软链接也能实现。Debian最小化安装的时候awk的链接是指向mawk的。

