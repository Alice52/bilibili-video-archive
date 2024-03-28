// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedLike = "archived_like"

// ArchivedLike 点赞视频
type ArchivedLike struct {
	Bvid       string         `gorm:"column:bvid;type:varchar(64);primaryKey" json:"bvid"`
	CreateTime *time.Time     `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
	Aid        int64          `gorm:"column:aid;type:bigint;not null;comment:bili aid" json:"aid"`                    // bili aid
	Cid        int64          `gorm:"column:cid;type:bigint;not null;comment:bili cid" json:"cid"`                    // bili cid
	Cover      *string        `gorm:"column:cover;type:varchar(256);comment:video cover" json:"cover"`                // video cover
	Duration   int64          `gorm:"column:duration;type:bigint;not null;comment:video duration" json:"duration"`    // video duration
	LikeTime   int64          `gorm:"column:like_time;type:bigint;not null;comment:video like time" json:"like_time"` // video like time
	SeasonID   int64          `gorm:"column:season_id;type:bigint;not null;comment:bili season id" json:"season_id"`  // bili season id
	Intro      *string        `gorm:"column:intro;type:text;comment:video intro" json:"intro"`                        // video intro
	Title      *string        `gorm:"column:title;type:text;comment:video title" json:"title"`                        // video title
	Type       int64          `gorm:"column:type;type:bigint;not null;comment:video type" json:"type"`                // video type
	Owner      *string        `gorm:"column:owner;type:json;comment:{"mid": 173986740, "name": "这个月-"}" json:"owner"` // {"mid": 173986740, "name": "这个月-"}
	Resp       *string        `gorm:"column:resp;type:json" json:"resp"`
	CntInfo    *string        `gorm:"column:cnt_info;type:json;comment:{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }" json:"cnt_info"` // {"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }
	VideoInfo  ArchivedVideo  `gorm:"foreignKey:bvid" json:"video_info"`
}

// TableName ArchivedLike's table name
func (*ArchivedLike) TableName() string {
	return TableNameArchivedLike
}
