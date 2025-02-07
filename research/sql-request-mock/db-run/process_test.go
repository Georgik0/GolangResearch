package db_run

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	normalizer "gitlab.wildberries.ru/advertising/monolib/norma"
	mockedindex "gitlab.wildberries.ru/advertising/monolib/norma/common/query-index/mock"
	"go.uber.org/mock/gomock"
)

func TestNormalizer(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	qIndex := mockedindex.NewMockMockQueryIndex(ctrl)
	qIndex.EXPECT().Search("some_word").
		Return("normalized_word", true).
		AnyTimes()

	normalizer.SetMockedIndex(qIndex)
	normalizer.New()
}

const (
	testQuery = `select
	    key_words,
	    views,
	    clicks,
	from
	    table`
)

func TestRun(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() {
		err := db.Close()
		require.NoError(t, err)
	}()

	rows := initRows()

	mock.ExpectQuery(testQuery).
		WillReturnRows(rows)

	err = Run(db)
	require.Error(t, err)
}

func initRows() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"arg1", "arg2", "arg3"})

	rows.AddRow("error_words", 1, 1)

	return rows
}
