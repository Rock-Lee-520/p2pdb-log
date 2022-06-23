
## 数据结构
```
{
    "id": "163538817972916224", // 原子事件ID
    "resource_id":""//资源id
    "timestamp": 1652788631, // 本机创建时间戳
    "cid": QmcRD4wkPPi6dig81r5sLj9Zm1gDCL4zgpEj9CfuRrGbzF, // 节点cid
    "last_id":"163538817972916224"  //因果关系事件id(上一个关联事件id),为空代表没有因果关系
    "revision": 5808, // 资源版本,单调递增
    "type": "log", // 原子事件自定义类型
    "operations": [ // 自定义操作集
        ...........
    ],
    "properties": { // 自定义属性信息
        ...........
    }
   
}
```



@startuml

class Node <<Entity>> {
    - Cid  // 节点ID,一个原子事件一个节点
    - Type // 节点类型 默认log
    - LamportClock //兰伯特时钟
    - ReceiveTimestamp // 本机接收到事件时间
    - SendTimestamp //对等节点发送事件时间
    - Last_id //因果关系事件id(上一个关联事件id) 为空代表没有因果关系                 
}


class Object <<Entity>> {
    - Cid  //节点id
    - ObjectId  // 对象ID  一个节点包含多个对象
    - Resource_id //资源id
    - Timestamp //本机创建时间戳
    - Last_id //因果关系事件id(上一个关联事件id) 为空代表没有因果关系
    - Operations //自定义操作集
    - Properties //自定义属性信息                 
}

class Link <<Entity>> {
    - LastCid  // 上一个节点CID
    - Cid //当前节点Cid
    - Size //数据大小
}
