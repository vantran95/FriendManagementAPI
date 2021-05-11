package relationship

import (
	"InternalUserManagement/models"
	"fmt"
	"strconv"
	"strings"
)

// CreateRelationship attempts to create relationship between 2 email addresses.
func (r RepositoryImpl) CreateRelationship(relationship models.Relationship) (bool, error) {
	query, err := r.DB.Prepare("insert into relationship (first_email_id, second_email_id, status) values ($1, $2, $3);")

	if err != nil {
		return false, err
	}

	query.Exec(relationship.FirstEmailId, relationship.SecondEmailId, relationship.Status)
	return true, nil
}

// FindByTwoEmailIdsAndStatus attempts to retrieve a friend relationship by two email addresses id and status
func (r RepositoryImpl) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error) {
	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.*
			from relationship x
			where x.first_email_id in (%s, %s)
			and x.second_email_id in (%s, %s)
			and x.status in (%s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strings.Join(strStatusIds, ","))

	results, err := r.DB.Query(query)
	if err != nil {
		return []models.Relationship{}, err
	}

	var relationships []models.Relationship
	for results.Next() {
		var id, firstEmailId, secondEmailId, status int64
		results.Scan(&id, &firstEmailId, &secondEmailId, &status)
		relationship := models.Relationship{Id: id, FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: status}
		relationships = append(relationships, relationship)
	}
	return relationships, nil
}

// FindByEmailIdAndStatus attempts to retrieve a friend relationship by email id and status.
func (r RepositoryImpl) FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error) {
	strStatusIds := make([]string, len(status))

	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	qr := `select x.*
			from relationship x
			where (x.first_email_id = %s
			or x.second_email_id = %s) 
			and x.status in (%s);
			`
	query := fmt.Sprintf(
		qr,
		strconv.FormatInt(emailId, 10),
		strconv.FormatInt(emailId, 10),
		strings.Join(strStatusIds, ","))

	results, err := r.DB.Query(query)

	if err != nil {
		return []models.Relationship{}, err
	}

	var relationships []models.Relationship

	for results.Next() {
		var id, firstEmailId, secondEmailId, status int64

		results.Scan(&id, &firstEmailId, &secondEmailId, &status)

		relationship := models.Relationship{Id: id, FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: status}

		relationships = append(relationships, relationship)
	}

	return relationships, nil
}
