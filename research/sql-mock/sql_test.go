package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSQLMock(t *testing.T) {
	t.Parallel()

	const (
		casesNumber = 100
	)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"advert_id", "excluded"})

	expected := map[uint64][]string{}

	for i := 0; i < casesNumber; i++ {
		strs := []string{strconv.Itoa(i), strconv.Itoa(i * 10)}

		rows.AddRow(i, strings.Join(strs, ","))

		expected[uint64(i)] = strs
	}

	// Ожидаем вызова метода Query с конкретным SQL-запросом и возвращаем заданные строки
	// Запрос требует экранирования
	mock.ExpectQuery(`
		select es.advert_id, excluded
		from \"Adverts\" a join excluded_search es on a.\"Id\" = es.advert_id
		where \(\"Type\" = 8 or \"Type\" = 9\) and \"StatusId\" = 9
		`).
		WillReturnRows(rows)
}
