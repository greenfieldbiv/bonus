package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const jwtTableName = "jwt"

type JwtTable struct {
	Id           int64  `db:"id"`
	AccountId    int64  `db:"account_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
}

type DatabaseJwtStorage struct {
	db *sqlx.DB
}

func NewDatabaseJwtStorage(db *sqlx.DB) JwtStorage {
	return &DatabaseJwtStorage{db: db}
}

func (d *DatabaseJwtStorage) GetAccessToken(id int64) string {
	if isExist := d.IsExistAccessTokenById(id); !isExist {
		return ""
	}
	var acToken = ""
	if err := d.db.Get(&acToken, fmt.Sprintf("SELECT access_token FROM %s where account_id = $1", jwtTableName), id); err != nil {
		return ""
	}
	return acToken
}

func (d *DatabaseJwtStorage) GetRefreshToken(id int64) string {
	if isExist := d.IsExistAccessTokenById(id); !isExist {
		return ""
	}
	var rToken = ""
	if err := d.db.Get(&rToken, fmt.Sprintf("SELECT refresh_token FROM %s where account_id = $1", jwtTableName), id); err != nil {
		return ""
	}
	return rToken
}

func (d *DatabaseJwtStorage) PutAccessToken(id int64, access string) bool {
	if isExistAccessToken := d.IsExistAccessToken(id, access); !isExistAccessToken {
		rows, err := d.db.Query(fmt.Sprintf("INSERT INTO %s (ACCOUNT_ID, ACCESS_TOKEN) VALUES ($1, $2)", jwtTableName), id, access)
		defer func() {
			err = rows.Close()
		}()
		return err == nil
	}
	rows, err := d.db.Query(fmt.Sprintf("UPDATE %s SET access_token = $1 where account_id = $2", jwtTableName), access, id)
	defer func() {
		err = rows.Close()
	}()
	return err == nil
}

func (d *DatabaseJwtStorage) PutRefreshToken(id int64, refresh string) bool {
	if isExistRefreshToken := d.IsExistRefreshToken(id, refresh); !isExistRefreshToken {
		rows, err := d.db.Query(fmt.Sprintf("INSERT INTO %s (ACCOUNT_ID, refresh_token) VALUES ($1, $2)", jwtTableName), id, refresh)
		defer func() {
			err = rows.Close()
		}()
		return err == nil
	}
	rows, err := d.db.Query(fmt.Sprintf("UPDATE %s SET refresh_token = $1 where account_id = $2", jwtTableName), refresh, id)
	defer func() {
		err = rows.Close()
	}()
	return err == nil
}

func (d *DatabaseJwtStorage) DeleteAccessToken(id int64, access string) bool {
	// для базы такая имплементация не нужна
	return false
}

func (d *DatabaseJwtStorage) DeleteRefreshToken(id int64, refreshToken string) bool {
	rows, err := d.db.Query(fmt.Sprintf("DELETE FROM %s where account_id = $1 and refresh_token = $2", jwtTableName), id, refreshToken)
	defer func() {
		err = rows.Close()
	}()
	return err == nil
}

func (d *DatabaseJwtStorage) IsExistAccessToken(id int64, access string) bool {
	var count = 0
	if err := d.db.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s where account_id = $1 and access_token = $2", jwtTableName), id, access); err != nil {
		return false
	}
	return count != 0
}

func (d *DatabaseJwtStorage) IsExistRefreshToken(id int64, refresh string) bool {
	var count = 0
	if err := d.db.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s where account_id = $1 and refresh_token = $2", jwtTableName), id, refresh); err != nil {
		return false
	}
	return count != 0
}

func (d *DatabaseJwtStorage) IsExistAccessTokenById(id int64) bool {
	var aToken *string
	if err := d.db.Get(&aToken, fmt.Sprintf("SELECT access_token FROM %s where account_id = $1", jwtTableName), id); err != nil {
		return false
	}
	return aToken != nil
}

func (d *DatabaseJwtStorage) IsExistRefreshTokenById(id int64) bool {
	var rToken *string
	if err := d.db.Get(&rToken, fmt.Sprintf("SELECT refresh_token FROM %s where account_id = $1", jwtTableName), id); err != nil {
		return false
	}
	return rToken != nil
}

func (d *DatabaseJwtStorage) PutTokens(id int64, access string, refresh string) bool {
	if d.IsExistAccessTokenById(id) || d.IsExistRefreshTokenById(id) {
		rows, err := d.db.Query(fmt.Sprintf("UPDATE %s SET access_token = $1 , refresh_token = $2 where account_id = $3", jwtTableName), access, refresh, id)
		defer func() {
			err = rows.Close()
		}()
		return err == nil
	}
	rows, err := d.db.Query(fmt.Sprintf("INSERT INTO %s (ACCOUNT_ID, access_token, refresh_token) VALUES ($1, $2, $3)", jwtTableName), id, access, refresh)
	defer func() {
		err = rows.Close()
	}()
	return err == nil
}

func (d *DatabaseJwtStorage) GetRefreshTokenByAccessToken(accessToken string) string {
	var rToken = ""
	if err := d.db.Get(&rToken, fmt.Sprintf("SELECT refresh_token FROM %s where access_token = $1", jwtTableName), accessToken); err != nil {
		return ""
	}
	return rToken
}

func (d *DatabaseJwtStorage) GetRefreshTokenById(id string) string {
	var rToken = ""
	idStr, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ""
	}
	if err := d.db.Get(&rToken, fmt.Sprintf("SELECT refresh_token FROM %s where account_id = $1", jwtTableName), idStr); err != nil {
		return ""
	}
	return rToken
}

func (d *DatabaseJwtStorage) DeleteTokenById(id string) bool {
	idStr, err := strconv.ParseInt(id, 10, 64)
	rows, err := d.db.Query(fmt.Sprintf("DELETE FROM %s where account_id = $1", jwtTableName), idStr)
	defer func() {
		err = rows.Close()
	}()
	return err == nil
}
