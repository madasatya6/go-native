package SettingModel

import (
	"fmt"
	"context"

	"github.com/madasatya6/go-native/applications/config"
)

type Setting struct{
	ID int `json:"id" query:"id" form:"id" validate:"required"`
	Name string `json:"name" query:"name" form:"name" validate:"required"` 
	Value string `json:"value" query:"value" form:"value" validate:"value" validate:"required"`
	FieldType string `json:"field_type" query:"field_type" form:"field_type" validate:"required"`
	Tab string `json:"tab" query:"tab" form:"tab" validate:"required"`
}

var Table = "setting"
var db = config.MySQL

func GetAll(ctx context.Context) ([]Setting, error) {

	var settings []Setting

	var queryText = fmt.Sprintf("SELECT*FROM setting")

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		return settings, err
	}

	for rowQuery.Next() {

		var setting Setting

		err = rowQuery.Scan(&setting.ID, &setting.Name, &setting.Value, &setting.FieldType, &setting.Tab,)

		if err != nil {
			return settings, err
		}

		settings = append(settings, setting)
	}

	return settings, nil
}

func GetOption(setting_id int) (Setting, error) {

	var opt Setting

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queryText := fmt.Sprintf("SELECT*FROM setting where id=%d", setting_id)

	var err = db.QueryRowContext(ctx, queryText).Scan(&opt.ID, &opt.Name, &opt.Value, &opt.FieldType, &opt.Tab,)
	if err != nil {
		return opt, err
	}

	return opt, nil
}

func GetValue(setting_id int) (string, error) {

	var opt Setting
	var value string

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queryText := fmt.Sprintf("SELECT*FROM setting where id=%d", setting_id)
	err := db.QueryRowContext(ctx, queryText).Scan(&opt.ID, &opt.Name, &opt.Value, &opt.FieldType, &opt.Tab,)
	if err != nil {
		return value, err
	}

	value = opt.Value
	return value, nil
}

func Update(ctx context.Context, setting_name string, value string) error {

	var query = fmt.Sprintf("update %s set value='%s' where name='%s'", Table, value, setting_name)

	tx, err := db.Begin()
	if err != nil {
		return err 
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return err 
	}

	err = tx.Commit()
	if err != nil {
		return err 
	}
	
	return nil
}