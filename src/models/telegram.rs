mod models {
    use serde::{Deserialize, Serialize};

    #[derive(Serialize, Deserialize, Debug)]
    pub struct Channel {
        pub id: i64,
        pub title: String,
        pub username: String,
        pub access_hash: i64,
        pub date: i32,
        pub version: i32,
        pub admin_rights: Option<AdminRights>,
        pub banned_rights: Option<BannedRights>,
        pub participants_count: i32,
        pub about: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct AdminRights {
        pub change_info: bool,
        pub post_messages: bool,
        pub edit_messages: bool,
        pub delete_messages: bool,
        pub ban_users: bool,
        pub invite_users: bool,
        pub pin_messages: bool,
        pub add_admins: bool,
        pub manage_call: bool,
        pub other_admin_rights: bool,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct BannedRights {
        pub view_messages: bool,
        pub send_messages: bool,
        pub send_media: bool,
        pub send_stickers: bool,
        pub send_gifs: bool,
        pub send_games: bool,
        pub send_inline: bool,
        pub embed_links: bool,
        pub send_polls: bool,
        pub change_info: bool,
        pub invite_users: bool,
        pub pin_messages: bool,
    }
}