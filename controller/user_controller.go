package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"lilith/common"
	"lilith/model/dto"
	"lilith/service"
	"net/http"
)

var (
	userService *service.UserService
)

func init() {
	userService = service.NewUserService()
}

func Login(c *gin.Context) {
	response := dto.Response{C: c}
	var loginDto dto.LoginDto
	e := c.BindJSON(&loginDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(loginDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userPwd, e := userService.GetUserPwdByUsername(loginDto.Username)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.UsernamePasswordError)
		return
	}
	if userPwd.Username == "" {
		response.Error(http.StatusUnauthorized, common.UsernamePasswordError)
		return
	}
	if userPwd.IsOK != "ok" {
		response.Error(http.StatusUnauthorized, common.AccountNoOKError)
		return
	}
	validatePwd := common.ValidatePwd(userPwd.Password, loginDto.Password, userPwd.Salt)
	if validatePwd {
		userInfo, e := userService.GetUserInfo(loginDto.Username)
		if e != nil {
			response.Error(http.StatusInternalServerError, common.ServerError)
			return
		}
		token := common.GenToken(userInfo.Id)
		userDto := dto.UserDto{
			UserInfo: userInfo,
			Token:    token,
		}
		response.Success(http.StatusOK, userDto)
		return
	}
	response.Error(http.StatusUnauthorized, common.UsernamePasswordError)
}

func Register(c *gin.Context) {
	response := dto.Response{C: c}
	var registerDto dto.RegisterDto
	e := c.BindJSON(&registerDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(registerDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userPwd, e := userService.GetUserPwdByUsername(registerDto.Username)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ServerError)
		return
	}
	if userPwd.Username != "" {
		response.Error(http.StatusUnauthorized, common.UserHasExistedError)
		return
	}
	userService.CreateUserPwd(registerDto)
	userService.CreateUserInfo(registerDto)
	response.Success(http.StatusOK, nil)
}

func ChangePwd(c *gin.Context) {
	response := dto.Response{C: c}
	var pwdDto dto.ChangePwdDto
	e := c.BindJSON(&pwdDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(pwdDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userId, _ := c.Get("userId")
	info, e := userService.GetUserInfoById(userId.(string))
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ServerError)
		return
	}
	pwd, e := userService.GetUserPwdByUsername(info.Username)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ServerError)
		return
	}
	enPwd := common.GenPwd(pwdDto.OPwd, pwd.Salt)
	if enPwd != pwd.Password {
		response.Error(http.StatusInternalServerError, common.InputPwdError)
		return
	}
	newPwd := common.GenPwd(pwdDto.NPwd, pwd.Salt)
	updatePwd := userService.UpdatePwd(info.Username, newPwd)
	if updatePwd != 1 {
		response.Error(http.StatusInternalServerError, common.UpdatePwdError)
		return
	}
	response.Success(http.StatusOK, info.Username)
}

func UpdateUserInfo(c *gin.Context) {
	response := dto.Response{C: c}
	var uDto dto.UpdateUserDto
	e := c.BindJSON(&uDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(uDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userId, _ := c.Get("userId")
	i := userService.UpdateUserInfo(uDto, userId.(string))
	if i != 1 {
		response.Error(http.StatusInternalServerError, common.UpdateInfoError)
		return
	}
	response.Success(http.StatusOK, uDto)
}
