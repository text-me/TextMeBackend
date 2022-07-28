package models

import (
	"encoding/json"
	"github.com/text-me/TextMeBackend/log"
)

type GroupJson struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func AddGroup(title string) (*Group, error) {
	insert := &Group{Title: title}
	return insert, db.Create(insert).Error
}

func SelectGroups() []Group {
	var groups []Group
	db.Find(&groups)

	return groups
}

func GroupsToJson(groups []Group) ([]byte, error) {
	list := make([]GroupJson, 0)
	for _, msg := range groups {
		list = append(list, GroupJson{
			Id:    msg.ID,
			Title: msg.Title,
		})
	}

	groupsJson, err := json.Marshal(list)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return groupsJson, nil
}

func (m Group) ToJson() ([]byte, error) {
	newMessageJson := &GroupJson{
		Id:    m.ID,
		Title: m.Title,
	}

	groupJson, err := json.Marshal(newMessageJson)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return groupJson, nil
}
