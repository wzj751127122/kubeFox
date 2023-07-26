package logic

import (
	"database/sql"
	"errors"

	"k8s-platform/dao"
	"k8s-platform/model"
	"k8s-platform/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, userInfo *model.AdminLoginInput) (string, error) {
	// fmt.Println("finduser",)
	user, err := dao.Find(ctx, &model.SysUser{UserName: userInfo.UserName})
	if err != nil {
		return "", err
	}
	// fmt.Println(user)
	if !utils.CheckPassword(userInfo.Password, user.Password) {
		return "", errors.New("密码错误,请重新输入")
	}

	token, err := utils.JwtToken.GenerateToken(utils.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.UserName,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func LoginOut(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid, Status: sql.NullInt64{Int64: 0, Valid: true}}
	return dao.Updates(ctx, user)
}

func GetUserInfo(ctx *gin.Context, uid int, aid uint) (*model.UserInfoOut, error) {
	user, err := dao.Find(ctx, &model.SysUser{ID: uid})
	if err != nil {
		return nil, err
	}
	menus, err := GetMenuByAuthorityID(ctx, aid)
	if err != nil {
		return nil, err
	}
	var outRules []string
	rules := GetPolicyPathByAuthorityId(aid)
	for _, rule := range rules {
		item := rule.Path + "," + rule.Method
		outRules = append(outRules, item)
	}
	return &model.UserInfoOut{
		User:      *user,
		Menus:     menus,
		RuleNames: outRules,
	}, nil
}

func SetUserAuth(ctx *gin.Context, uid int, aid uint) error {
	user := &model.SysUser{ID: uid, AuthorityId: aid}
	return dao.Updates(ctx, user)
}

func DeleteUser(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid}
	return dao.Delete(ctx, user)
}

func ChangePassword(ctx *gin.Context, uid int, info *model.ChangeUserPwdInput) error {
	userDB := &model.SysUser{ID: uid}
	user, err := dao.Find(ctx, userDB)
	if err != nil {
		return err
	}

	if !utils.CheckPassword(info.OldPwd, user.Password) {
		return errors.New("原密码错误,请重新输入")
	}

	//生成新密码
	user.Password, err = utils.GenSaltPassword(info.NewPwd)
	if err != nil {
		return err
	}
	return dao.Updates(ctx, user)
}

func ResetPassword(ctx *gin.Context, uid int) error {
	newPwd, err := utils.GenSaltPassword("kubeFox")
	if err != nil {
		return err
	}
	user := &model.SysUser{ID: uid, Password: newPwd}
	return dao.Updates(ctx, user)
}
