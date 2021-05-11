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
}

type FriendService interface {
	CreateFriend(friendDto dto.FriendDto) (bool, *exception.Exception)
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

	relationship := models.Relationship{FirstEmailId: firstEmailId, SecondEmailID: secondEmailId, Status: enum.FRIEND}
	result := r.RelationshipService.CreateRelationship(relationship)

	if result != true {
		return false, &exception.Exception{Code: http.StatusInternalServerError, Message: "Can not make friend"}
	}

	return true, nil
}
