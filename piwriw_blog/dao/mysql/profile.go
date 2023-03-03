package mysql

import "piwriw_blog/models"

func GetProfile(id string) (profile *models.Profile, err error) {
	sqlStr := `select * from profile where id=?`
	profile = new(models.Profile)
	err = db.Get(profile, sqlStr, id)
	if err != nil {
		return
	}
	return
}
func UpdateProfile(uid string, profile *models.Profile) (err error) {
	sqlStr := "update profile set name=?,`desc`=?,qqchat=?,wechat=?,weibo=?,bili=?,email=?,img=?,avatar=? ,github=?,gitee=? ,public_account=? where id=?"
	_, err = db.Exec(sqlStr, profile.Name, profile.Desc, profile.Qqchat, profile.Wechat, profile.Weibo, profile.Bili, profile.Email, profile.Img, profile.Avatar, profile.GitHub, profile.GitEE, profile.PublicAccount, uid)
	if err != nil {
		return
	}
	return
}
