package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
	"strings"
	"time"
)

type UserMemberGet struct {
	UserId int64 `json:"-"`
}

type UserMemberResult struct {
	City                                           AdminDivResult `json:"city,omitempty"`
	Type                                           string         `json:"type,omitempty"`
	RemainingTimesForViewingForRentContactPerDay   int            `json:"remaining_times_for_viewing_for_rent_contact_per_day,omitempty"`
	TotalTimesForViewingForRentContactPerDay       int            `json:"total_times_for_viewing_for_rent_contact_per_day,omitempty"`
	RemainingTimesForPublishingSeekHouse           int            `json:"remaining_times_for_publishing_seek_house,omitempty"`
	TotalTimesForPublishingSeekHouse               int            `json:"total_times_for_publishing_seek_house,omitempty"`
	RemainingTimesForViewingSeekHouseContactPerDay int            `json:"remaining_times_for_viewing_seek_house_contact_per_day,omitempty"`
	TotalTimesForViewingSeekHouseContactPerDay     int            `json:"total_times_for_viewing_seek_house_contact_per_day,omitempty"`
	RemainingTimesForPublishingForRent             int            `json:"remaining_times_for_publishing_for_rent,omitempty"`
	TotalTimesForPublishingForRent                 int            `json:"total_times_for_publishing_for_rent,omitempty"`
	RemainingTopTimesPerMonth                      int            `json:"remaining_top_times_per_month,omitempty"`
	TotalTopTimesPerMonth                          int            `json:"total_top_times_per_month,omitempty"`
	RemainingTimesForchangingCity                  int            `json:"remaining_times_for_changing_city,omitempty"`
	TotalTimesForchangingCity                      int            `json:"total_times_for_changing_city,omitempty"`
	ExpiredAt                                      time.Time      `json:"expired_at,omitempty"`
}

func (u *UserLogin) Login() (result *LoginResult, resCode int, errDetail *util.ErrDetail) {
	//如果开启了验证码验证
	if global.Config.Captcha.EnabledForLogin {
		//如果没有传入验证码id或验证码内容
		if u.CaptchaId == "" || u.Captcha == "" {
			return nil, util.ErrorMissingCaptcha, nil
		}

		//开始校验
		permitted := u.Verify()
		if !permitted {
			return nil, util.ErrorWrongCaptcha, nil
		}
	}

	//获取用户记录
	var user model.User
	err := global.Db.Model(model.User{}).
		Where("username = ?", u.Username).
		First(&user).Error
	if err != nil {
		return nil, util.ErrorInvalidUsernameOrPassword, util.GetErrDetail(err)
	}

	// 校验密码
	var encryptAndCompare util.EncryptAndCompare
	permitted := encryptAndCompare.Compare(u.Password, user.Password)
	if !permitted {
		return nil, util.ErrorInvalidUsernameOrPassword, util.GetErrDetail(err)
	}

	// 生成token
	token, resCode, errDetail := util.GenerateToken(user.Id)
	if resCode != util.Success {
		return nil, resCode, errDetail
	}

	result = &LoginResult{AccessToken: token}

	return result, util.Success, nil
}

func (u *UserGet) Get() (result *UserResult, resCode int, errDetail *util.ErrDetail) {
	//获取用户记录
	var user model.User
	err := global.Db.
		Where("id = ?", u.Id).
		First(&user).Error
	if err != nil {
		return nil, util.ErrorFailToGetUser, util.GetErrDetail(err)
	}

	var tmpRes UserResult
	tmpRes.Id = user.Id

	if user.MobilePhone != nil {
		tmpRes.MobilePhone = *user.MobilePhone
	}

	if user.EmailAddress != nil {
		tmpRes.EmailAddress = *user.EmailAddress
	}

	return &tmpRes, util.Success, nil
}

func (u *UserCreate) Create() (result *UserResult, resCode int, errDetail *util.ErrDetail) {
	var user model.User

	// 检查用户名是否已存在
	var count int64
	global.Db.Model(model.User{}).
		Where("username =?", u.Username).Count(&count)
	if count > 0 {
		return nil, util.ErrorUsernameExists, nil
	}

	//校验用户注册限制
	if global.Config.RegisterLimit.Enabled {
		parts := strings.Split(*u.Ip, ".")
		if len(parts) == 4 {
			ipPrefix := strings.Join(parts[:3], ".")
			var count int64
			timeLimit := time.Now().AddDate(0, 0, -1*global.Config.RegisterLimit.Interval)
			global.Db.Model(model.User{}).
				Where("register_ip LIKE ?", ipPrefix+".%").
				Where("created_at > ?", timeLimit).
				Count(&count)
			if count > 0 {
				return nil, util.ErrorTooFrequentRegistration, nil
			}
		}
	}

	user.Username = u.Username
	var err error
	// 加密密码
	var encryptAndCompare util.EncryptAndCompare
	user.Password, err = encryptAndCompare.Encrypt(u.Password)

	if err != nil {
		return nil, util.ErrorFailToEncrypt, util.GetErrDetail(err)
	}

	user.RegisterIp = u.Ip
	isValid := true
	user.IsValid = &isValid
	user.TimesForViewingContact = 5

	err = global.Db.Create(&user).Error
	if err != nil {
		return nil, util.ErrorFailToCreateUser, util.GetErrDetail(err)
	}

	var userLogin UserLogin
	userLogin.Username = u.Username
	userLogin.Password = u.Password
	tmpRes, resCode, errDetail := userLogin.Login()
	if resCode != util.Success {
		return nil, util.ErrorFailToCreateUser, errDetail
	}

	return &UserResult{AccessToken: tmpRes.AccessToken}, util.Success, nil
}

func (u *UserUpdate) Update() (resCode int, errDetail *util.ErrDetail) {
	user := make(map[string]any)

	if u.Id > 0 {
		user["last_modifier"] = u.Id
	}

	if u.EmailAddress != nil {
		user["email_address"] = *u.EmailAddress
	}

	if u.MobilePhone != nil {
		user["mobile_phone"] = *u.MobilePhone
	}

	err := global.Db.Model(model.User{}).
		Where("id = ?", u.Id).
		Updates(user).Error
	if err != nil {
		return util.ErrorFailToUpdateUser, util.GetErrDetail(err)
	}

	return util.Success, nil
}

func (u *UserDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	tx := global.Db.Begin()

	//获取用户记录
	var user model.User
	err := tx.Where("id = ?", u.Id).
		First(&user).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToGetUser, util.GetErrDetail(err)
	}

	//将用户记录存入归档表
	var archivedUser model.ArchivedUser
	archivedUser.Archive.Delete(u.Id, "用户注销")
	archivedUser.User = user
	err = tx.Create(&archivedUser).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteUser, util.GetErrDetail(err)
	}

	//删除用户记录
	err = tx.Delete(&user).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteUser, util.GetErrDetail(err)
	}

	tx.Commit()
	return util.Success, nil
}
