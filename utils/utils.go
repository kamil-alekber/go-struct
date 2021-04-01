package utils

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JSON struct {
}

func Stringify(val interface{}) (string, error) {
	data, err := json.Marshal(val)

	if err != nil {
		fmt.Printf("Error in Stringify JSON: %s", val)
	}
	return string(data), err
}

func Parse(data []byte) (interface{}, error) {
	var val interface{}
	err := json.Unmarshal(data, &val)

	if err != nil {
		fmt.Printf("Error in Parse JSON: %s", val)
	}

	return val, err
}

func MigrationUp(dbpool *pgxpool.Pool) {
	migrate(dbpool, "UP")

}

func MigrationDown(dbpool *pgxpool.Pool) {
	migrate(dbpool, "DOWN")
}

func migrate(dbpool *pgxpool.Pool, kind string) {
	kind = strings.ToLower(kind)
	kindList := []string{"down", "up"}

	isKind := InArray(kind, kindList)

	if !isKind {
		fmt.Printf("Kind '%s' is not present in the array: %s \n", kind, kindList)
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error reading working dir: %s\n", err)
	}
	migrationDir := fmt.Sprintf("%s/migrations", pwd)
	// fmt.Println(path, err)
	migrations, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		fmt.Printf("Error reading migrations folder: %s\n", err)
	}

	// filter migrations

	for _, v := range migrations {
		fileName := strings.ToLower(v.Name())
		if !strings.Contains(fileName, kind) {
			continue
		}
		filePath := fmt.Sprintf("%s/%s", migrationDir, v.Name())

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file content: %s \n %s", filePath, err)
		}
		fmt.Printf("Applying migratoion: %s\n", v.Name())
		row, err := dbpool.Query(context.Background(), string(data))
		row.Close()

		if err != nil {
			fmt.Printf("Error applying migrations %s\n", err)
		} else {
			fmt.Println("Done. Adding to migration table...")
		}

		row, err = dbpool.Query(context.Background(),
			"INSERT INTO migrations (id, title, tag, body, migration_type, migrated_at) VALUES (gen_random_uuid(), $1, $2, $3, $4, CURRENT_TIMESTAMP)", v.Name(), v.Name(), string(data), strings.ToUpper(kind))
		if err != nil {
			fmt.Printf("Error adding '%s' to migration's table: %s\n", v.Name(), err)
		} else {
			fmt.Println("Added.")
		}
		row.Close()
		fmt.Println("------ \n\r")
	}
}

func InArray(locate string, arr []string) bool {
	for _, v := range arr {
		if v == locate {
			return true
		}
	}
	return false
}

func ForEachStr(arr []string, cb func(v interface{}, i int, arr interface{})) {
	for i, v := range arr {
		cb(v, i, arr)
	}
}

func PgSqlRowsToJson(rows pgx.Rows) []byte {
	fields := rows.FieldDescriptions()
	fieldsArray := make([]string, len(fields))
	for i, field := range fields {
		fieldsArray[i] = string(field.Name)
	}

	res := make([]map[string]interface{}, 0)

	for rows.Next() {
		row, _ := rows.Values()
		obj := make(map[string]interface{})
		for i, val := range row {

			if fieldsArray[i] == "id" {
				assertedV, ok := val.([16]uint8)
				if ok {
					val = B2S(assertedV)
				}
			}

			obj[fieldsArray[i]] = val
		}
		res = append(res, obj)
	}

	b, _ := json.Marshal(res)
	return b
}

func B2S(bs [16]uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = v
	}
	return hex.EncodeToString(b)
}
