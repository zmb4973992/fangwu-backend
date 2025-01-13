package global

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	v      = viper.New()
	Config config
	// Logger zap的标准logger，速度更快，但是输入麻烦，用于取代gin的logger
	Logger *zap.Logger
	// SugaredLogger zap的加糖logger，速度慢一点点，但是输入方便，自己用
	SugaredLogger *zap.SugaredLogger
	Db            *gorm.DB
	Rdb           *redis.Client
	Captcha       *base64Captcha.Captcha
)

// 这层只是中间的汇总层，只是包内引用、不展示，所以小写
type config struct {
	Access        accessConfig
	Gin           ginConfig
	Db, Redis     dbConfig
	Jwt           jwtConfig
	Log           logConfig
	Upload        uploadConfig
	Download      downloadConfig
	Email         emailConfig
	Paging        pagingConfig
	RateLimit     rateLimitConfig
	Captcha       captchaConfig
	SuperPassword superPasswordConfig
	RegisterLimit registerLimitConfig
}

type accessConfig struct {
	Port int
}

type ginConfig struct {
	Mode string
}

type dbConfig struct {
	Enabled  bool
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	// Data Source Name 数据源名称，用于描述数据库连接的字符串
	DSN string
}

type jwtConfig struct {
	SecretKey    string
	ValidityDays int
}

type logConfig struct {
	FileName  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
}

type uploadConfig struct {
	StoragePath      string
	TmpStoragePath   string
	ThumbnailPath    string
	TmpThumbnailPath string
	MaxSize          int64
	AllowedExts      []string
}

type downloadConfig struct {
	PublicIp     string
	RelativePath string
	FullPath     string
}

type emailConfig struct {
	OutgoingMailServer string
	Port               int
	Account            string
	Password           string
}

type pagingConfig struct {
	PageSize    int
	MaxPageSize int
}

type rateLimitConfig struct {
	Limit float64
	Burst int
}

type captchaConfig struct {
	Length          int
	Width           int
	Height          int
	NoiseCount      int
	EnabledForLogin bool
}

type superPasswordConfig struct {
	Enabled  bool
	Password string
}

type registerLimitConfig struct {
	Enabled  bool
	Interval int
}

func initConfig() {
	Config.Access.Port = v.GetInt("access.port")

	Config.Gin.Mode = v.GetString("gin.mode")
	if Config.Gin.Mode != "debug" && Config.Gin.Mode != "test" && Config.Gin.Mode != "release" {
		Config.Gin.Mode = "debug"
	}

	Config.Db.Host = v.GetString("db.host")
	Config.Db.Port = v.GetString("db.port")
	Config.Db.DbName = v.GetString("db.name")
	Config.Db.Username = v.GetString("db.username")
	Config.Db.Password = v.GetString("db.password")
	Config.Db.DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		Config.Db.Host, Config.Db.Username, Config.Db.Password, Config.Db.DbName, Config.Db.Port)

	Config.Redis.Enabled = v.GetBool("redis.enabled")
	Config.Redis.Host = v.GetString("redis.host")
	Config.Redis.Port = v.GetString("redis.port")
	Config.Redis.Password = v.GetString("redis.password")
	Config.Redis.DSN = fmt.Sprintf("%s:%s", Config.Redis.Host, Config.Redis.Port)

	Config.Jwt.SecretKey = v.GetString("jwt.secret-key")
	Config.Jwt.ValidityDays = v.GetInt("jwt.validity-days")

	Config.Log.FileName = v.GetString("log.Log-path") + "/status.Log"
	Config.Log.MaxSize = v.GetInt("log.log-max-size")
	Config.Log.MaxBackup = v.GetInt("log.log-max-backup")
	Config.Log.MaxAge = v.GetInt("log.log-max-age")
	Config.Log.Compress = v.GetBool("log.log-compress")

	if runtime.GOOS == "windows" {
		Config.Upload.StoragePath = v.GetString("upload.storage-path.windows") + "/"
		Config.Upload.TmpStoragePath = v.GetString("upload.storage-path.windows") + "_tmp/"
		Config.Upload.ThumbnailPath = v.GetString("upload.thumbnail-path.windows") + "/"
		Config.Upload.TmpThumbnailPath = v.GetString("upload.thumbnail-path.windows") + "_tmp/"
	} else if runtime.GOOS == "linux" {
		Config.Upload.StoragePath = v.GetString("upload.storage-path.linux") + "/"
		Config.Upload.TmpStoragePath = v.GetString("upload.storage-path.linux") + "_tmp/"
		Config.Upload.ThumbnailPath = v.GetString("upload.thumbnail-path.linux") + "/"
		Config.Upload.TmpThumbnailPath = v.GetString("upload.thumbnail-path.linux") + "_tmp/"
	} else {
		SugaredLogger.Panicln(errors.New("不支持的系统环境，请更换为windows或linux"))
	}

	Config.Upload.MaxSize = v.GetInt64("upload.max-size") << 20
	Config.Upload.AllowedExts = v.GetStringSlice("upload.allowed-exts")

	Config.Download.RelativePath = v.GetString("Download.relative-path") + "/"
	var err error
	if v.GetString("download.public-ip") != "" {
		Config.Download.PublicIp = v.GetString("download.public-ip")
	} else {
		Config.Download.PublicIp, err = getLocalIp()
		if err != nil {
			SugaredLogger.Errorln(err)
		}
	}

	Config.Download.FullPath = "http://" + Config.Download.PublicIp +
		":" + strconv.Itoa(Config.Access.Port) + Config.Download.RelativePath

	Config.Email.OutgoingMailServer = v.GetString("email.outgoing-mail-server")
	Config.Email.Port = v.GetInt("email.port")
	Config.Email.Account = v.GetString("email.account")
	Config.Email.Password = v.GetString("email.password")

	Config.Paging.PageSize = v.GetInt("paging.page-size")
	Config.Paging.MaxPageSize = v.GetInt("paging.max-page-size")

	Config.RateLimit.Limit = v.GetFloat64("rate-limit.limit")
	Config.RateLimit.Burst = v.GetInt("rate-limit.burst")

	Config.Captcha.Length = v.GetInt("captcha.length")
	Config.Captcha.Width = v.GetInt("captcha.width")
	Config.Captcha.Height = v.GetInt("captcha.height")
	Config.Captcha.NoiseCount = v.GetInt("captcha.noise-count")
	Config.Captcha.EnabledForLogin = v.GetBool("captcha.enabled-for-login")

	Config.SuperPassword.Enabled = v.GetBool("superPassword.enabled")
	Config.SuperPassword.Password = v.GetString("superPassword.password")

	Config.RegisterLimit.Enabled = v.GetBool("register-limit.enabled")
	Config.RegisterLimit.Interval = v.GetInt("register-limit.interval")
}

func LoadConfig() {
	//配置文件的路径
	v.AddConfigPath("./config/")
	//配置文件的前缀
	v.SetConfigName("config")
	//配置文件的后缀
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("配置文件无法读取: %w", err)
	}

	//配置文件热更新
	v.OnConfigChange(func(e fsnotify.Event) {
		initConfig()
	})

	v.WatchConfig()
	initConfig()
}

func getLocalIp() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("未找到本机IP地址")
}
