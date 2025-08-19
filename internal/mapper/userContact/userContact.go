package mapper

import (
	"github.com/google/uuid"
	"github.com/habit-tracker-api/entities"
	"github.com/habit-tracker-api/enums"

	_userContactModel "github.com/habit-tracker-api/model/userContact"
)

func ToUserContactCreateEntity(userID uuid.UUID, userContactReq []*_userContactModel.UserContactReq) ([]entities.UserContact, []uuid.UUID, error) {
	var userContact []entities.UserContact
	var userContactId []uuid.UUID

	for _, h := range userContactReq {
		if h.ID == uuid.Nil {
			userContactID := uuid.New()
			userContactEntity := entities.UserContact{
				ID:          userContactID,
				UserID:      userID,
				ContactType: enums.ContactType(h.ContactType),
				Value:       h.Value,
			}
			userContactId = append(userContactId, userContactID)
			userContact = append(userContact, userContactEntity)
		}
	}

	return userContact, userContactId, nil
}

func ToUserContactUpdateEntity(userID uuid.UUID, userContactReq []*_userContactModel.UserContactReq) ([]entities.UserContact, []uuid.UUID, error) {
	var userContact []entities.UserContact
	var userContactId []uuid.UUID

	for _, h := range userContactReq {
		if h.ID != uuid.Nil && !h.IsDelete {
			userContactEntity := entities.UserContact{
				ID:          h.ID,
				UserID:      userID,
				ContactType: enums.ContactType(h.ContactType),
				Value:       h.Value,
			}
			userContactId = append(userContactId, h.ID)
			userContact = append(userContact, userContactEntity)
		}
	}

	return userContact, userContactId, nil
}

func ToUserContactDeleteEntity(userID uuid.UUID, userContactReq []*_userContactModel.UserContactReq) ([]entities.UserContact, []uuid.UUID, error) {
	var userContact []entities.UserContact
	var userContactId []uuid.UUID

	for _, h := range userContactReq {
		if h.ID != uuid.Nil && h.IsDelete {
			userContactEntity := entities.UserContact{
				ID:     h.ID,
				UserID: userID,
			}
			userContactId = append(userContactId, h.ID)
			userContact = append(userContact, userContactEntity)
		}
	}

	return userContact, userContactId, nil
}

func ToUserContactRes(userContacts []*entities.UserContact) []*_userContactModel.UserContactRes {
	var userContactModels []*_userContactModel.UserContactRes
	for _, h := range userContacts {
		userContactRes := &_userContactModel.UserContactRes{
			ID:          h.ID,
			ContactType: string(h.ContactType),
			Value:       h.Value,
		}
		userContactModels = append(userContactModels, userContactRes)
	}

	return userContactModels
}
