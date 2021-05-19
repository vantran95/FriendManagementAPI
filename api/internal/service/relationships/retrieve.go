package relationships

import (
	"database/sql"
	"errors"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

type (
	// userRetrieverRepo interface represents the retriever from user repository
	userRetrieverRepo interface {
		GetUser(email string) (*models.User, error)
	}
	// retrieveRepository interface represents the retrieve from relationship repository
	retrieveRepository interface {
		GetRelationships(requestID, targetID int64) (*[]models.Relationship, error)
		GetFriendsList(emailID int64) (*[]models.User, error)
	}
)

// GetFriendsList attempts to retrieve a list of friends through a email.
func (s ServiceImpl) GetFriendsList(email string) ([]string, error) {
	var emails []string
	// Check email already created
	user, err := s.UserRetrieverRepo.GetUser(email)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return []string{}, errors.New("user does not exists")
		case err != nil:
			return []string{}, err
		}
	}
	// Get list friend
	friendsList, err := s.RetrieveRepo.GetFriendsList(user.ID)
	if err != nil {
		return []string{}, err
	}
	// check length
	if len(*friendsList) == 0 {
		return []string{}, errors.New("user does not have friend")
	}
	for _, f := range *friendsList {
		friendEmail := f.Email
		emails = append(emails, friendEmail)
	}
	return emails, nil
}

// GetCommonFriends attempts to retrieve a list of common friends
func (s ServiceImpl) GetCommonFriends(requestEmail, targetEmail string) ([]string, error) {
	var commonEmails []string
	requestFriendsList, err := s.GetFriendsList(requestEmail)
	if err != nil {
		return []string{}, err
	}
	targetFriendsList, err := s.GetFriendsList(targetEmail)
	if err != nil {
		return []string{}, err
	}
	for _, v := range requestFriendsList {
		for _, item := range targetFriendsList {
			if item == v {
				commonEmails = append(commonEmails, v)
			}
		}
	}
	if len(commonEmails) == 0 {
		return []string{}, errors.New("do not have common friends between two emails")
	}
	return commonEmails, nil
}

// getRelationships get relationship between two emails.
func getRelationships(repo retrieveRepository, requestID, targetID int64) (*[]models.Relationship, error) {
	return repo.GetRelationships(requestID, targetID)
}
