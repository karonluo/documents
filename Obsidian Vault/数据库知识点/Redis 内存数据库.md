#T200 #Redis #DEV #OPS #数据库/非关系型数据库/内存数据库/Redis 

# Redis 内存数据库
---
## 概述
<font size=4>Redis is often referred to as a data structures server. 
What this means is that Redis provides access to mutable data structures via a set of commands, which are sent using a server-client model with TCP sockets and a simple protocol. So different processes can query and modify the same data structures in a shared way.
Data structures implemented into Redis have a few special properties:
Redis cares to store them on disk, even if they are always served and modified into the server memory. This means that Redis is fast, but that it is also non-volatile.
The implementation of data structures emphasizes memory efficiency, so data structures inside Redis will likely use less memory compared to the same data structure modelled using a high-level programming language.
Redis offers a number of features that are natural to find in a database, like replication, tunable levels of durability, clustering, and high availability.
Another good example is to think of Redis as a more complex version of memcached, where the operations are not just SETs and GETs, but operations that work with complex data types like Lists, Sets, ordered data structures, and so forth.</font>


## 安装部署

## 知识点
### [[Redis内存数据库设计优化和规范]]
