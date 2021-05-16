package relationships

import (
	"fmt"
	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"strconv"
	"testing"
)

func TestRepositoryImpl_GetRelationships(t *testing.T) {
	tcs := []struct {
		scenario         string
		fromID           int64
		toID             int64
		mockGetRelOutput *[]models.Relationship
		mockErr          error
		expResult        interface{}
		expErr           error
	}{
		{
			scenario: "success",
			fromID:   int64(1),
			toID:     int64(2),
			mockGetRelOutput: &[]models.Relationship{
				{
					ID:            1,
					FirstEmailID:  1,
					SecondEmailID: 2,
					Status:        "FRIEND",
				},
			},
			expResult: &[]models.Relationship{
				{
					ID:            1,
					FirstEmailID:  1,
					SecondEmailID: 2,
					Status:        "FRIEND",
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			rows := sqlmock.NewRows([]string{"id", "first_email_id", "second_email_id", "status"})
			for _, v := range *tc.mockGetRelOutput {
				rows.AddRow(v.ID, v.FirstEmailID, v.SecondEmailID, v.Status)
			}

			stmt := `select x.id, x.first_email_id, x.second_email_id, x.status
			from relationships x
			where x.first_email_id in (%s, %s)
			and x.second_email_id in (%s, %s);
			`
			query := fmt.Sprintf(
				stmt,
				strconv.FormatInt(tc.fromID, 10),
				strconv.FormatInt(tc.toID, 10),
				strconv.FormatInt(tc.fromID, 10),
				strconv.FormatInt(tc.toID, 10))
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			myDB := &RepositoryImpl{db}
			result, _ := myDB.GetRelationships(tc.fromID, tc.toID)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
			}
		})
	}
}

func TestRepositoryImpl_GetFriendsList(t *testing.T) {
	tcs := []struct {
		scenario                 string
		emailID                  int64
		mockGetFriendsListOutput *[]models.User
		mockErr                  error
		expResult                interface{}
		expErr                   error
	}{
		{
			scenario: "success",
			emailID:  int64(1),
			mockGetFriendsListOutput: &[]models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
			},
			expResult: &[]models.User{
				{
					ID:    1,
					Email: "a@gmail.com",
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			rows := sqlmock.NewRows([]string{"id", "email"})
			for _, v := range *tc.mockGetFriendsListOutput {
				rows.AddRow(v.ID, v.Email)
			}

			stmt := `select u.id, u.email
			from users u
        		join relationships r on r.second_email_id = u.id
			where r.first_email_id = %s and r.status = 'FRIEND'
			union
			select u.id, u.email
			from users u
        		join relationships r on r.first_email_id = u.id
			where r.second_email_id = %s and r.status = 'FRIEND';
			`
			query := fmt.Sprintf(
				stmt,
				strconv.FormatInt(tc.emailID, 10),
				strconv.FormatInt(tc.emailID, 10))
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			myDB := &RepositoryImpl{db}
			result, _ := myDB.GetFriendsList(tc.emailID)
			assert.Equal(t, tc.expErr, tc.mockErr)
			if tc.expErr == nil {
				assert.Equal(t, tc.expResult, result)
			}
		})
	}
}
