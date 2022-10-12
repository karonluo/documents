# Redis 优化相关知识
---
- [# Redis 优化相关知识](#-redis-优化相关知识)
- [前言](#前言)
- [一、数据存储优化](#一数据存储优化)
  - [1.控制key的长度](#1控制key的长度)
  - [2.控制元素的大小](#2控制元素的大小)
  - [3.存储合适数据类型](#3存储合适数据类型)
  - [4.设置过期key](#4设置过期key)
  - [5.冷热分离](#5冷热分离)
  - [6.数据压缩](#6数据压缩)
- [二、内存淘汰优化](#二内存淘汰优化)
- [三、过期策略优化](#三过期策略优化)
- [四、持久化优化](#四持久化优化)
- [五、部署优化](#五部署优化)
  - [1.物理机部署](#1物理机部署)
  - [2.内网部署](#2内网部署)
- [六、集群优化](#六集群优化)
  - [1.只使用db0](#1只使用db0)
  - [2.架构优化](#2架构优化)
  - [3.实例内存优化](#3实例内存优化)
- [七、其他优化](#七其他优化)
  - [1.lazy-free机制](#1lazy-free机制)
  - [2.批量命令](#2批量命令)
  - [3.避免复杂度过高命令](#3避免复杂度过高命令)
  - [4.注意容器类型数据](#4注意容器类型数据)

---
## 前言

<font size=4>Redis都是基于内存的操作，CPU不是redis的性能瓶颈，则服务器的内存利用率和网络IO就是redis的性能瓶颈，redis优化主要是从这2个维度做优化。
</font>

## 一、数据存储优化

### 1.控制key的长度 
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>在保证key在简单、清晰的前提下，尽可能把key定义得短一些来控制key的长度，如 uerOreder:0001 缩写为uo:0001。<br/>
建议：做好key规范管理的文档，定义好每个key缩写的含义。
如按以上规范来做当redis中有大量的key时也会节约redis大量的内存空间，使其性能更高。
。</font>

### 2.控制元素的大小
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>在决定使用redis建议强制规定元素的大小，推荐规范如下：
1. String类型数据的值控制在10K以下；
2. List/Hash/Set/ZSet数据类型的元素要控制在1W以内。
以上的是杜绝bigkey有效措施，在bigkey对redis的性能影响是最为致命的。
</font>

### 3.存储合适数据类型
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>除非业务的强要求，建议选择合适的数据类型来优化内存，具体如下：<br/>
1. String、Set：尽可能存储 int 类型数据；<br/>
2. Hash、ZSet：存储的元素数量控制在转换阈值之下，以压缩列表存储，节约内存。
</font>

### 4.设置过期key
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>强制要求所有的key必须设置过期时间，以优化redis内存。</font>

### 5.冷热分离
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>热key需要单独存放并分配合理的资源，防止大流量下直接冲垮整个缓存系统。</font>

### 6.数据压缩
<font size=4>可以采用 snappy、gzip 等压缩算法来先将数据压缩后再存入缓存中，来节约redis的内存空间， 但这种方法会使客户端在读取时还需要解压缩，在这期间会消耗更多CPU资源，你需要根据实际情况进行权衡。 <br/>
建议，只是在redis匮乏时的一种方案。</font>

## 二、内存淘汰优化
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>杜绝使用默认的内存淘汰策略，避免在业务扩展下Redis的内存持续膨胀，需要根据你的业务设置对应内存淘汰策略。</font>

## 三、过期策略优化
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>由于redis采用的是定期删除+懒加载删除策略，且这个过程在redis 6.0之前是在主线程上执行的，
建议所有key的过期时间用随机数打散，杜绝大批量的数据同时过期，拉胯redis的性能和造成缓存雪崩。</font>

## 四、持久化优化
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>除非业务的强要求，否则不要开启AOF，避免写磁盘拖垮redis的性能。<br />
如果有业务要求开启AOF，建议配置为 appendfsync everysec 把数据持久化的刷盘操作，放到后台线程中去执行，尽量降低 Redis 写磁盘对性能的影响。</font>

## 五、部署优化
### 1.物理机部署
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>Redis在做数据持久化时，采用创建子进程的方式进行（ 会调用操作系统的 fork 系统 ），而 虚拟机环境执行 fork 的耗时，要比物理机慢得多。</font>

### 2.内网部署
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>Redis集群有做副本数据同步，为了解决网络问题建议整个集群都部署在同一个局域网中。</font>

## 六、集群优化
### 1.只使用db0
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>具体原因如下：<br />
1.在一个连接上操作多个db数据时，每次都会先执行查询，有额外开销；<br />
2.建议不同业务数据尽量存储到不同的实例中而不是存放在不同db中；<br />
3.redis cluster只支持db0。</font>

### 2.架构优化
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>读写分离能最大限度提高redis的性能，其中主库负责数据写入，从库负责数据读取；<br />
分片集群是解决超大量数据导致性能瓶颈方案，如rediscluster。<br />
以上是在大流量下提高redis性能在架构上的优化。</font>

### 3.实例内存优化
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>Redis集群有做副本数据同步，如果redis的内存过大，那么在做持久化或数据同步时，
磁盘IO和网络IO会拉垮redis性能，导致卡顿。<br/>
建议redis每一个实例都不要超过8G，要从业务上解耦，不同的业务用不同的redis实例，杜绝所有业务都使用同一个redis实例。</font>

## 七、其他优化

### 1.lazy-free机制
<font size=4 style='font-weight:bold'>强制规范</font>

<font size=4>在redis4.0+中支持，开启lazy-free机制后，由主线程删除bigkey，而较耗时的内存释放会在后台线程中执行，不会影响到主线程。</font>
	
### 2.批量命令
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>尽量减少客户端和服务端之间来回的io次数，建议string/hash使用mget/mset代替get/set，hmget/hmset代替hget/hset;其他数据类型使用pipeline。</font>
	
### 3.避免复杂度过高命令
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>redis是单线程模型处理请求，在执行复杂度过高的命令（消耗更多cpu资源）时后面的请求会排队导致延迟，如 SORT、SINTER、SINTERSTORE、ZUNIONSTORE、ZINTERSTORE 等聚合类命令。</font>

### 4.注意容器类型数据
<font size=4 style='font-weight:bold'>建议规范</font>

<font size=4>1. 容器类型数据查询
<br/>list、hash、set、zset类型数据在查询时要先确认元素的数量，元素非常多时需要分批查询（ LRANGE/HASCAN/SSCAN/ZSCAN ），否则大量数据会导致延迟（如网络传输问题等）。
<br/>2. 容器类型数据删除
<br/>建议分批删除，如list执行多次lpop/rpop；hash/set/zset先执行hscan/sscan/scan查询元素，再执行hdel/srem/zrem,不关注元素的数量直接del会有大量内存释放，拖垮主线程性能。
</font>
