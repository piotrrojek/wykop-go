package wykop

type Profile struct {
	Login             string `json:"login"`
	RegistrationEmail string `json:"email"`
	Email             string `json:"public_email"`
	Name              string `json:"name"`
	WWW               string `json:"www"`
	Jabber            string `json:"jabber"`
	GG                string `json:"gg"`
	City              string `json:"city"`
	About             string `json:"about"`
	AuthorGroup       int    `json:"author_group"`
	LinksAdded        int    `json:"links_added"`
	LinksPublished    int    `json:"links_published"`
	Comments          int    `json:"comments"`
	Rank              int    `json:"rank"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	Entries           int    `json:"entries"`
	EntriesComments   int    `json:"entries_comments"`
	Diggs             int    `json:"diggs"`
	Buries            int    `json:"buries"`
	Groups            int    `json:"groups"`
	RelatedLinks      int    `json:"related_links"`
	SignupDate        string `json:"signup_date"`
	Avatar            string `json:"avatar"`
	AvatarBig         string `json:"avatar_big"`
	AvatarMed         string `json:"avatar_med"`
	AvatarLo          string `json:"avatar_lo"`
	IsObserved        bool   `json:"is_observed"`
	Sex               string `json:"sex"`
	URL               string `json:"url"`
	ViolationURL      string `json:"violation_url"`
}

type Comment struct {
	Body            string `json:"body,omitempty"`
	Status          string `json:"status,omitempty"`
	Type            string `json:"type,omitempty"`
	App             string `json:"app,omitempty"`
	Date            string `json:"date,omitempty"`
	ID              int    `json:"id,omitempty"`
	ParentID        int    `json:"parent_id,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Blocked         bool   `json:"blocked,omitempty"`
	UserFavorite    bool   `json:"user_favorite,omitempty"`
	CanVote         bool   `json:"can_vote,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	VoteCountPlus   int    `json:"vote_count_plus,omitempty"`
	VoteCountMinus  int    `json:"vote_count_minus,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	Link            Link   `json:"link,omitempty"`
	Source          string `json:"source,omitempty"`
	Embed           Embed  `json:"embed,omitempty"`
}

type RelatedLink struct {
	Title           string `json:"title,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	URL             string `json:"url,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	UserVote        int    `json:"user_vote,omitempty"`
	Link            Link   `json:"link,omitempty"`
	Type            string `json:"type,omitempty"`
	ID              int    `json:"id,omitempty"`
	Plus18          bool   `json:"plus18,omitempty"`
}

type Entry struct {
	App               string         `json:"app,omitempty"`
	ID                int            `json:"id,omitempty"`
	Author            string         `json:"author,omitempty"`
	AuthorAvatar      string         `json:"author_avatar,omitempty"`
	AuthorAvatarBig   string         `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed   string         `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo    string         `json:"author_avatar_lo,omitempty"`
	AuthorGroup       int            `json:"author_group,omitempty"`
	Date              string         `json:"date,omitempty"`
	Body              string         `json:"body,omitempty"`
	URL               string         `json:"url,omitempty"`
	Receiver          string         `json:"receiver,omitempty"`
	ReceiverSex       string         `json:"receiver_sex,omitempty"`
	ReceiverAvatar    string         `json:"receiver_avatar,omitempty"`
	ReceiverAvatarMed string         `json:"receiver_avatar_med,omitempty"`
	ReceiverAvatarLo  string         `json:"receiver_avatar_lo,omitempty"`
	ReceiverAvatarBig string         `json:"receiver_avatar_big,omitempty"`
	ReceiverGroup     int            `json:"receiver_group,omitempty"`
	Comments          []EntryComment `json:"comments,omitempty"`
	CommentCount      int            `json:"comment_count"`
	VoteCount         int            `json:"vote_count,omitempty"`
	UserVote          int            `json:"user_vote,omitempty"`
	Voters            []Dig          `json:"voters,omitempty"`
	UserFavorite      bool           `json:"user_favorite,omitempty"`
	Embed             Embed          `json:"embed,omitempty"`
	Blocked           bool           `json:"blocked,omitempty"`
	Type              string         `json:"type,omitempty"`
	Deleted           bool           `json:"deleted,omitempty"`
	AuthorSex         string         `json:"author_sex"`
	ViolationURL      string         `json:"violation_url,omitempty"`
	Source            string         `json:"source"`
	CanComment        bool           `json:"can_comment"`
}

type EntryComment struct {
	ID              int    `json:"id,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Date            string `json:"date,omitempty"`
	Body            string `json:"body,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	UserVote        int    `json:"user_vote,omitempty"`
	Voters          []Dig  `json:"voters,omitempty"`
	Embed           Embed  `json:"embed,omitempty"`
	Entry           Entry  `json:"entry,omitempty"`
	ViolationUrl    string `json:"violation_url,omitempty"`
	App             string `json:"app,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	Source          string `json:"source,omitempty"`
	Type            string `json:"entry_comment,omitempty"`
	Blocked         bool   `json:"blocked,omitempty"`
}

type Embed struct {
	Type    string `json:"type,omitempty"`
	Preview string `json:"preview,omitempty"`
	URL     string `json:"url,omitempty"`
	Source  string `json:"source,omitempty"`
	Plus18  bool   `json:"plus18,omitempty"`
}

type Dig struct {
	Author          string `json:"author,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	Date            string `json:"date,omitempty"`
}

type Link struct {
	App             string `json:"app,omitempty"`
	ID              int    `json:"id,omitempty"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	Tags            string `json:"tags,omitempty"`
	URL             string `json:"url,omitempty"`
	SourceURL       string `json:"source_url,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	CommentCount    int    `json:"comment_count,omitempty"`
	ReportCount     int    `json:"report_count,omitempty"`
	Date            string `json:"date,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Preview         string `json:"preview,omitempty"`
	UserLists       []int  `json:"user_lists,omitempty"`
	Plus18          bool   `json:"plus18,omitempty"`
	Status          string `json:"status,omitempty"`
	CanVote         bool   `json:"can_vote,omitempty"`
	HasOwnContent   bool   `json:"has_own_content,omitempty"`
	CategoryName    string `json:"category_name,omitempty"`
	UserVote        bool   `json:"user_vote,omitempty"`
	UserObserve     bool   `json:"user_observe,omitempty"`
	UserFavorite    bool   `json:"user_favorite,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	Info            string `json:"info,omitempty"`
	RelatedCount    int    `json:"related_count,omitempty"`
	IsHot           bool   `json:"is_hot,omitempty"`
	Type            string `json:"type,omitempty"`
}
