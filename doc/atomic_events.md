
## 数据结构
```
{
    "id": "31212", // 原子事件ID
    "timestamp": 1652788631, // 本机创建时间戳
    "cid": QmcRD4wkPPi6dig81r5sLj9Zm1gDCL4zgpEj9CfuRrGbzF, // 节点cid
     "version": 5808, // 事件版本
    "type": "log", // 原子事件类型
    "operations": [ // 自定义操作集
        {
            "cmd": "DeleteViews",
            "actions": [
                {
                    "n": "LD",
                    "p": [
                        "sheet",
                        "163538817972916224",
                        "views",
                        "viwf32f5fe7b4bb0b1f9aec56585ed2c17b"
                    ],
                    "LD": {
                        "beforeId": "viw642b8a6e7585abdf6e1c6c0a3bf8aa5b",
                        "view": {
                            "viewId": "viwf32f5fe7b4bb0b1f9aec56585ed2c17b",
    								 ... 此处省略相关视图配置
                        }
                    }
                }
            ]
        }
        ...........
    ],
    "properties": { // 用户额外定义属性信息
        "sheet_id": "1234455"
        ...........
    }
   
}
```