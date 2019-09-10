package service

import (
	uuid "github.com/satori/go.uuid"
	"lilith/common"
	"lilith/dao"
	"lilith/model/dto"
	"lilith/model/entity"
	"log"
	"time"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService() *UserService {
	service := &UserService{
		dao: dao.NewUserDao(),
	}
	return service
}

func (s *UserService) GetUserPwdByUsername(username string) (entity.UserPwd, error) {
	pwd, e := s.dao.GetUserPwdByUsername(username)
	if e != nil {
		log.Println(e)
		return pwd, e
	}
	return pwd, e
}

func (s *UserService) GetUserInfo(username string) (entity.UserInfo, error) {
	return s.dao.GetUserInfo(username)
}

func (s *UserService) GetUserInfoById(id string) (entity.UserInfo, error) {
	return s.dao.GetUserInfoById(id)
}

func (s *UserService) CreateUserPwd(dto dto.RegisterDto) {
	var userPwd entity.UserPwd
	salt := common.RandString(10)
	userPwd.Id = uuid.NewV4().String()
	userPwd.IsOK = "ok"
	userPwd.Salt = salt
	userPwd.Password = common.GenPwd(dto.Password, salt)
	userPwd.Username = dto.Username
	userPwd.CreateTime = time.Now().UnixNano() / 1e6
	i, e := s.dao.CreateUserPwd(userPwd)
	if e != nil {
		log.Println(e)
	}
	if i != 1 {
		log.Println("create user pwd error...")
	}
}

func (s *UserService) CreateUserInfo(dto dto.RegisterDto) {
	var userInfo entity.UserInfo
	userInfo.Id = uuid.NewV4().String()
	userInfo.Username = dto.Username
	userInfo.Nick = dto.Username
	userInfo.Position = "普通咸鱼"
	userInfo.CreateTime = time.Now().UnixNano() / 1e6
	i, e := s.dao.CreateUserInfo(userInfo)
	if e != nil {
		log.Println(e)
	}
	if i != 1 {
		log.Println("create user info error...")
	}
}

func (s *UserService) UpdatePwd(username, pwd string) int {
	return s.dao.UpdatePwd(username, pwd)
}

func (s *UserService) UpdateUserInfo(uDto dto.UpdateUserDto, id string) int {
	return s.dao.UpdateUserInfo(uDto, id)
}
