package common

import "time"

type PSModel struct {
	Id        int        `json:"-"  gorm:"column:id;"`
	FakeID    *UID       `json:"id"  gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"create_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"update_at"`
}

func (m *PSModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeID = &uid
}
