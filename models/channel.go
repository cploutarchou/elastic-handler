package models

// Channel represents a Telegram channel with its various attributes.
type Channel struct {
	ID                int64         `json:"id"`                 // Unique identifier for the channel
	Title             string        `json:"title"`              // Title of the channel
	Username          string        `json:"username"`           // Username of the channel (if any)
	AccessHash        int64         `json:"access_hash"`        // Access hash for the channel
	Date              int32         `json:"date"`               // Creation date of the channel
	Version           int32         `json:"version"`            // Version of the channel
	AdminRights       *AdminRights  `json:"admin_rights"`       // Admin rights of the channel
	BannedRights      *BannedRights `json:"banned_rights"`      // Banned rights of the channel
	ParticipantsCount int32         `json:"participants_count"` // Number of participants in the channel
	About             string        `json:"about"`              // Description of the channel
}

// AdminRights represents the admin rights for a channel.
type AdminRights struct {
	ChangeInfo       bool `json:"change_info"`
	PostMessages     bool `json:"post_messages"`
	EditMessages     bool `json:"edit_messages"`
	DeleteMessages   bool `json:"delete_messages"`
	BanUsers         bool `json:"ban_users"`
	InviteUsers      bool `json:"invite_users"`
	PinMessages      bool `json:"pin_messages"`
	AddAdmins        bool `json:"add_admins"`
	ManageCall       bool `json:"manage_call"`
	OtherAdminRights bool `json:"other_admin_rights"`
}

// BannedRights represents the banned rights for a channel.
type BannedRights struct {
	ViewMessages bool `json:"view_messages"`
	SendMessages bool `json:"send_messages"`
	SendMedia    bool `json:"send_media"`
	SendStickers bool `json:"send_stickers"`
	SendGifs     bool `json:"send_gifs"`
	SendGames    bool `json:"send_games"`
	SendInline   bool `json:"send_inline"`
	EmbedLinks   bool `json:"embed_links"`
	SendPolls    bool `json:"send_polls"`
	ChangeInfo   bool `json:"change_info"`
	InviteUsers  bool `json:"invite_users"`
	PinMessages  bool `json:"pin_messages"`
}
