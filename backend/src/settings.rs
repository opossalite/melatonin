use crate::AppState;
use rocket::serde::json::Json;
use serde::Serialize;



#[derive(Serialize, Debug)]
pub struct Settings {
    pub folders: Vec<String>,
    pub main_color: (u8, u8, u8),
}
impl Settings {
    pub fn new_default() -> Self {
        Settings {
            folders: vec!["~/Music".to_string()],
            main_color: (215, 142, 30), //default: #d78e1e or (215, 142, 30)
        }
    }

    pub fn read_settings() -> Self {
        todo!()
    }

    pub fn update_settings() {
        todo!()
    }
}


#[derive(Serialize)]
pub struct GetSettingsResponse {
    settings: Settings,
    
}
#[get("/get_settings", format = "json")]
pub fn get_settings(state: &rocket::State<AppState>) -> Json<GetSettingsResponse> {
    Json(GetSettingsResponse {
        settings: Settings::new_default(),
    })
}




// todo: save and retrieve settings from ~/.config/melatonin/settings.json



