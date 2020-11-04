package blacklist

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
	"database/sql"
	"fmt"
	"log"
)

const (
	errorNoRows = "no rows in result set"
	queryBlackList    = "SELECT bl.id, CONCAT('PROD_',bl.sku) sku FROM black_list as bl;"
	queryBlackListExclude    = "SELECT bl.id, CONCAT('PROD_',bl.sku) sku FROM black_list as bl LEFT JOIN black_list_exclude as ble ON (bl.sku=ble.sku) WHERE ble.sku is NULL;"
	queryListForExclude = "SELECT ble.id, CONCAT('PROD_',ble.sku) sku FROM black_list_exclude as ble;"
)

type blackListMysqlRepository struct {
	client *sql.DB
}

func NewBlackListMysqlRepository(client *sql.DB) *blackListMysqlRepository {
	return &blackListMysqlRepository{client: client}
}

func (bl *blackListMysqlRepository) ShowBlackList(option string) ([]*models.BlackList, error) {

	var query string

	query = queryBlackList

	if option == "WITHEXCLUDE" {
		query = queryBlackListExclude
	}
	if option == "FOREXCLUDE" {
		query = queryListForExclude
	}

	stmt, err := bl.client.Prepare(query)

	if err != nil {
		log.Println("error when trying to prepare get user statement")
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Println("error when trying to find user")

		return nil, err
	}

	defer rows.Close()

	results := make([]*models.BlackList,0)

	listado := ""

	for rows.Next() {

		var bl models.BlackList

		if err := rows.Scan(&bl.ID, &bl.Sku); err != nil {
			log.Println("error when trying to scan user")
			return nil, err
		}

		listado += bl.Sku + ","

		results = append(results, &bl)

	}

	fmt.Printf(listado)

	if len(results) == 0 {
		return nil, err
	}

	return results, nil
}
