package models

import (
	"log"
)

func GetDataByQuery(query string) ([]map[string]interface{}, error) {
	data := []map[string]interface{}{}

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	colNames, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	for rows.Next() {
		myMap := map[string]interface{}{}

		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Println(err)
		}
		for i, col := range cols {
			myMap[colNames[i]] = BytesToString(col.([]byte))
		}
		data = append(data, myMap)
	}
	return data, nil
}

func BytesToString(data []byte) string {
	return string(data[:])
}
