# TensorRT INT8 量化原理与实现
## 一、模型量化是什么？

模型量化是由模型、量化两个词组成。我们要准确理解模型量化，要看这两个词分别是什么意思。
在计算机视觉、深度学习的语境下，模型特指卷积神经网络，用于提取图像/视频视觉特征。
量化是指将信号的连续取值近似为有限多个离散值的过程，可理解成一种信息压缩的方法。在计算机系统上考虑这个概念的话，量化有若干相似的术语，低精度可能是最通用的概念。常规精度一般使用 FP32（32位浮点，单精度）存储模型权重；低精度则表示 FP16（半精度浮点），INT8（8位的定点整数）等等数值格式。目前，低精度往往指代INT8，因此也有人称量化为“定点化”，但是严格来讲所表示的范围是缩小的。定点化特指scale为2的幂次的线性量化，是一种更加实用的量化方法。
简而言之，我们常说的模型量化就是将浮点存储（运算）转换为整型存储（运算）的一种模型压缩技术。举个例子，即原来表示一个权重或偏置需要使用FP32表示，使用了INT8量化后只需要使用一个INT8来表示就可以了。
注：以下主要基于INT8量化介绍。

## 二、为什么要做模型量化？
现有的深度学习框架，比如：TensorFlow，Pytorch，Caffe， MixNet等，在训练深度神经网络时，往往都会使用FP32的数据精度来表示权值、偏置、激活值等。在深度学习模型性能提高的同时，计算也越来越复杂，计算开销和内存需求逐渐增加。仅 8 层的 AlexNet需要0.61 亿个网络参数和 7.29 亿次浮点型计算，花费约 233MB 内存。随后的 VGG-16的网络参数达到 1.38 亿，浮点型计算次数为 156 亿，需要约 553MB 内存。为了克服深层网络的梯度消失问题。He 提出了 ResNet网络，首次在ILSVRC 比赛中实现了低于 5%的 top-5 分类错误，偏浅的 ResNet-50 网络参数就达到 0.25 亿，浮点型计算次数高达 41.2亿，内存花费约 102MB。

![[Pasted image 20221013163353.png]]
表1 不同模型的模型大小及浮点运算次数

庞大的网络参数意味着更大的内存存储，而增长的浮点型计算次数意味着训练成本和计算时间的增长，这极大地限制了在资源受限设备，例如智能手机、智能手环等上的部署。如表2所示，深度模型在 Samsung Galaxy S6 的推理时间远超 Titan X 桌面级显卡，实时性较差，无法满足实际应用的需要。

![[Pasted image 20221013163241.png]]
表2 不同模型在不同设备上的推理时间（单位：ms）

## 三、模型量化的目标是什么？

![[Pasted image 20221013163605.png]]
1.  **更小的模型尺寸。**
2.  **更低的运算功耗。**
3.  **更低的运存占用。**
4.  **更快的计算速度。**
5.  **持平的推理精度。**

## 四、模型量化的必要条件
量化是否一定能加速计算？回答是否定的，许多量化算法都无法带来实质性加速。
引入一个概念：**理论计算峰值**。
在高性能计算领域，这概念一般被定义为：**单位时钟周期内能完成的计算个数乘上芯片频率**。
什么样的量化方法可以带来潜在、可落地的速度提升呢？
我们总结需要满足两个条件：
1. 量化数值的计算在部署硬件上的峰值性能更高 。
2. 量化算法引入的额外计算（overhead）少 。
要准确理解上述条件，需要有一定的高性能计算基础知识，限于篇幅就不展开讨论了。现直接给出如下结论，已知提速概率较大的量化方法主要有如下三类:
1. **二值化**，其可以用简单的位运算来同时计算大量的数。对比从nvdia gpu到x86平台，1bit计算分别有5到128倍的理论性能提升。且其只会引入一个额外的量化操作，该操作可以享受到SIMD（单指令多数据流）的加速收益。
2. **线性量化**，又可细分为非对称，对称几种。在nvdia gpu，x86和arm平台上，均支持8bit的计算，效率提升从1倍到16倍不等，其中tensor core甚至支持4bit计算，这也是非常有潜力的方向。由于线性量化引入的额外量化/反量化计算都是标准的向量操作，也可以使用SIMD进行加速，带来的额外计算耗时不大。
3. **对数量化**，一个比较特殊的量化方法。可以想象一下，两个同底的幂指数进行相乘，那么等价于其指数相加，降低了计算强度。同时加法也被转变为索引计算。但没有看到有在三大平台上实现对数量化的加速库，可能其实现的加速效果不明显。只有一些专用芯片上使用了对数量化。

