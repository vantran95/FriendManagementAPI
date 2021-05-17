package relationships

import (
	"errors"

	"github.com/s3corp-github/S3_FriendManagement_VanTran/api/internal/models"
)

type userServiceRetriever interface {
	GetUser(email string) (*models.User, error)
}

type retrieveRepository interface {
	GetRelationships(fromID, toID int64) (*[]models.Relationship, error)
	GetFriendsList(emailID int64) (*[]models.User, error)
}

// GetFriendsList attempts to retrieve a list of friends through a email.
func (s ServiceImpl) GetFriendsList(email string) ([]string, error) {
	var emails []string

	// Check email already created
	getUser, _ := s.UserServiceRetriever.GetUser(email)

	if getUser == nil {
		return []string{}, errors.New("user does not exist")
	}

	// Get list friend
	getFriendsList, _ := s.RetrieveRepo.GetFriendsList(getUser.ID)
	if getFriendsList == nil {
		return []string{}, errors.New("user does not have friend")
	}
	for _, f := range *getFriendsList {
		friendEmail := f.Email
		emails = append(emails, friendEmail)
	}
	return emails, nil
}

// GetCommonFriends attempts to retrieve a list of common friends
func (s ServiceImpl) GetCommonFriends(firstEmail, secondEmail string) ([]string, error) {
	var commonEmails []string

	firstFriendsList, err := s.GetFriendsList(firstEmail)
	if err != nil {
		return []string{}, err
	}
	secondFriendsList, err := s.GetFriendsList(secondEmail)
	if err != nil {
		return []string{}, err
	}

	for _, v := range firstFriendsList {
		for _, item := range secondFriendsList {
			if item == v {
				commonEmails = append(commonEmails, v)
			}
		}
	}

	//if len(commonEmails) == 0 {
	//	return nil, errors.New("do not have common friends between two emails")
	//}

	return commonEmails, nil
}

// getRelationships get relationship between two emails.
func getRelationships(repo retrieveRepository, fromID, toID int64) (*[]models.Relationship, error) {
	return repo.GetRelationships(fromID, toID)
}
