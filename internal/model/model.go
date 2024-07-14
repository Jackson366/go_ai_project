package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goAiproject/global"
	"goAiproject/pkg/setting"
	"time"
)

type Model struct {
	ID         uint64    `gorm:"primary_key;column:id" json:"id"`
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:createTime;comment:'创建时间'" json:"createTime"`
	UpdateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:updateTime;comment:'更新时间'" json:"updateTime"`
	IsDelete   uint8     `gorm:"type:tinyint;default:0;column:isDelete;comment:'是否删除'" json:"isDelete"`
}

func NewDBEngin(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(time.Now())
			}
		}

		if updateTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if updateTimeField.IsBlank {
				_ = updateTimeField.Set(time.Now())
			}
		}
	}
}

// 更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdateTime", time.Now())
	}
}

// 删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("IsDelete")
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}

}

// 添加额外的空格
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
