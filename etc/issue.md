# issue

## 启动

    geth --datadir ./data --networkid 15 --port 30303 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcvhosts "*" --rpcapi "db,net,eth,web3,personal" --rpccorsdomain "*" --ws --wsaddr "localhost" --wsport "8546" --wsorigins "*" --nat "any" --nodiscover --dev --dev.period 1 console 2> 1.log
    
    go run main.go -c ../etc/eilieili.dev.toml
    go run main.go -c ../etc/eilieili_home.dev.toml

## 知识回顾

    1. 使用layout.html时, 无法直接使用xx.html, 作为新连接网站

    2. 静态文件: public,  网站入口:/public/index.html

    3. 从前端提交的json数据, 使用c.Ctx.ReadJSON()进行解析

    4. 定义viewmodels.Obj, 当作中间介进行前端与数据库进行交互

    5. 设置cookies(本地), session(服务器). 

    6. 返回mvc.Result, 不需要返回resp结构, 需在返回的Data中设置所需返回给客服端的数据

## 进行投票时的有关作品数据

    1. 通过编写爬虫, 写入文件备注部分

## 暂时遇到了bug

    1. /contents, 无法点击<原创认证>

## 数据存储redis+mysql

    1、mysql支持sql查询，可以实现一些关联的查询以及统计；
    2、redis对内存要求比较高，在有限的条件下不能把所有数据都放在redis；
    3、mysql偏向于存数据，redis偏向于快速取数据，但redis查询复杂的表关系时不如mysql，所以可以把热门的数据放redis，mysql存基本数据

    如果你数据有20G，只使用Redis，就需要占用20G内存
    如果Redis+MySQL，Redis只储存热数据，MySQL储存冷数据，那么就不需要把全部数据放内存里了。内存里只有热数据。20G数据储存在硬盘。
