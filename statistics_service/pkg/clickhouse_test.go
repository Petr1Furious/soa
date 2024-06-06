package statistics_listener

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSaveView(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	c := &Clickhouse{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM views WHERE user_id = \? AND post_id = \? AND author_id = \?`).WithArgs(1, 2, 3).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO views \(user_id, post_id, author_id\) VALUES \(\?, \?, \?\)`).WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = c.SaveView(1, 2, 3)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveLike(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	c := &Clickhouse{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM likes WHERE user_id = \? AND post_id = \? AND author_id = \?`).WithArgs(1, 2, 3).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO likes \(user_id, post_id, author_id\) VALUES \(\?, \?, \?\)`).WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = c.SaveLike(1, 2, 3)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
