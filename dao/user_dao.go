package dao

import (
	"database/sql"
	"lilith/config"
	"lilith/model/dto"
	"lilith/model/entity"
)

type UserDao struct {
	db *sql.DB
}

func NewUserDao() *UserDao {
	dao := &UserDao{
		db: config.GetDB(),
	}
	return dao
}

func (d *UserDao) GetUserPwdByUsername(username string) (entity.UserPwd, error) {
	sqlStr := "select t.id,t.username,t.password,t.salt,t.is_ok,t.create_time from t_pwd t where t.username=?"
	row := d.db.QueryRow(sqlStr, username)
	var userPwd entity.UserPwd
	_ = row.Scan(&userPwd.Id, &userPwd.Username, &userPwd.Password, &userPwd.Salt, &userPwd.IsOK, &userPwd.CreateTime)
	return userPwd, nil
}

func (d *UserDao) GetUserInfo(username string) (entity.UserInfo, error) {
	sqlStr := "select t.id,t.username,t.nick,t.email,t.photo,t.desc,t.position,t.home_page,t.tags,t.create_time from t_user t where t.username=?"
	row := d.db.QueryRow(sqlStr, username)
	var userInfo entity.UserInfo
	_ = row.Scan(&userInfo.Id, &userInfo.Username, &userInfo.Nick, &userInfo.Email, &userInfo.Photo, &userInfo.Desc, &userInfo.Position, &userInfo.HomePage, &userInfo.Tags, &userInfo.CreateTime)
	return userInfo, nil
}

func (d *UserDao) GetUserInfoById(id string) (entity.UserInfo, error) {
	sqlStr := "select t.id,t.username,t.nick,t.email,t.photo,t.desc,t.position,t.home_page,t.tags,t.create_time from t_user t where t.id=?"
	row := d.db.QueryRow(sqlStr, id)
	var userInfo entity.UserInfo
	_ = row.Scan(&userInfo.Id, &userInfo.Username, &userInfo.Nick, &userInfo.Email, &userInfo.Photo, &userInfo.Desc, &userInfo.Position, &userInfo.HomePage, &userInfo.Tags, &userInfo.CreateTime)
	return userInfo, nil
}

func (d *UserDao) CreateUserPwd(uEn entity.UserPwd) (int, error) {
	sqlStr := "insert into t_pwd(id,username,password,salt,is_ok,create_time) values (?,?,?,?,?,?)"
	result, e := d.db.Exec(sqlStr, uEn.Id, uEn.Username, uEn.Password, uEn.Salt, uEn.IsOK, uEn.CreateTime)
	if e != nil {
		return 0, e
	}
	i, e := result.RowsAffected()
	return int(i), e
}

func (d *UserDao) CreateUserInfo(uEn entity.UserInfo) (int, error) {
	sql := "insert into t_user(`id`,`username`,`nick`,`email`,`photo`,`desc`,`position`,`home_page`,`tags`,`create_time`) values (?,?,?,?,?,?,?,?,?,?)"
	result, e := d.db.Exec(sql, uEn.Id, uEn.Username, uEn.Nick, uEn.Email, uEn.Photo, uEn.Desc, uEn.Position, uEn.HomePage, uEn.Tags, uEn.CreateTime)
	if e != nil {
		return 0, e
	}
	i, e := result.RowsAffected()
	return int(i), e
}

func (d *UserDao) UpdatePwd(username, pwd string) int {
	sqlStr := "update t_pwd as p set p.password=? where p.username=?"
	result, e := d.db.Exec(sqlStr, pwd, username)
	if e != nil {
		return 0
	}
	i, e := result.RowsAffected()
	return int(i)
}

func (d *UserDao) UpdateUserInfo(uDto dto.UpdateUserDto, id string) int {
	sqlStr := "update t_user as p set p.nick=?, p.position=? where p.id=?"
	result, e := d.db.Exec(sqlStr, uDto.Nick, uDto.Position, id)
	if e != nil {
		return 0
	}
	i, e := result.RowsAffected()
	return int(i)
}
