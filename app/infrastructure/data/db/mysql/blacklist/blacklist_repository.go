package blacklist

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
	"database/sql"
	"log"
)

const (
	queryBlackList    = "SELECT bl.id, CONCAT('PROD_',bl.sku) sku FROM black_list as bl;"
	queryBlackListExclude    = "SELECT bl.id, CONCAT('PROD_',bl.sku) sku FROM black_list as bl LEFT JOIN black_list_exclude as ble ON (bl.sku=ble.sku) WHERE ble.sku is NULL;"
	queryListForExclude = "SELECT ble.id, CONCAT('PROD_',ble.sku) sku FROM black_list_exclude as ble;"
	queryDeleteExclude = "DELETE FROM black_list_exclude;"

)

type blackListMysqlRepository struct {
	client *sql.DB
}

func NewBlackListMysqlRepository(client *sql.DB) *blackListMysqlRepository {
	return &blackListMysqlRepository{client: client}
}

func (bl *blackListMysqlRepository) ShowBlackList(option string) (string, error) {

	var query string
	listado := ""

	query = queryBlackList

	if option == "WITHOUTEXCLUDE" {
		query = queryBlackListExclude
	}
	if option == "EXCLUDEITEMS" {
		query = queryListForExclude
	}

	stmt, err := bl.client.Prepare(query)

	if err != nil {
		log.Println("error when trying to prepare get user statement")
		return listado, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Println("error when trying to find user")

		return listado, err
	}

	defer rows.Close()

	results := make([]*models.BlackList,0)


	for rows.Next() {

		var bl models.BlackList

		if err := rows.Scan(&bl.ID, &bl.Sku); err != nil {
			log.Println("error when trying to scan user")
			return listado, err
		}

		listado += bl.Sku + ","

		results = append(results, &bl)

	}

	//fmt.Printf(listado)

	if len(results) == 0 {
		return listado, err
	}

	return listado, nil
}

func (bl *blackListMysqlRepository) RemoveExcludeItems()  error {

	stmt, err := bl.client.Prepare(queryDeleteExclude)

	if err != nil {
	log.Println("error when trying to prepare get user statement")
	return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
	log.Println("error when trying to delete exclude items")

	return  err
	}

	defer bl.client.Close()

	return nil
}