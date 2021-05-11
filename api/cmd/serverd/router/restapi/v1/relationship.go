package v1

import (
	"InternalUserManagement/models"
	"InternalUserManagement/pkg/dto"
	"InternalUserManagement/pkg/enum"
	"InternalUserManagement/pkg/exception"
	"InternalUserManagement/pkg/utils"
	"net/http"
)

// RelationshipService interface represents the criteria used to retrieve a relationship service
type RelationshipService interface {
	CreateRelationship(relationship models.Relationship) bool
	FindByTwoEmailIdsAndStatus(firstEmailId int64, secondEmailId int64, status []int64) ([]models.Relationship, error)
	FindByEmailIdAndStatus(emailId int64, status []int64) ([]models.Relationship, error)
	IsFriended(firstEmailId int64, secondEmailId int64) bool
}

type FriendService interface {
	CreateFriend(friendDto dto.FriendDto) (bool, *exception.Exception)
	GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *exception.Exception)
	GetCommonFriends(friendDto dto.FriendDto) ([]string, *exception.Exception)
}

// RelationshipImpl stores info to retrieve relationship
type RelationshipImpl struct {
	RelationshipService RelationshipService
	UserService         UserService
}

// CreateFriend attempts to create relationship between to email.
func (r RelationshipImpl) CreateFriend(friendDto dto.FriendDto) (bool, *exception.Exception) {
	firstEmail := friendDto.Friends[0]
	secondEmail := friendDto.Friends[1]

	if !utils.IsFormatEmail(firstEmail) || !utils.IsFormatEmail(secondEmail) {
		return false, &exception.Exception{Code: http.StatusBadRequest, Message: "Email invalid format"}
	}

	firstEmailId, err1 := r.UserService.FindUserIdByEmail(firstEmail)
	if err1 != nil {
		return false, &exception.Exception{Code: http.StatusNotFound, Message: err1.Error()}
	}

	secondEmailId, err2 := r.UserService.FindUserIdByEmail(secondEmail)
	if err2 != nil {
		return false, &exception.Exception{Code: http.StatusNotFound, Message: err2.Error()}
	}

	// Check friend
	if r.RelationshipService.IsFriended(firstEmailId, secondEmailId) {
		return false, &exception.Exception{Code: http.StatusInternalServerError, Message: "Two emails already friended"}
	}
	relationship := models.Relationship{FirstEmailId: firstEmailId, SecondEmailId: secondEmailId, Status: enum.FRIEND}
	result := r.RelationshipService.CreateRelationship(relationship)

	if result != true {
		return false, &exception.Exception{Code: http.StatusInternalServerError, Message: "Can not make friend"}
	}

	return true, nil
}

// GetFriendsListByEmail attempts to retrieve a list of friends
func (r RelationshipImpl) GetFriendsListByEmail(emailDto dto.EmailDto) ([]string, *exception.Exception) {
	emails := []string{}
	if !utils.IsFormatEmail(emailDto.Email) {
		return nil, &exception.Exception{Code: http.StatusBadRequest, Message: "Email invalid format"}
	}

	// Check email already created
	emailId, err := r.UserService.FindUserIdByEmail(emailDto.Email)
	if err != nil {
		return nil, &exception.Exception{Code: http.StatusNotFound, Message: err.Error()}
	}

	// Find list relationship of email by emailId and status
	relationships, _ := r.RelationshipService.FindByEmailIdAndStatus(emailId, []int64{enum.FRIEND})

	if relationships != nil {
		emailIds := getEmailIdsFromListRelationships(relationships)
		emailIds = utils.RemoveItemFromList(emailIds, emailId)

		// Get list emails by list emailIds
		if emailIds != nil && len(emailIds) > 0 {
			emails, _ = r.UserService.FindEmailByIds(emailIds)
		}
	}
	return emails, nil
}

// GetCommonFriends attempts to retrieve a list of common friends
func (r RelationshipImpl) GetCommonFriends(friendDto dto.FriendDto) ([]string, *exception.Exception) {
	commonEmails := []string{}
	var firstEmail dto.EmailDto
	firstEmail.Email = friendDto.Friends[0]
	var secondEmail dto.EmailDto
	secondEmail.Email = friendDto.Friends[1]

	firstEmailRelationships, _ := r.GetFriendsListByEmail(firstEmail)
	secondEmailRelationships, _ := r.GetFriendsListByEmail(secondEmail)

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
