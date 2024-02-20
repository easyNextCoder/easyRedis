package dao

import (
	"easyRedis/redisPool"
	"easyRedis/redisTest/basic/sortedSet/proto"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var conn = redisPool.GetRedis()

func StrToInt64(str string) int64 {
	val64, _ := strconv.ParseInt(str, 10, 64)
	return val64
}

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func StrToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}



var expireday = 15

func expired(t time.Time) bool {
	if t.AddDate(0, 0, expireday).Compare(time.Now()) < 0 {
		return true
	}
	return false
}


//存储聊天逻辑
func alertCntKey(uid int64) string {
	return "cnt_" + Int64ToStr(uid)
}

func alertTimeKey(uid int64) string {
	return "time_" + Int64ToStr(uid)
}

//存储聊天信息
func chatId(a, b int64) string {
	if a > b {
		a, b = b, a
	}
	return "chat_" + Int64ToStr(a) + "_" + Int64ToStr(b)
}

//存储聊天已读位置最新时间戳
func readKey(uid int64)string{
	return "read_"+Int64ToStr(uid)
}







type Log struct {

}

func (self *Log)Debug(a string, v interface{}...){

}

func (self *Log)Error(a string, v ...interface{}){

}

var log Log


func deleteExpiredAlert(rightUid int64) error {
	uidTimesMp, err := redis.Int64Map(conn.Do("hgetall", alertTimeKey(rightUid)))

	if err != nil {
		return fmt.Errorf("deleteExpiredAlert err %s rightUid %d", err, rightUid)
	}

	for leftUid, t := range uidTimesMp {
		if expired(time.UnixMilli(t)) {
			_, err2 := conn.Do("hdel", alertTimeKey(rightUid), leftUid)
			if err2 != nil {
				return  fmt.Errorf("deleteExpiredAlert err2 %s rightUid %d leftUid %d", err2, rightUid, leftUid)
			}
			_, err3 := conn.Do("hdel", alertCntKey(rightUid), leftUid)
			if err3 != nil {
				return  fmt.Errorf("deleteExpiredAlert err3 %s rightUid %d leftUid %d", err3, rightUid, leftUid)
			}
		}
	}

	return nil
}

func saveUids(leftUid, rightUid int64) error {

	err := deleteExpiredAlert(leftUid)
	if err != nil {
		return fmt.Errorf("saveUids err %s", err)
	}

	_, err = conn.Do("hincrby", alertCntKey(leftUid), Int64ToStr(rightUid), 1)
	if err != nil {
		return fmt.Errorf("saveUids hincrby err %s", err)
	}
	_, err = conn.Do("hset", alertTimeKey(leftUid), Int64ToStr(rightUid), time.Now().UnixMilli())
	if err != nil {
		return fmt.Errorf("saveUids hset err %s", err)
	}

	return nil
}

func getUids(uid int64) ([]int64, error) {

	err := deleteExpiredAlert(uid)
	if err != nil {
		return nil, fmt.Errorf("getUids err %s", err)
	}

	senders, err := redis.Int64s(conn.Do("hkeys", alertCntKey(uid)))
	if err != nil {
		return []int64{}, fmt.Errorf("getUids err %s uid %d", err, uid)
	}
	return senders, nil
}




func SaveMsg(rightUid, leftUid int64, content string) error{//right为发送着， left为接收者

	millUnix := time.Now().UnixMilli()

	err := saveUids(rightUid, leftUid)
	if err != nil {
		return fmt.Errorf("saveMsg err1 %s", err)
	}

	err2 := saveUids(leftUid, rightUid)
	if err != nil {
		return fmt.Errorf("saveMsg err1 %s", err2)
	}

	var msg proto.Msg

	msg.Content = content
	msg.SenderUid = rightUid
	msg.ReceiverUid = leftUid
	msg.TimeMillUnix = time.Now().UnixMilli()
	msg.Now = time.Now().Format(time.DateTime)

	bs, err := json.Marshal(&msg)
	if err != nil{
		return fmt.Errorf("saveMsg marshal err %s", err)
	}

	_, err = conn.Do("zadd", chatId(leftUid, rightUid), millUnix, bs)
	if err != nil {
		return fmt.Errorf("saveMsg zadd err %s", err)
	}

	return nil
}

