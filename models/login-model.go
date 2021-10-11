package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nikitamirzani323/gobackend_mastertoto/configs"
	"github.com/nikitamirzani323/gobackend_mastertoto/db"
	"github.com/nikitamirzani323/gobackend_mastertoto/helpers"
	"github.com/nleeper/goment"
)

func Login_Model(username, password string) (bool, error) {
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	flag := false

	var passwordDB string

	sql_select := `
		SELECT
		password    
		FROM ` + configs.DB_tbl_mst_admin + ` 
		WHERE username = ?
		AND statuslogin = "Y" 
	`
	row := con.QueryRowContext(ctx, sql_select, username)
	switch err := row.Scan(&passwordDB); err {
	case sql.ErrNoRows:
		return false, errors.New("Username and Password not found")
	case nil:
		flag = true
	default:
		return false, errors.New("Username and Password not found")
	}
	hashpass := helpers.HashPasswordMD5(password)
	if hashpass != passwordDB {
		return false, errors.New("Username and Password not found")
	}
	if flag {
		sql_update := `
			UPDATE ` + configs.DB_tbl_mst_admin + ` 
			SET lastlogin=?  
			WHERE username = ? 
		`
		rows_update, err_update := con.PrepareContext(ctx, sql_update)
		helpers.ErrorCheck(err_update)
		res_update, err_update := rows_update.ExecContext(ctx,
			tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)
		helpers.ErrorCheck(err_update)
		update, e := res_update.RowsAffected()
		helpers.ErrorCheck(e)
		if update > 0 {
			flag = true
			log.Println("Data Berhasil di save")
		}
	}
	return true, nil
}
