package svc

import (
	"easyRedis/redisTest/basic/sortedSet/dao"
	"easyRedis/redisTest/basic/sortedSet/proto"
)

var log = new(dao.Log)

type AllChatReq struct {
	Uid int64 `json:"uid"`
}

type AllChatRsp struct {
	Chats []*proto.OneChat `json:"chats"`
}

func GetAllChat(req *AllChatReq, rsp *AllChatRsp) error {
	chats, err := dao.GetAllChat(req.Uid)
	if err != nil {
		log.Error("GetAllChat err %s req %+v", err, *req)
		return err
	}

	rsp.Chats = chats

	return nil
}

type SaveMsgReq struct {
	SenderUid   int64  `json:"sender_uid"`
	ReceiverUid int64  `json:"receiver_uid"`
	Content     string `json:"content"`
}

type SaveMsgRsp struct {
	Result bool `json:"result"`
}

func SaveMsg(req *SaveMsgReq, rsp *SaveMsgRsp) error {
	err := dao.SaveMsg(req.SenderUid, req.ReceiverUid, req.Content)
	if err != nil {
		log.Error("SaveMsg err %s req %+v", err, *req)
		return err
	}

	rsp.Result = true
	return nil
}

type MoreOldMsgReq struct {
	SenderUid   int64 `json:"sender_uid"`
	ReceiverUid int64 `json:"receiver_uid"`
	MaxTimeUnix int64 `json:"max_time_unix"`
}

type MoreOldMsgRsp struct {
	Msgs *proto.OneChat `json:"msgs"`
}

func GetMoreOldMsg(req *MoreOldMsgReq, rsp *MoreOldMsgRsp) error {
	msgs, err := dao.GetMoreOldMsg(req.ReceiverUid, req.SenderUid, req.MaxTimeUnix)
	if err != nil {
		log.Error("GetMoreOldMsg err %s req %+v", err, *req)
		return err
	}
	rsp.Msgs = msgs
	return nil
}

type NewMsgReq struct {
	SenderUid   int64 `json:"sender_uid"`
	ReceiverUid int64 `json:"receiver_uid"`
	MinTimeUnix int64 `json:"min_time_unix"`
}

type NewMsgRsp struct {
	Msgs *proto.OneChat `json:"msgs"`
}

func GetNewMsg(req *NewMsgReq, rsp *NewMsgRsp) error {
	msgs, err := dao.GetNewMsg(req.ReceiverUid, req.SenderUid, req.MinTimeUnix)
	if err != nil {
		log.Error("GetNewMsg err %s req %+v", err, *req)
		return err
	}
	rsp.Msgs = msgs
	return nil
}

type UpdateReadPointReq struct {
	SenderUid    int64 `json:"sender_uid"`
	ReceiverUid  int64 `json:"receiver_uid"`
	ReadTimeUnix int64 `json:"Read_time_unix"`
}

type UpdateReadPointRsp struct {
	Result bool `json:"result"`
}

func UpdateReadMaxTimeUnix(req *UpdateReadPointReq, rsp *UpdateReadPointRsp) error {
	err := dao.UpdateReadMaxTimeUnix(req.ReceiverUid, req.ReceiverUid, req.ReadTimeUnix)
	if err != nil {
		log.Error("UpdateReadMaxTimeUnix err %s req %+v", err, *req)
		return err
	}
	rsp.Result = true
	return nil
}
