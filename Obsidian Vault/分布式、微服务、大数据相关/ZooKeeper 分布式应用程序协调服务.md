# ZooKeeper 分布式应用程序协调服务
## 概述
ZooKeeper是一个[分布式](https://baike.baidu.com/item/%E5%88%86%E5%B8%83%E5%BC%8F/19276232?fromModule=lemma_inlink)的，开放源码的[分布式应用程序](https://baike.baidu.com/item/%E5%88%86%E5%B8%83%E5%BC%8F%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F/9854429?fromModule=lemma_inlink)协调服务，是[Google](https://baike.baidu.com/item/Google?fromModule=lemma_inlink)的Chubby一个[开源](https://baike.baidu.com/item/%E5%BC%80%E6%BA%90/246339?fromModule=lemma_inlink)的实现，是Hadoop和[Hbase](https://baike.baidu.com/item/Hbase/7670213?fromModule=lemma_inlink)的重要组件。它是一个为分布式应用提供一致性服务的软件，提供的功能包括：配置维护、域名服务、分布式同步、组服务等。

ZooKeeper的目标就是封装好复杂易出错的关键服务，将简单易用的接口和性能高效、功能稳定的系统提供给用户。

ZooKeeper包含一个简单的原语集，提供Java和C的接口。

ZooKeeper代码版本中，提供了分布式独享锁、选举、队列的接口，代码在\$zookeeper_home\\src\\recipes。其中分布锁和队列有[Java](https://baike.baidu.com/item/Java/85979?fromModule=lemma_inlink)和C两个版本，选举只有Java版本。
