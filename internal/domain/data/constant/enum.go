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