func getLastMsg(leftUid, rightUid int64)(*proto.OneChat, error){//拉取最后的100条信息
	bytes, err := redis.ByteSlices(conn.Do("zrevrangebyscore", chatId(leftUid, rightUid), "+inf", "-inf", "limit", 0, 100))
	if err != nil{
		return nil, fmt.Errorf("getLastMsg err %s", err)
	}
	return bytes2OneChat(bytes, leftUid, rightUid)
}

func getSenderInfo()*proto.Sender{
	return nil
}

func GetAllChat(uid int64) ([]*proto.OneChat, error) {
	uids, err := getUids(uid)
	if err != nil {
		return []*proto.OneChat{}, fmt.Errorf("getAllMsg err %s uid %d", err, uid)
	}

	var res []*proto.OneChat

	for _, leftUid := uids{
		v, err := getLastMsg(uid, leftUid)
		if err != nil{
			log.Error("getAllMsg err %s uid %d leftUid %d", err, uid, leftUid)
			continue
		}
		res = append(res, v)
	}


	return res, nil
}


func bytes2OneChat(bytes [][]byte, leftUid, rightUid int64) (*proto.OneChat, error){
	var res proto.OneChat
	for _, b := range bytes{
		var msg proto.Msg
		err := json.Unmarshal(b, &msg)
		if err != nil{
			log.Error("getNewMsg %s", err)
			continue
		}
		res.Chats = append(res.Chats, &msg)
	}

	res.Sender = getSenderInfo()

	rt, err := GetReadMaxTimeUnix(leftUid, rightUid)
	if err != nil{
		return nil, fmt.Errorf("getNewMsg err %s", err)
	}

	res.ReadAtTimeUnix = rt

	return &res, err
}


func GetMoreOldMsg(leftUid, rightUid int64,  maxTimeUnix int64)(*proto.OneChat, error) {
	bytes, err := redis.ByteSlices(conn.Do("zrevrangebyscore", chatId(leftUid, rightUid), maxTimeUnix-1, "-inf", "limit", 0, 100))
	if err != nil{
		return nil, fmt.Errorf("getMoreMsg err %s", err)
	}
	return bytes2OneChat(bytes, leftUid, rightUid)
}

func GetNewMsg(leftUid, rightUid int64, minTimeUnix int64)(*proto.OneChat, error){
	bytes, err := redis.ByteSlices(conn.Do("zrangebyscore", chatId(leftUid, rightUid), minTimeUnix+1, "+inf", "limit", 0, 100))
	if err != nil{
		return nil, fmt.Errorf("getNewMsg err %s", err)
	}
	return bytes2OneChat(bytes, leftUid, rightUid)
}

func GetReadMaxTimeUnix(leftUid, rightUid int64) (int64, error){
	res, err := redis.Int64(conn.Do("hget", rightUid, readKey(leftUid)))
	if err != nil{
		return 0, fmt.Errorf("GetReadMaxTimeUnix err %s", err)
	}
	return res, nil
}

func UpdateReadMaxTimeUnix(leftUid, rightUid, readTimeUnix int64)error{
	old, err := GetReadMaxTimeUnix(leftUid, rightUid)
	if err != nil{
		return fmt.Errorf("UpdateReadMaxTimeUnix err %s", err)
	}

	if readTimeUnix < old{
		return fmt.Errorf("UpdateReadMaxTimeUnix old %d update %d", old, readTimeUnix)
	}

	_, err = conn.Do("hset", rightUid, readKey(leftUid), readTimeUnix)
	if err != nil {
		return fmt.Errorf("UpdateReadMaxTimeUnix hset err %s", err)
	}

	return nil
}



//共用一个msgList
//onLoad->加载全部
//bind触发拉取最新的,更新msgList，同时setdata//解决刚好在当前变动问题
//onshow->展示第一条，同时更新正在看的最新的时间

