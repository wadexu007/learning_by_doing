package models

import (
	"encoding/csv"
	"log"
	"os"

	"main.go/config"
)

func ReadData(fileName string) ([][]string, error) {

	f, err := os.Open(config.Conf.FILE_PATH + fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	// if _, err := r.Read(); err != nil {
	// 	return [][]string{}, err
	// }

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func WriteData(fileName string, record []string) {
	full_path_file := config.Conf.FILE_PATH + fileName
	f, err := os.OpenFile(full_path_file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if os.IsNotExist(err) {
		log.Println("[INFO] file not exist, now try to crete " + full_path_file)
		error := os.MkdirAll(config.Conf.FILE_PATH, os.ModePerm)
		if error != nil {
			log.Println(error)
		}
		f, err = os.Create(full_path_file)

	}
	if err != nil {
		log.Println("[ERROR] Failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(record); err != nil {
		log.Println("[ERROR] Error writing record to file", err)
	}
}
