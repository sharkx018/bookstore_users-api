package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	Status      string `json:"status"`
}

func (user Users) Marshal(isPublic bool) []interface{} {
	results := make([]interface{}, len(user))
	for idx, user := range user {
		results[idx] = user.Marshal(isPublic)
	}
	return results
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser

}
