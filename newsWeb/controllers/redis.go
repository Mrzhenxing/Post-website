package controllers

import (
	"gomodule/redigo/redis"
	"github.com/astaxie/beego"
)

func init() {

	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	beego.Error("redis链接失败")
	//	return
	//}
	//
	//resp, err := conn.Do("get", "newsWeb")
	//
	//result, _ := redis.String(resp, err)
	//
	//beego.Info("获取的数据为：", result)

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		beego.Info("redis连接失败", err)
		return
	}
	//有连接就有关闭
	defer conn.Close()
	//操作函数

	//conn.Send()发送

	//conn.Flush() 刷新

	//conn.Receive()接收结果  三个为一组

	//来一条执行一条   由于返回值是个接口类型 处理单个数据用
	//resp,err:=conn.Do("get","c1")
	//result,err:=redis.String(resp,err)
	//if err!=nil{
	//	beego.Info("获取数据失败")
	//	return
	//}

	resp,err:=conn.Do("mget","t1","t2","t3")
	result,err:=redis.Strings(resp,err)
	if err!=nil{
		beego.Info("获取数据失败")
		return
	}

	beego.Info("获取数据为：", result)


	//获取多个数据
	//resp, err := conn.Do("mget", "t1", "t2", "t3")
	////redis中获取的多个数据都是同一个类型 用切片strings  不同类型用value和scan
	//result, err := redis.Values(resp, err)
	//if err != nil {
	//	beego.Info("获取数据失败")
	//	return
	//}
	//
	////把对应的函数扫描到相应变量里面
	//var v1 ,v2 int
	//var v3 string
	//redis.Scan(result,&v1,&v2,&v3)


	//beego.Info("获取数据为：", v1,v2,v3)

}
