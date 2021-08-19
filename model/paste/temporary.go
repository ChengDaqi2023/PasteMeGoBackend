package paste

import (
	"github.com/PasteUs/PasteMeGoBackend/model/dao"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	EnumTime  = "time"
	EnumCount = "count"
	OneMonth  = 31 * 24 * 60
	MaxCount  = 3
)

func init() {
	dao.CreateTable(&Temporary{})
}

// Temporary 临时
type Temporary struct {
	*AbstractPaste        // 公有字段
	ExpireType     string // 过期类型
	Expiration     uint64 // 过期的数据
}

func (Temporary) TableName() string {
	return "temporary"
}

// Save 成员函数，保存
func (paste *Temporary) Save() error {
	paste.Key = Generator(8, true, &paste)
	return dao.DB.Create(&paste).Error
}

// Delete 成员函数，删除
func (paste *Temporary) Delete() error {
	return dao.DB.Delete(&paste).Error
}

// Get 成员函数，查看
func (paste *Temporary) Get(password string) error {
	err := dao.DB.Transaction(func(tx *gorm.DB) error {
		if e := tx.Find(&paste).Error; e != nil {
			return e
		}

		if paste.ExpireType == EnumTime {
			duration := time.Minute * time.Duration(paste.Expiration)
			if time.Now().After(paste.CreatedAt.Add(duration)) {
				if e := tx.Delete(&paste).Error; e != nil {
					return e
				}
				return gorm.ErrRecordNotFound
			}
		}

		if e := paste.checkPassword(password); e != nil {
			return e
		}

		if paste.ExpireType == EnumCount {
			if paste.Expiration <= 1 {
				if e := tx.Delete(&paste).Error; e != nil {
					return e
				}
			} else {
				return tx.Model(&paste).Update("expiration", paste.Expiration-1).Error
			}
		}
		return nil
	})
	return err
}

func exist(key string, model interface{}) bool {
	count := uint8(0)
	dao.DB.Model(model).Where("`key` = ?", key).Count(&count)
	return count > 0
}
