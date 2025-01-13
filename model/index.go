package model

import (
	"fangwu-backend/global"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

type Base struct {
	Id             int64      `json:"id"`
	CreatedAt      *time.Time `json:"created_at" gorm:"type:timestamp(3)"`
	LastModifiedAt *time.Time `json:"last_modified_at" gorm:"autoUpdateTime:milli;type:timestamp(3)"`
	Creator        *int64     `json:"creator" gorm:"index;"`
	LastModifier   *int64     `json:"last_modifier" gorm:"index;"`
}

// 需要软删除功能的model都要加上这个
type Delete struct {
	IsDeleted bool       `json:"is_deleted" gorm:"index;"`
	Deleter   *int64     `json:"deleter" gorm:"index;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:timestamp(3)"`
}

type Archive struct {
	ArchiveId     int64      `json:"archive_id"`
	ArchivedBy    *int64     `json:"archived_by" gorm:"index;"`
	ArchivedAt    *time.Time `json:"archived_at" gorm:"type:timestamp(3)"`
	ArchiveReason *string    `json:"archive_reason"`
}

func ConnectToDb() {
	ConnectToPgsql()

	//如果开启了redis，就连接到redis数据库
	if global.Config.Redis.Enabled {
		ConnectToRedis()
	}
}

func ConnectToPgsql() {
	//通过gorm连接数据库
	global.Db, err = gorm.Open(
		postgres.Open(global.Config.Db.DSN),
		&gorm.Config{},
	)
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}

	//使用gorm标准格式，创建连接池
	sqlDB, _ := global.Db.DB()
	// Set Max Idle Connections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// Set Max Open Connections 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// Set Connection Max Lifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = global.Db.AutoMigrate(
		&User{},              //用户
		&ArchivedUser{},      //已归档用户
		&RequestLog{},        //请求日志
		&File{},              //文件
		&AdminDiv{},          //行政区划
		&DictionaryType{},    //字典类型
		&DictionaryDetail{},  //字典详情
		&ForRent{},           //出租信息
		&ArchivedForRent{},   //已归档出租信息
		&Comment{},           //评论
		&Favorite{},          //收藏
		&SeekHouse{},         //求租信息
		&ArchivedSeekHouse{}, //已归档求租信息
		&ContactBlacklist{},  //联系方式黑名单
		&Complaint{},         //投诉
		&ArchivedComment{},   //已归档评论
		&UserBlacklist{},     //用户黑名单
		&Notification{},      //消息
		&Favorite{},          //收藏
		&ArchivedComplaint{}, //已归档投诉
		&Footprint{},         //足迹
		&ViewContact{},       //查看联系方式
		&Community{},         //小区
		&Member{},            //会员类型
		&UserMember{},        //用户的会员情况
	)
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}

	//初始化数据
	InitData()
}

func InitData() {
	executeSql()
	initDictionaryType()
	initDictionaryDetail()
	initMember()
}

func (a *Archive) Delete(archivedBy int64, archiveReason string) {
	a.ArchiveId = idgen.NextId()
	a.ArchivedBy = &archivedBy
	archivedAt := time.Now()
	a.ArchivedAt = &archivedAt
	a.ArchiveReason = &archiveReason
}

func ConnectToRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.DSN,
		Password: global.Config.Redis.Password,
		DB:       0,
	})

	if global.Rdb == nil {
		global.SugaredLogger.Panicln(err)
	}
}

func executeSql() {
	executeSqlForAdminDiv()
	executeSqlForCommunity()
}
