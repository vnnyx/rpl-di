package entity

type VideoEntity struct {
	Id         int
	PlaylistId PlaylistEntity
	ExamId     ExamEntity
	Title      string
	URLVideo   string
	Deskripsi  string
}
