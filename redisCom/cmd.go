package redisCom

const (
	Set = "set"
	//set key value [EX seconds | PX milliseconds] [NX|XX] [KEEPTTL]
	//EX 设置超时时间，单位是秒
	//PX 设置超时时间，单位是毫秒
	//NX 当且仅当对应的key不存在时才进行设置
	//XX 当且仅当对应的key存在时才进行设置
	Setex       = "setex"
	Get         = "get"
	Del         = "del"
	Incr        = "incr"
	decr        = "decr"
	incrby      = "incrby"
	Decrby      = "decrby"
	Setnx       = "setnx"
	Mset        = "mset"
	Mget        = "mget"
	APPEND      = "APPEND"
	Strlen      = "strlen"
	Incrbyfloat = "incrbyfloat" //可以处理负值
	Getrange    = "getrange"    //getrange key [startIndex, endIndex] //获得对应key的值的指定区间的字符串
	Setrange    = "setrange"
	Scan        = "scan" //迭代的是当前数据库中的所有数据库键,默认的count是10，可以扫出所有的key

	Rpush   = "rpush"
	Lpush   = "lpush"
	Linsert = "linsert" //linsert key BEFORE|AFTER pivot element
	BEFORE  = "BEFORE"
	AFTER   = "AFTER"
	rinsert = "rinsert"
	Lrange  = "lrange" //lrange key start stop [start, stop]
	Lpop    = "lpop"
	Rpop    = "rpop"
	Lrem    = "lrem"
	// Ltrim 根据参数count的值，移除列表中与参数value相等的元素
	//lrem key count value
	//count > 0 从表头开始向表尾搜索，移除与value相等的元素，数量为count
	//count < 0 从表尾开始向表头搜索，移除与value相等的元素，数量为count的绝对值
	//count = 0 移除表中所有与value相等的值
	Ltrim  = "ltrim"
	Lindex = "lindex" //lindex key index
	Llen   = "llen"
	Lset   = "lset"

	Hget    = "hget"
	Hset    = "hset"
	Hgetall = "hgetall" //!谨慎使用,可能阻塞生产环境
	Hexists = "hexists"
	Hlen    = "hlen"
	Hmset   = "hmset"
	Hmget   = "hmget"
	Hvals   = "hvals"
	//返回哈希表所有子键的值，并不返回key
	//hvals key
	Hkeys        = "hkeys"
	Hincrby      = "hincrby"
	Hincrbyfloat = "hincrbyfloat"
	Hdel         = "hdel"  //hdel 是删除key中的field，del是删除key
	Hscan        = "hscan" //用于迭代hash类型中的键值对，第一个参数是一个数据库键

	Sadd        = "sadd"        //可以一次add多个//sadd key member [member ...]
	Scard       = "scard"       //计算集合的大小
	Sismember   = "sismember"   //某个key是不是集合中的元素
	Srandmember = "srandmember" //随机获取一些集合中的元素
	Smembers    = "smembers"    //获取集合中的全部元素，会发生阻塞，无序的 !谨慎使用,可能阻塞生产环境
	Sdiff       = "sdiff"       //计算两个集合之间的差集 //sdiff key [key ...] //seta - setb - setc - ...
	Sinter      = "sinter"      //计算两个集合的交集 //sinter key [key ...]// seta 交 setb 交 setc 交 ...
	Sunion      = "sunion"      //计算两个集合的并集 //重复项只取一项
	Spop        = "spop"        //随机弹出一个集合中的元素
	Srem        = "srem"        //从集合中移除元素，可以一次操作多个
	Sscan       = "sscan"       //用于迭代set集合中的元素，第一个参数是一个数据库键

	// Zadd zadd key score value ... scoren valuen
	Zadd = "zadd"
	//zrange key min max//[min, max]//0表示第一个成员 1 表示第二个成员; 负数下标 -1表示倒数第一个成员 -2表示倒数第二个成员
	Zrange = "zrange"
	//http://www.redis.cn/commands/zrangebylex.html
	//时间复杂度：O(log(N)+M) with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it O(log(N)).
	Zrangebylex      = "zrangebylex"
	Zrevrange        = "zrevrange"
	Zrangebyscore    = "zrangebyscore"
	Zrevrangebyscore = "zrevrangebyscore"
	Zscore           = "zscore"
	Zrank            = "zrank"
	Zrevrank         = "zrevrank"
	Zcount           = "zcount"
	Zinterstore      = "zinterstore"
	//ZINTERSTORE 时间复杂度: O(N*K)+O(M*log(M)) 这里N表示有序集合中成员数最少的数字，K表示有序集合数量。M表示结果集中重合的数量。
	//交集并集的复杂度很高，如果有bigkey的情况会严重阻塞主线程，建议一般不要使用，可以把两个zset的元素取出来，在内存中进行交并集运算这样不会阻塞redis主线程
	Zunionstore      = "zinterstore"
	Zincrby          = "zincrby"
	Zcard            = "zcard"
	Zrem             = "zrem"
	Zremrangebyrank  = "zremrangebyrank"
	Zremrangebyscore = "zremrangebyscore"
	Zpopmax          = "zpopmax" //移除有序集合中分数最大的count个元素
	Zpopmin          = "zpopmin" //移除有序组合中分数最小的count个元素
	Zscan            = "zscan"   //用于迭代sortset集合中的元素和元素对应的分值，第一个参数是一个数据库键，

	Eval = "eval"
	//使用方法
	//eval script numkeys key [key ...] arg [arg ...]

)

//其他类型的command
//keys * 打印出所有的key
//keys he*
//keys he[h-l] 第三个字母是h到l的范围
//keys he?  问号表示任意一位
////key相关命令在生产服严禁使用，因为生产服的key很多会把redis阻塞挂掉，复杂度为O(n)
//dbsize 计算key的总数时间复杂读O(1)
//expire name 3 设置3秒过期
//ttl name  查看name还有多长时间过期
//persist name 去掉name的过期时间
//type name 查看name的类型
//
////其他命令
//info 内存，cpu和主从相关
//client list 正在连接的绘画
//client kill ip:端口
//flushall 清空所有
//flushdb 只清空当前库
//select 数字 选择某个库，总共16个库
//monitor 记录操作日志
