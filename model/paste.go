package model

import (
	"errors"
	"github.com/PasteUs/PasteMeGoBackend/util/convert"
	"time"
)

type IPaste interface {
	Save() error
	Get() error
	Delete() error
	GetContent() string
	GetLang() string
	GetPassword() string
}

type AbstractPaste struct {
	IPaste
	Lang      string    `json:"lang" gorm:"type:varchar(16)"`     // 语言类型
	Content   string    `json:"content" gorm:"type:mediumtext"`   // 内容，最大长度为 16777215(2^24-1) 个字符
	Password  string    `json:"password" gorm:"type:varchar(32)"` // 密码
	ClientIP  string    `gorm:"type:varchar(64)"`                 // 用户 IP
	CreatedAt time.Time // 存储记录的创建时间
}

func (paste *AbstractPaste) GetContent() string {
	return paste.Content
}

func (paste *AbstractPaste) GetPassword() string {
	return paste.Password
}

func (paste *AbstractPaste) GetLang() string {
	return paste.Password
}

func (paste *AbstractPaste) beforeSave() error {
	if paste.Content == "" {
		return errors.New("empty content") // 内容为空，返回错误信息 "empty content"
	}
	if paste.Lang == "" {
		return errors.New("empty lang") // 语言类型为空，返回错误信息 "empty lang"
	}
	if paste.Password != "" {
		paste.Password = convert.String2md5(paste.Password) // 加密存储，设置密码
	}
	return nil
}
