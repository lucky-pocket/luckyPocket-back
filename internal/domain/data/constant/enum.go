package constant

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

func (g Gender) Valid() bool {
	switch g {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (Gender) Values() (genders []string) {
	for _, g := range []Gender{GenderMale, GenderFemale} {
		genders = append(genders, string(g))
	}
	return
}

type Role string

const (
	RoleMember Role = "ROLE_MEMBER"
	RoleAdmin  Role = "ROLE_ADMIN"
)

func (r Role) Valid() bool {
	switch r {
	case RoleMember, RoleAdmin:
		return true
	}
	return false
}

func (Role) Values() (roles []string) {
	for _, r := range []Role{RoleAdmin, RoleMember} {
		roles = append(roles, string(r))
	}
	return
}

type UserType string

const (
	TypeStudent  UserType = "STUDENT"
	TypeTeacher  UserType = "TEACHER"
	TypeGraduate UserType = "GRADUATE"
)

func (t UserType) Valid() bool {
	switch t {
	case TypeStudent, TypeTeacher, TypeGraduate:
		return true
	}
	return false
}

func (UserType) Values() (userTypes []string) {
	for _, ut := range []UserType{TypeTeacher, TypeGraduate, TypeStudent} {
		userTypes = append(userTypes, string(ut))
	}
	return
}

type NoticeType string

const (
	NoticeTypeRevealed NoticeType = "REVEALED"
	NoticeTypeReceived NoticeType = "RECEIVED"
)

func (n NoticeType) Valid() bool {
	switch n {
	case NoticeTypeReceived, NoticeTypeRevealed:
		return true
	}
	return false
}

func (NoticeType) Values() (noticeTypes []string) {
	for _, nt := range []NoticeType{NoticeTypeReceived, NoticeTypeRevealed} {
		noticeTypes = append(noticeTypes, string(nt))
	}
	return
}

type SortType string

const (
	SortTypePocket SortType = "POCKET"
	SortTypeCoins  SortType = "COIN"
)

func (s SortType) Valid() bool {
	switch s {
	case SortTypePocket, SortTypeCoins:
		return true
	}
	return false
}
