package user

import (
	"gin_weibo/app/models"
	"gin_weibo/app/requests"
)

type UserUpdateForm struct {
	Name                 string
	Password             string
	PasswordConfirmation string
}

// Validate : 验证函数
func (u *UserUpdateForm) Validate() (errors []string) {
	nameValidators := []requests.ValidatorFunc{
		requests.RequiredValidator(u.Name),
		requests.MaxLengthValidator(u.Name, 50),
	}
	nameMsgs := []string{
		"名称不能为空",
		"名称长度不能大于 50 个字符",
	}
	pwdValidators := []requests.ValidatorFunc{
		requests.RequiredValidator(u.Password),
		requests.MixLengthValidator(u.Password, 6),
		requests.EqualValidator(u.Password, u.PasswordConfirmation),
	}
	pwdMsgs := []string{
		"密码不能为空",
		"密码长度不能小于 6 个字符",
		"两次输入的密码不一致",
	}

	if u.Password == "" {
		errors = requests.RunValidators(
			requests.ValidatorMap{
				"name": nameValidators,
			},
			requests.ValidatorMsgArr{
				"name": nameMsgs,
			},
		)
	} else {
		errors = requests.RunValidators(
			requests.ValidatorMap{
				"name":     nameValidators,
				"password": pwdValidators,
			},
			requests.ValidatorMsgArr{
				"name":     nameMsgs,
				"password": pwdMsgs,
			},
		)
	}

	return errors
}

// ValidateAndSave 验证参数并且创建用户
func (u *UserUpdateForm) ValidateAndSave(user *models.User) (errors []string) {
	var err error
	errors = u.Validate()

	if len(errors) != 0 {
		return errors
	}

	// 更新用户
	user.Name = u.Name
	if u.Password != "" {
		user.Password = u.Password
	}

	if err = user.Encrypt(); err != nil {
		errors = append(errors, "用户更新失败: "+err.Error())
		return errors
	}

	if err = user.Update(); err != nil {
		errors = append(errors, "用户更新失败: "+err.Error())
		return errors
	}

	return []string{}
}