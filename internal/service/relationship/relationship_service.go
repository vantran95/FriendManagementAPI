package relationship

import (
	"InternalUserManagement/models"
	"InternalUserManagement/pkg/enum"
	"InternalUserManagement/pkg/exception"
	"InternalUserManagement/pkg/utils"
	"net/http"
)

// FindByTwoEmailIdsAndStatus attempts to retrieve a list friend relationship by two emails and status
func (s ServiceImpl) FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error) {
	return s.Repository.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status)
}

// FindByEmailIdAndStatus attempts to retrieve a friend relationship by email address and status
func (s ServiceImpl) FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error) {
	return s.Repository.FindByEmailIdAndStatus(emailId, status)
}

// getRelationship get relationship between two emails.
func getRelationship(repo Repository, fromID, toID int64) ([]models.Relationship, error) {
	return repo.GetRelationship(fromID, toID)
}

func (s ServiceImpl) MakeFriend(firstEmail, secondEmail string) (bool, *exception.Exception) {
	firstEmailId, err := s.UserService.FindUserIdByEmail(firstEmail)
	if err != nil {
		return false, &exception.Exception{Code: http.StatusNotFound, Message: err.Error()}
	}

	secondEmailId, err2 := s.UserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return false, &exception.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// Get relationship and check friend
	rs, err := getRelationship(s.Repository, firstEmailId, secondEmailId)
	if err != nil {
		return false, &exception.Exception{Code: http.StatusNotFound, Message: err.Error()}
	}

	for _, item := range rs {
		if item.Status == enum.FRIEND {
			return false, &exception.Exception{Code: http.StatusInternalServerError, Message: "Two emails already friended"}
		}
	}

	relationship := models.Relationship{FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: enum.FRIEND}
	result, _ := s.Repository.CreateRelationship(relationship)

	if result != true {
		return false, &exception.Exception{Code: http.StatusInternalServerError, Message: "Can not make friend"}
	}

	return true, nil
}

// GetFriendsListByEmail attempts to retrieve a list of friends
func (s ServiceImpl) GetFriendsListByEmail(email string) ([]string, *exception.Exception) {
	emails := []string{}

	// Check email already created
	emailId, err := s.UserService.FindUserIdByEmail(email)
	if err != nil {
		return nil, &exception.Exception{Code: http.StatusNotFound, Message: err.Error()}
	}

	// Find list relationship of email by emailId and status
	relationships, _ := s.FindByEmailIdAndStatus(emailId, []int64{enum.FRIEND})

	if relationships != nil {
		emailIds := getEmailIdsFromListRelationships(relationships)
		emailIds = utils.RemoveItemFromList(emailIds, emailId)

		// Get list emails by list emailIds
		if emailIds != nil && len(emailIds) > 0 {
			emails, _ = s.UserService.FindEmailByIds(emailIds)
		}
	}
	return emails, nil
}

// GetCommonFriends attempts to retrieve a list of common friends
func (s ServiceImpl) GetCommonFriends(firstEmail, secondEmail string) ([]string, *exception.Exception) {
	commonEmails := []string{}

	firstEmailRelationships, _ := s.GetFriendsListByEmail(firstEmail)
	secondEmailRelationships, _ := s.GetFriendsListByEmail(secondEmail)

	for _, v := range firstEmailRelationships {
		if getEmailExists(secondEmailRelationships, v) {
			commonEmails = append(commonEmails, v)
		}
	}

	if len(commonEmails) == 0 {
		return nil, &exception.Exception{Code: http.StatusNotFound, Message: "Do not have common friends between two emails"}
	}

	return commonEmails, nil
}

// getEmailExists check email exist on slice
func getEmailExists(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// getEmailIdsFromListRelationships attempts to retrieve a list email ids from list relationships.
func getEmailIdsFromListRelationships(relationships []models.Relationship) []int64 {
	keys := make(map[int64]bool)
	set := []int64{}
	if relationships != nil && len(relationships) > 0 {
		for _, rela := range relationships {
			if _, ok := keys[rela.FirstEmailId]; !ok {
				keys[rela.FirstEmailId] = true
				set = append(set, rela.FirstEmailId)
			}

			if _, ok := keys[rela.SecondEmailId]; !ok {
				keys[rela.SecondEmailId] = true
				set = append(set, rela.SecondEmailId)
			}
		}
	}
	return set
}
