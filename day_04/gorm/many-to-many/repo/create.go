package repo

import (
	"many-to-many/model"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func NewID(length ...int) (id string) {
	id, _ = gonanoid.New(8)
	return
}

func AddMemberToClub() (err error) {
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	//---- Tạo members
	john := model.Member{
		Id:   NewID(),
		Name: "John",
	}

	anna := model.Member{
		Id:   NewID(),
		Name: "Anna",
	}

	bob := model.Member{
		Id:   NewID(),
		Name: "Bob",
	}

	alice := model.Member{
		Id:   NewID(),
		Name: "Alice",
	}

	var members []model.Member
	members = append(members, john, anna, bob, alice)

	// Insert các bản ghi member vào trong CSDL
	if err := tx.Create(&members).Error; err != nil {
		tx.Rollback()
		return err
	}

	//--- Club
	math := model.Club{
		Id:   NewID(),
		Name: "Math",
	}

	sport := model.Club{
		Id:   NewID(),
		Name: "Sport",
	}

	music := model.Club{
		Id:   NewID(),
		Name: "Music",
	}
	var clubs []model.Club
	clubs = append(clubs, math, sport, music)

	// Insert các bản ghi club vào trong CSDL
	if err := tx.Create(&clubs).Error; err != nil {
		tx.Rollback()
		return err
	}

	//---- Thêm các thành viên vào club
	err = assignMembersToClub(tx, math, []model.Member{john, anna})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = assignMembersToClub(tx, sport, []model.Member{bob, alice})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = assignMembersToClub(tx, music, []model.Member{john, bob, alice})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Function thực hiện gắn các bản ghi member vào 1 club cụ thể
func assignMembersToClub(tx *gorm.DB, club model.Club, members []model.Member) (err error) {
	for _, member := range members {
		err := tx.Create(&model.MemberClub{
			MemberId: member.Id,
			ClubId:   club.Id,
			Active:   random.Intn(2) == 1, //random true or false
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
