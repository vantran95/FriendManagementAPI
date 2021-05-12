package relationship

import (
	"InternalUserManagement/models"
	"errors"
)

// getRelationships get relationship between two emails.
func getRelationships(repo Repository, fromID, toID int64) ([]models.Relationship, error) {
	return repo.GetRelationships(fromID, toID)
}

// GetFriendsList attempts to retrieve a list of friends through a email.
func (s ServiceImpl) GetFriendsList(email string) ([]string, error) {
	var emails []string

	// Check email already created
	getUser, err := s.UserService.GetUser(email)
	if err != nil {
		return nil, err
	}

	// Get list friend
	getFriendsList, _ := s.Repository.GetFriendsList(getUser.ID)
	for _, f := range getFriendsList {
		friendEmail := f.Email
		emails = append(emails, friendEmail)
	}
	return emails, nil
}

// GetCommonFriends attempts to retrieve a list of common friends
func (s ServiceImpl) GetCommonFriends(firstEmail, secondEmail string) ([]string, error) {
	var commonEmails []string

	firstFriendsList, _ := s.GetFriendsList(firstEmail)
	secondFriendsList, _ := s.GetFriendsList(secondEmail)

	for _, v := range firstFriendsList {
		for _, item := range secondFriendsList {
			if item == v {
				commonEmails = append(commonEmails, v)
			}
		}
	}

	if len(commonEmails) == 0 {
		return nil, errors.New("do not have common friends between two emails")
	}

	return commonEmails, nil
}
