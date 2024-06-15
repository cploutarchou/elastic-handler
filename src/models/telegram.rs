mod models {
    use serde::{Deserialize, Serialize};
    use serde::{Deserialize, Serialize};
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
    #[derive(Serialize, Deserialize, Debug)]
    pub struct Token {
        pub token_address: String,
        pub balance: f64,
        pub decimals: u8,
        pub ui_amount_string: String,
        pub token_name: String,
        pub token_symbol: String,
        pub metadata_source: String,
        pub token_type: String,
        pub market_overview: MarketOverview,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct Metadata {
        pub key: String,
        pub update_authority: String,
        pub mint: String,
        pub name: String,
        pub symbol: String,
        pub uri: String,
        pub seller_fee_basis_points: String,
        pub primary_sale_happened: String,
        pub is_mutable: String,
        pub edition_nonce: String,
        pub token_standard: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct MarketOverview {
        pub price: f64,
        pub market_cap: f64,
        pub current_supply: f64,
    }

    #[derive(Serialize, Deserialize, Debug)]
    pub struct NftData {
        pub metadata: Metadata,
        pub market_overview: Option<MarketOverview>,
    }
}
