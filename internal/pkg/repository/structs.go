package repository

type Post struct {
	ID         int64     `db:"id"          json:"id"`
	Heading    string    `db:"heading"     json:"heading"`
	Text       string    `db:"text"        json:"text"`
	LikesCount int       `db:"likes_count" json:"likes_count"`
	Comments   []Comment `                 json:"comments"`
}

type Comment struct {
	ID         int64  `db:"id"          json:"id"`
	PostID     int64  `db:"post_id"     json:"-"`
	Text       string `db:"text"        json:"text"`
	LikesCount int    `db:"likes_count" json:"likes_count"`
}
