package repo

import (
	"many-to-many/model"
)

// Tìm club theo tên, đồng thời lấy danh sách thành viên
func GetClubByName(name string) (club model.Club, err error) {
	err = DB.Preload("Members").Find(&club, "name = ?", name).Error
	if err != nil {
		return model.Club{}, err
	}
	return club, nil
}

// Tìm thành viên, lấy danh sách club mà người đó tham gia
func GetMemberByName(name string) (member model.Member, err error) {
	err = DB.Preload("Clubs").Find(&member, "name = ?", name).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}

// Xem 2 thành viên có chung club nào
func FindShareClubOfMember(name string) (member model.Member, err error) {
	err = DB.Preload("Clubs").Find(&member, "name = ?", name).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}
