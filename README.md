## 使用go_mod搭建的线上订餐系统后端  
#### 在命令行输入 ```go mod tidy ``` 导入包关系
#### 运行 ```go run```
### version 1.0
功能实现 ：登录功能，提供短信登录和密码登录，想使用短信登录，必须更改app.json中阿里云的秘钥信息，详情参照阿里云短信官方文档


 *注* 需要自行修改数据库密码等等（app.json）
 
 
 ---------------------------
 ### version 2.0
 基本实现了前台用户的功能，对于用户信息模块功能还有待增加，需要在此基础上加入分布式文件系统或者使用对象存储oss，其次还要考虑到session的读取，加入redis缓存等。
 

 ##  主体功能
 
 ####	用户登录
 
 ##### 获取验证码
 
 1. URL地址：`/api/login_sms`
 
 2. 请求方式： get
 
 3. 请求参数：
 
 | 参数名 | 参数类型 | 参数长度 | 参数说明 | 必填 |
 | ------ | -------- | -------- | -------- | ---- |
 | phone  | 字符串   | 11       | 手机号   | y    |
 
 4. A 成功
 
    ```json
    {
        "code": 0,
        "data": "发送成功"
    }
    ```
 
    B 失败
 
    ```json
    {
        "code": 1,
        "data": "参数解析失败"
    }
    ```
 
    
 
 ##### 		通过短信登录
 
 1. URL地址：`/api/login_sms`
 
 2. 请求方式：post
 
 3. 请求参数：
 
 | 参数名 | 参数类型 | 参数长度 | 参数说明   | 必填 |
 | ------ | -------- | -------- | ---------- | ---- |
 | phone  | 字符串   | 11       | 手机号     | y    |
 | code   | 字符串   | 6        | 短信验证码 | y    |
 
 4. A 成功
 
    ```json
    {
        "code": 0,
        "data": {
            "id": 1,
            "user_name": "15362124173",
            "mobile": "15362124173",
            "password": "",
            "register_time": 1600413679,
            "avatar": "",
            "balance": 0,
            "is_active": 0,
            "city": ""
        },
        "msg": "成功"
    }
    ```
 
    B 失败
 
    ```json
    {
        "code": 1,
        "data": "登录失败"
    }
    ```
 
    
 
 ##### 通过密码登录
 
 1. URL地址：`/api/login_pwd`
 
 2. 请求方式：post
 
 3. 请求参数：
 
 | 参数名 | 参数类型 | 参数长度 | 参数说明 | 必填 |
 | ------ | -------- | -------- | -------- | ---- |
 | phone  | 字符串   | 11       | 手机号   | y    |
 | pwd    | 字符串   |          | 密码     | y    |
 
 4. A 成功
 
    ```json
    {
        "code": 0,
        "data": {
            "id": 2,
            "user_name": "13011110916",
            "mobile": "13011110916",
            "password": "e10adc3949ba59abbe56e057f20f883e",
            "register_time": 1600502543,
            "avatar": "",
            "balance": 0,
            "is_active": 0,
            "city": ""
        },
        "msg": "成功"
    }
    ```
 
    B 失败
 
    ```json
    {
        "code": 1,
        "data": "登录失败"
    }
    ```
 
    
 
 #### 	获取食品类别
 
 1. URL地址：`/api/goods/:id`
 
 2. 请求方式：get
 
 3. 请求参数：
 
 | 参数名 | 参数类型 | 参数长度 | 参数说明 | 必填 |
 | ------ | -------- | -------- | -------- | ---- |
 | id     | 字符串   |          | 商铺id   | y    |
 
 4. 接口文档
 
    A 成功
 
    ```json
    {
        "code": 0,
        "data": [
            {
                "id": 1,
                "name": "小小鲜肉包",
                "description": "滑蛋牛肉粥(1份)+小小鲜肉包(4只)",
                "icon": "",
                "sell_count": 14,
                "price": 25,
                "old_price": 29,
                "shop_id": 1
            },
            {
                "id": 2,
                "name": "滑蛋牛肉粥+小小鲜肉包",
                "description": "滑蛋牛肉粥(1份)+小小鲜肉包(3只)",
                "icon": "",
                "sell_count": 6,
                "price": 35,
                "old_price": 41,
                "shop_id": 1
            },
            {
                "id": 3,
                "name": "滑蛋牛肉粥+绿甘蓝馅饼",
                "description": "滑蛋牛肉粥(1份)+绿甘蓝馅饼(1张)",
                "icon": "",
                "sell_count": 2,
                "price": 25,
                "old_price": 30,
                "shop_id": 1
            },
            {
                "id": 4,
                "name": "茶香卤味蛋",
                "description": "咸鸡蛋",
                "icon": "",
                "sell_count": 688,
                "price": 2.5,
                "old_price": 3,
                "shop_id": 1
            },
            {
                "id": 5,
                "name": "韭菜鸡蛋馅饼(2张)",
                "description": "韭菜鸡蛋馅饼",
                "icon": "",
                "sell_count": 381,
                "price": 10,
                "old_price": 12,
                "shop_id": 1
            }
        ],
        "msg": "成功"
    }
    ```
 
    B 失败
 
    ```json
    {
        "code": 0,
        "data": "网络异常，请稍后再试！"
    }
    ```
 
    
 
 #### 	商家在售食品信息的展示
 
 1. URL地址：`/api/categories`
 
 2. 请求方式：get
 
 3. 请求参数：none
 
 4. 接口文档
 
    A 成功
 
    ```json
    {
        "code": 0,
        "data": [
            {
                "id": 1,
                "title": "品质美食",
                "description": "好吃的品质美食",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 2,
                "title": "甜点饮品",
                "description": "好吃的甜点饮品",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 3,
                "title": "超市便利",
                "description": "快捷的超市购物",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 4,
                "title": "蔬菜水果",
                "description": "新鲜的蔬菜水果",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 5,
                "title": "药品",
                "description": "保证安全的药品",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 6,
                "title": "汉堡披萨",
                "description": "好吃的汉堡披萨",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 7,
                "title": "快食简餐",
                "description": "好吃的快食简餐",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            },
            {
                "id": 8,
                "title": "家常菜",
                "description": "好吃的家常菜",
                "image_url": "",
                "link_url": "",
                "is_in_serving": false
            }
        ],
        "msg": "成功"
    }
    ```
 
    B 失败
 
    ```json
    {
        "code": 0,
        "msg": "获取食物类别失败",
    }
    ```
 
    
 
 #### 	  商家信息的查询的拓展
 
 #####     	获取附近所有商家信息
 
 1. URL地址：`/api/shops`
 
 2. 请求方式：get
 
 3. 请求参数：
 
 | 参数名    | 参数类型 | 参数长度 | 参数说明 | 必填 |
 | --------- | -------- | -------- | -------- | ---- |
 | longitude | 字符串   |          | 经度     |      |
 | latitude  | 字符串   |          | 纬度     |      |
 
 4. 接口返回
 
    A 成功:
 
    ```json
    {
        "code": 0,
        "data": [
            {
                "id": 486,
                "name": "东来顺旗舰店",
                "promotion_info": "老北京正宗涮羊肉,非物质文化遗产",
                "address": "北京市天河区东圃镇汇彩路38号1领汇创展商务中心401",
                "phone": "13544323775",
                "status": 1,
                "longitude": 113.41724,
                "latitude": 23.1127,
                "image_path": "",
                "is_new": true,
                "is_premium": true,
                "rating": 4.2,
                "rating_count": 372,
                "recent_order_num": 542,
                "minimum_order_amount": 20,
                "delivery_fee": 5,
                "opening_hours": "09:00/21:30",
                "Supports": [
                    {
                        "id": 1,
                        "name": "准时达",
                        "description": "准时达",
                        "icon_name": "",
                        "icon_color": ""
                    },
                    {
                        "id": 3,
                        "name": "美味鲜",
                        "description": "美味鲜",
                        "icon_name": "",
                        "icon_color": ""
                    }
                ]
            },
            {
                "id": 487,
                "name": "北京酒家",
                "promotion_info": "北京第一家传承300年酒家",
                "address": "北京市海淀区上下九商业步行街内",
                "phone": "13257482341",
                "status": 0,
                "longitude": 113.24826,
                "latitude": 23.11488,
                "image_path": "",
                "is_new": true,
                "is_premium": true,
                "rating": 4.2,
                "rating_count": 871,
                "recent_order_num": 923,
                "minimum_order_amount": 20,
                "delivery_fee": 5,
                "opening_hours": "8:30/20:30",
                "Supports": [
                    {
                        "id": 1,
                        "name": "准时达",
                        "description": "准时达",
                        "icon_name": "",
                        "icon_color": ""
                    },
                    {
                        "id": 2,
                        "name": "食无忧",
                        "description": "食无忧",
                        "icon_name": "",
                        "icon_color": ""
                    }
                ]
            },
            {
                "id": 488,
                "name": "和平鸽饺子馆",
                "promotion_info": "吃饺子就来和平鸽饺子馆",
                "address": "北京市越秀区德政中路171",
                "phone": "17098764762",
                "status": 1,
                "longitude": 113.27521,
                "latitude": 23.12092,
                "image_path": "",
                "is_new": true,
                "is_premium": true,
                "rating": 4.2,
                "rating_count": 273,
                "recent_order_num": 483,
                "minimum_order_amount": 20,
                "delivery_fee": 5,
                "opening_hours": "8:30/20:30",
                "Supports": [
                    {
                        "id": 1,
                        "name": "准时达",
                        "description": "准时达",
                        "icon_name": "",
                        "icon_color": ""
                    },
                    {
                        "id": 3,
                        "name": "美味鲜",
                        "description": "美味鲜",
                        "icon_name": "",
                        "icon_color": ""
                    }
                ]
            }
        ],
        "msg": "成功"
    }
    ```
 
    B 失败：
 
    ```"json
    {
        "code": 0,
        "msg": "暂未获取到商户信息",
    }
    ```
 
    
 
    
 
 #####     模糊搜索具体商家信息
 
 1. URL地址：`/api/search_shops`
 
 2. 请求方式：get
 
 3. 请求参数：
 
 | 参数名    | 参数类型 | 参数长度 | 参数说明   | 必填 |
 | --------- | -------- | -------- | ---------- | ---- |
 | longitude | 字符串   |          | 经度       | n    |
 | latitude  | 字符串   |          | 纬度       | n    |
 | keywprds  | 字符串   |          | 查询关键字 | n    |
 
 4. 接口返回
 
    A 成功
 
    ```json
    {
        "code": 0,
        "data": [
            {
                "id": 486,
                "name": "东来顺旗舰店",
                "promotion_info": "老北京正宗涮羊肉,非物质文化遗产",
                "address": "北京市天河区东圃镇汇彩路38号1领汇创展商务中心401",
                "phone": "13544323775",
                "status": 1,
                "longitude": 113.41724,
                "latitude": 23.1127,
                "image_path": "",
                "is_new": true,
                "is_premium": true,
                "rating": 4.2,
                "rating_count": 372,
                "recent_order_num": 542,
                "minimum_order_amount": 20,
                "delivery_fee": 5,
                "opening_hours": "09:00/21:30",
                "Supports": [
                    {
                        "id": 1,
                        "name": "准时达",
                        "description": "准时达",
                        "icon_name": "",
                        "icon_color": ""
                    },
                    {
                        "id": 3,
                        "name": "美味鲜",
                        "description": "美味鲜",
                        "icon_name": "",
                        "icon_color": ""
                    }
                ]
            }
        ],
        "msg": "成功"
    }
    ```
 
    B 失败：
 
    ```"json
    {
        "code": 0,
        "msg": "暂未获取到商户信息",
    }
    ```
 
    