## 五、模型量化的分类
### 5.1线性量化和非线性量化
根据映射函数是否是线性可以分为两类，即线性量化和非线性量化，本文主要研究的是线性量化技术。
### 5.2 逐层量化、逐组量化和逐通道量化
根据量化的粒度（共享量化参数的范围）可以分为逐层量化、逐组量化和逐通道量化。
- 逐层量化，以一个层为单位，整个layer的权重共用一组缩放因子S和偏移量Z；
- 逐组量化，以组为单位，每个group使用一组S和Z；
- 逐通道量，化则以通道为单位，每个channel单独使用一组S和Z；
当 group=1 时，逐组量化与逐层量化等价；当 group=num_filters （即dw卷积）时，逐组量化逐通道量化等价。
### 5.3 N比特量化 
根据存储一个权重元素所需的位数，可以将其分为8bit量化、4bit量化、2bit量化和1bit量化等。

### 5.4 权重量化和权重激活量化
#### 5.4.1 权重与激活的概念
我们来看一个简单的深度学习网络，如图2所示
![[Pasted image 20221013173858.png]]图2 深度学习网络维度示意图
其中滤波器就是权重，而输入和输出数据则是分别是上一层和当前层的激活值，假设输入数据为[3,224,224]，滤波器为[2,3,3,3]，使用如下公式可以计算得到输出数据为[2,222,222]
![[Pasted image 20221013174008.png]]
因此，权重有2 x 3 x 3 x 3= 54 个（不含偏置），上一层的激活值有3 x 224 x 224 = 150528 个，下一层的激活值有2 x 222 x 222=98568个，显然激活值的数量远大于权重。
#### 5.4.2 权重量化和权重激活量化
根据需要量化的参数可以分类两类：权重量化和权重激活量化。
1. **权重量化**，即仅仅需要对网络中的权重执行量化操作。由于网络的权重一般都保存下来了，因而我们可以提前根据权重获得相应的量化参数S和Z，而不需要额外的校准数据集。一般来说，推理过程中，权重值的数量远小于激活值，仅仅对权重执行量化的方法能带来的压缩力度和加速效果都一般。
2. **权重激活量化**，即不仅对网络中的权重进行量化，还对激活值进行量化。由于激活层的范围通常不容易提前获得，因而需要在网络推理的过程中进行计算或者根据模型进行大致的预测。
#### 5.4.3 激活量化方式
根据激活值的量化方式，可以分为在线量化和离线量化。

- **在线量化**，即指激活值的S和Z在实际推断过程中根据实际的激活值动态计算；
- **离线量化**，即指提前确定好激活值的S和Z，需要小批量的一些校准数据集支持。
由于不需要动态计算量化参数，通常离线量化的推断速度更快些。
通常使用以下的三种方法来确定相关的量化参数。
- **指数平滑法**，即将校准数据集送入模型，收集每个量化层的输出特征图，计算每个batch的S和Z值，并通过指数平滑法来更新S和Z值。
- **直方图截断法**，即在计算量化参数Z和S的过程中，由于有的特征图会出现偏离较远的奇异值，导致max非常大，所以可以通过直方图截取的形式，比如抛弃最大的前1%数据，以前1%分界点的数值作为max计算量化参数。
- **KL散度校准法**，即通过计算KL散度（也称为相对熵，用以描述两个分布之间的差异）来评估量化前后的两个分布之间存在的差异，搜索并选取KL散度最小的量化参数Z和S作为最终的结果。TensorRT中就采用了这种方法。

