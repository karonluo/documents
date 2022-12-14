#T200 #Python语言 #DEV #OPS #边缘计算 #嵌入式 #物联网 
# Python 语言
## 概述

Python由荷兰数学和计算机科学研究学会的吉多·范罗苏姆于1990年代初设计，作为一门叫做ABC语言的替代品。
Python提供了高效的高级数据结构，还能简单有效地面向对象编程。Python语法和动态类型，以及解释型语言的本质，使它成为多数平台上写脚本和快速开发应用的编程语言，随着版本的不断更新和语言新功能的添加，逐渐被用于独立的、大型项目的开发。Python解释器易于扩展，可以使用C语言或C++（或者其他可以通过C调用的语言）扩展新的功能和数据类型。
Python也可用于可定制化软件中的扩展程序语言。
Python丰富的标准库，提供了适用于各个主要系统平台的源码或机器码。


## 优点

**简单**：Python是一种代表[简单主义](https://baike.baidu.com/item/%E7%AE%80%E5%8D%95%E4%B8%BB%E4%B9%89/6711624?fromModule=lemma_inlink)思想的语言。阅读一个良好的Python程序就感觉像是在读英语一样。它使你能够专注于解决问题而不是去搞明白语言本身。
**易学**：Python极其容易上手，因为Python有极其简单的说明文档 [8]  。
**易读、易维护**：风格清晰划一、强制缩进
**用途广泛**
**速度较快**：Python的底层是用C语言写的，很多标准库和第三方库也都是用C写的，运行速度非常快。
**免费、开源**：Python是FLOSS（自由/开放源码软件）之一。使用者可以自由地发布这个软件的拷贝、阅读它的源代码、对它做改动、把它的一部分用于新的自由软件中。FLOSS是基于一个团体分享知识的概念。
**高层语言**：用Python语言编写程序的时候无需考虑诸如如何管理你的程序使用的内存一类的底层细节。
**可移植性**：由于它的开源本质，Python已经被移植在许多平台上（经过改动使它能够工作在不同平台上）。这些平台包括Linux、Windows、FreeBSD、Macintosh、Solaris、OS/2、Amiga、AROS、AS/400、BeOS、OS/390、z/OS、Palm OS、QNX、VMS、Psion、Acom RISC OS、VxWorks、PlayStation、Sharp Zaurus、Windows CE、PocketPC、Symbian以及Google基于linux开发的android平台。
**解释性**：一个用编译性语言比如C或C++写的程序可以从源文件（即C或C++语言）转换到一个你的计算机使用的语言（二进制代码，即0和1）。这个过程通过编译器和不同的标记、选项完成。

运行程序的时候，连接/转载器软件把你的程序从硬盘复制到内存中并且运行。而Python语言写的程序不需要编译成二进制代码。你可以直接从源代码运行程序。

在计算机内部，Python解释器把源代码转换成称为字节码的中间形式，然后再把它翻译成计算机使用的机器语言并运行。这使得使用Python更加简单。也使得Python程序更加易于移植。

**面向对象**：Python既支持面向过程的编程也支持面向对象的编程。在“面向过程”的语言中，程序是由过程或仅仅是可重用代码的函数构建起来的。在“面向对象”的语言中，程序是由数据和功能组合而成的对象构建起来的。

Python是完全[面向对象](https://baike.baidu.com/item/%E9%9D%A2%E5%90%91%E5%AF%B9%E8%B1%A1/2262089?fromModule=lemma_inlink)的语言。[函数](https://baike.baidu.com/item/%E5%87%BD%E6%95%B0/18686609?fromModule=lemma_inlink)、模块、数字、[字符串](https://baike.baidu.com/item/%E5%AD%97%E7%AC%A6%E4%B8%B2/1017763?fromModule=lemma_inlink)都是对象。并且完全支持继承、重载、派生、多继承，有益于增强源代码的复用性。Python支持重载运算符和动态类型。相对于[Lisp](https://baike.baidu.com/item/Lisp/22083?fromModule=lemma_inlink)这种传统的函数式编程语言，Python对函数式设计只提供了有限的支持。有两个标准库（functools，itertools）提供了Haskell和Standard ML中久经考验的函数式程序设计工具。

**可扩展性、可扩充性**：如果需要一段关键代码运行得更快或者希望某些算法不公开，可以部分程序用C或C++编写，然后在Python程序中使用它们。

Python本身被设计为可扩充的。并非所有的特性和功能都集成到语言核心。Python提供了丰富的[API](https://baike.baidu.com/item/API/10154?fromModule=lemma_inlink)和工具，以便程序员能够轻松地使用[C语言](https://baike.baidu.com/item/C%E8%AF%AD%E8%A8%80/105958?fromModule=lemma_inlink)、[C++](https://baike.baidu.com/item/C%2B%2B/99272?fromModule=lemma_inlink)、Cython来编写扩充模块。Python编译器本身也可以被集成到其它需要脚本语言的程序内。因此，很多人还把Python作为一种“胶水语言”（glue language）使用。使用Python将其他语言编写的程序进行集成和封装。在Google内部的很多项目，例如Google Engine使用C++编写性能要求极高的部分，然后用Python或Java/Go调用相应的模块。《Python技术手册》的作者马特利（Alex Martelli）说：“这很难讲，不过，2004年，Python已在Google内部使用，Google 召募许多 Python 高手，但在这之前就已决定使用Python，他们的目的是 Python where we can，C++ where we must，在操控硬件的场合使用[C++](https://baike.baidu.com/item/C%2B%2B/99272?fromModule=lemma_inlink)，在快速开发时候使用Python。”

**可嵌入性**：可以把Python嵌入[C](https://baike.baidu.com/item/C/7252092?fromModule=lemma_inlink)/[C++](https://baike.baidu.com/item/C%2B%2B/99272?fromModule=lemma_inlink)程序，从而向程序用户提供脚本功能。
**丰富的库**：Python标准库确实很庞大。它可以帮助处理各种工作，包括正则表达式、文档生成、单元测试、线程、数据库、[网页浏览器](https://baike.baidu.com/item/%E7%BD%91%E9%A1%B5%E6%B5%8F%E8%A7%88%E5%99%A8/8309940?fromModule=lemma_inlink)、CGI、FTP、[电子邮件](https://baike.baidu.com/item/%E7%94%B5%E5%AD%90%E9%82%AE%E4%BB%B6/111106?fromModule=lemma_inlink)、XML、XML-RPC、HTML、WAV文件、密码系统、GUI（图形用户界面）、Tk和其他与系统有关的操作。这被称作Python的“功能齐全”理念。除了标准库以外，还有许多其他高质量的库，如wxPython、Twisted和Python图像库等等。
**规范的代码**：Python采用强制缩进的方式使得代码具有较好可读性。而Python语言写的程序不需要编译成二进制代码。Python的作者设计限制性很强的语法，使得不好的编程习惯（例如if语句的下一行不向右缩进）都不能通过编译。其中很重要的一项就是Python的[缩进](https://baike.baidu.com/item/%E7%BC%A9%E8%BF%9B/7337492?fromModule=lemma_inlink)规则。一个和其他大多数语言（如C）的区别就是，一个模块的界限，完全是由每行的首字符在这一行的位置来决定（而C语言是用一对大括号“{}”（不含引号）来明确的定出模块的边界，与字符的位置毫无关系）。通过强制程序员们缩进（包括if，for和函数定义等所有需要使用模块的地方），Python确实使得程序更加清晰和美观。

**高级动态编程**：虽然Python可能被粗略地分类为“脚本语言”（script language），但实际上一些大规模软件开发计划例如Zope、Mnet及[BitTorrent](https://baike.baidu.com/item/BitTorrent/142795?fromModule=lemma_inlink)，[Google](https://baike.baidu.com/item/Google/86964?fromModule=lemma_inlink)也广泛地使用它。Python的支持者较喜欢称它为一种高级动态编程语言，原因是“脚本语言”泛指仅作简单程序设计任务的语言，如shellscript、VBScript等只能处理简单任务的编程语言，并不能与Python相提并论。

**做科学计算优点多**：说起科学计算，首先会被提到的可能是MATLAB。除了MATLAB的一些专业性很强的工具箱还无法被替代之外，[MATLAB](https://baike.baidu.com/item/MATLAB/263035?fromModule=lemma_inlink)的大部分常用功能都可以在Python世界中找到相应的扩展库。和MATLAB相比，用Python做科学计算有如下优点：
●首先，[MATLAB](https://baike.baidu.com/item/MATLAB/263035?fromModule=lemma_inlink)是一款商用软件，并且价格不菲。而Python完全免费，众多开源的[科学](https://baike.baidu.com/item/%E7%A7%91%E5%AD%A6/10406?fromModule=lemma_inlink)计算库都提供了Python的调用接口。用户可以在任何计算机上免费安装Python及其绝大多数扩展库。
●其次，与[MATLAB](https://baike.baidu.com/item/MATLAB/263035?fromModule=lemma_inlink)相比，Python是一门更易学、更严谨的程序设计语言。它能让用户编写出更易读、易维护的代码。
●最后，[MATLAB](https://baike.baidu.com/item/MATLAB/263035?fromModule=lemma_inlink)主要专注于工程和科学计算。然而即使在计算领域，也经常会遇到文件管理、界面设计、网络通信等各种需求。而Python有着丰富的扩展库，可以轻易完成各种高级任务，开发者可以用Python实现完整应用程序所需的各种功能。

## 缺点

**单行语句和命令行输出问题**：很多时候不能将程序连写成一行，如import sys；for i in sys.path：print i。而perl和awk就无此限制，可以较为方便的在shell下完成简单程序，不需要如Python一样，必须将程序写入一个.py文件。

**给初学者带来困惑**：独特的语法，这也许不应该被称为局限，但是它用缩进来区分语句关系的方式还是给很多初学者带来了困惑。即便是很有经验的Python程序员，也可能陷入陷阱当中。

**运行速度慢**：这里是指与C和C++相比。Python开发人员尽量避开不成熟或者不重要的优化。一些针对非重要部位的加快运行速度的补丁通常不会被合并到Python内。所以很多人认为Python很慢。不过，根据二八定律，大多数程序对速度要求不高。在某些对运行速度要求很高的情况，Python设计师倾向于使用JIT技术，或者用使用[C](https://baike.baidu.com/item/C/7252092?fromModule=lemma_inlink)/[C++](https://baike.baidu.com/item/C%2B%2B/99272?fromModule=lemma_inlink)语言改写这部分程序。可用的JIT技术是[PyPy](https://baike.baidu.com/item/PyPy/9780733?fromModule=lemma_inlink)。

**和其他语言区别**

**对于一个特定的问题，只要有一种最好的方法来解决**

这在由Tim Peters写的Python格言（称为The Zen of Python）里面表述为：There should be one-and preferably only one-obvious way to do it。这正好和Perl语言（另一种功能类似的高级动态语言）的中心思想TMTOWTDI（There's More Than One Way To Do It）完全相反。

Python的设计哲学是“优雅”、“明确”、“简单”。因此，[Perl语言](https://baike.baidu.com/item/Perl%E8%AF%AD%E8%A8%80/1346108?fromModule=lemma_inlink)中“总是有多种方法来做同一件事”的理念在Python开发者中通常是难以忍受的。Python开发者的哲学是“用一种方法，最好是只有一种方法来做一件事”。在设计Python语言时，如果面临多种选择，Python开发者一般会拒绝花俏的语法，而选择明确的没有或者很少有歧义的语法。由于这种设计观念的差异，Python源代码通常被认为比Perl具备更好的可读性，并且能够支撑大规模的软件开发。这些准则被称为Python格言。在Python解释器内运行import this可以获得完整的列表。

**更高级的Virtual Machine**
Python在执行时，首先会将.py文件中的[源代码](https://baike.baidu.com/item/%E6%BA%90%E4%BB%A3%E7%A0%81/3969?fromModule=lemma_inlink)编译成Python的byte code（字节码），然后再由Python Virtual Machine（Python虚拟机）来执行这些编译好的byte code。这种机制的基本思想跟Java，.NET是一致的。然而，Python Virtual Machine与Java或.NET的Virtual Machine不同的是，Python的Virtual Machine是一种更高级的Virtual Machine。这里的高级并不是通常意义上的高级，不是说Python的Virtual Machine比Java或.NET的功能更强大，而是说和Java 或.NET相比，Python的Virtual Machine距离真实机器的距离更远。或者可以这么说，Python的Virtual Machine是一种抽象层次更高的Virtual Machine。基于C的Python编译出的[字节码](https://baike.baidu.com/item/%E5%AD%97%E8%8A%82%E7%A0%81?fromModule=lemma_inlink)文件，通常是.pyc格式。除此之外，Python还可以以交互模式运行，比如主流操作系统Unix/[Linux](https://baike.baidu.com/item/Linux/27050?fromModule=lemma_inlink)、Mac、Windows都可以直接在命令模式下直接运行Python交互环境。直接下达操作指令即可实现交互操作。


## 相关技术和知识

### [[Python语言规范]]

### [[Python语言性能优化]]

