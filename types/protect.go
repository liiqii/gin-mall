package types

import "time"

type ProtectCreateReq struct {
	ID        string    `form:"id" json:"id"`
	Uid       string    `form:"uid" json:"uid"`
	Remark    string    `form:"remark" json:"remark"`
	StartTime time.Time `form:"start_time" json:"start_time"`
	EndTime   time.Time `form:"end_time" json:"end_time" `
}

type ProtectListReq struct {
	Id          uint      `form:"id" json:"id"`
	Uid         string    `form:"uid" json:"uid"`
	ApiId       string    `form:"apiId" json:"apiId"`
	ApiKey      string    `form:"apiKey" json:"apiKey"`
	Remark      string    `form:"remark" json:"remark"`
	Status      uint      `form:"status" json:"status"`
	Hidden      uint      `form:"hidden" json:"hidden"`
	StartTime   time.Time `form:"startTime" json:"startTime"`
	EndTime     time.Time `form:"endTime" json:"endTime" `
	CreatedTime time.Time `form:"createdTime" json:"createdTime"`
	UpdatedTime time.Time `form:"updatedTime" json:"updatedTime" `
}