### 5.5 训练时量化和训练后量化
 **训练后量化（Post-Training Quantization,PTQ）**，PTQ不需要再训练，因此是一种轻量级的量化方法。在大多数情况下，PTQ足以实现接近FP32性能的INT8量化。然而，它也有局限性，特别是针对激活更低位的量化，如4bit、2bit。这时就有了训练时量化的用武之地。
 
**训练时量化也叫量化感知训练（Quantization-Aware-Training,QAT）**，它可以获得高精度的低位量化，但是缺点也很明显，就是需要修改训练代码，并且反向传播过程对梯度的量化误差很大，也容易导致不能收敛。

本篇内容主要介绍训练后量化（PTQ）

## 六、量化的数学基础

### 6.1 定点数和浮点数
量化过程可以分为两部分：将模型从 FP32 转换为INT8，以及使用INT8 进行推理。本节说明这两部分背后的算术原理。如果不了解基础算术原理，在考虑量化细节时通常会感到困惑。
从事计算机科学的人很少了解算术运算的执行方式。由于量化桥接了固定点和浮点，在接触相关研究和解决方案之前，有必要先了解它们的基础知识。
定点和浮点都是数值的表示方式，它们区别在于，将整数部分和小数部分分开的点，位于哪里。定点保留特定位数整数和小数，而浮点保留特定位数的有效数字和指数。

![[Pasted image 20221014104225.png]]
表3 定点和浮点的格式与示例

在指令集的内置数据类型中，定点是整数，浮点是二进制格式。一般来说，指令集层面的定点是连续的，因为它是整数，且两个邻近的可表示数字的间隙是 1 。而浮点代表实数，其数值间隙由指数确定，因而具有非常宽的值域。同时也可以知道浮点的数值间隙是不均匀的，在相同的指数范围内，可表示数值数量也相同，且值越接近零就越准确。例如，\[1,2)中浮点值的数量与\[0.5,1)、\[2,4)、\[4,8) 等相同。另外，我们也可以得知定点数数值与想要表示的真值是一致的，而浮点数数值与想要表示的真值是有偏差的。

![[Pasted image 20221014104537.png]]

表4 FP32和INT32的数值范围及可取值数量
![[Pasted image 20221014105447.png]]
图3 浮点数与定点数对照关系示意图

 举个例子，假设每个指数范围内可表示数值量为2，例如\[2^0,2^1)范围内的数值转换成浮点数只能表示成{1, 1.5}。
 
 ![[Pasted image 20221014105654.png]]
表5 浮点数数值间隙不同的示例

### 6.2 线性量化（线性映射）
#### 6.2.1 量化
TensorRT 使用的就是线性量化，它可以用以下数学表达式来表示：

![[Pasted image 20221014105913.png]]

其中，
X表示原始的FP32数值；
Z表示映射的零点Zero Point；
S表示缩放因子Scale；
![[Pasted image 20221014110043.png]]表示的是近似取整的数学函数，可以是四舍五入、向上取整、向下取整等；
![[Pasted image 20221014110003.png]]表示的是量化后的一个整数值。

clip函数如下：

![[Pasted image 20221014110242.png]]

根据参数 Z 是否为零可以将线性量化分为两类—即对称量化和非对称量化，TensorRT 使用的时对称量化，即Z=0。
![[Pasted image 20221014110348.png]]
图4 对称带符号量化、对称无符号量化和非对称量化

#### 6.2.2 反量化

根据量化的公式不难推导出，反量化公式如下：
![[Pasted image 20221014110459.png]]
当Z=0时，![[Pasted image 20221014110527.png]]。

可以发现**当S取大时，可以扩大量化域，但同时，单个INT8数值可表示的FP32范围也变广了，因此INT8数值与FP32数值的误差（量化误差）会增大；而当S取小时，量化误差虽然减小了，但是量化域也缩小了，被舍弃的参数会增多。**
举个例子，假设Z=0,使用向下取整。
![[Pasted image 20221014110701.png]]
表6 不同缩放尺度的影响
## 七、TensorRT INT8 量化原理
### 7.1 TensorRT是什么
