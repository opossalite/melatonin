use crate::AppState;

use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};



#[derive(Serialize, Debug, Clone)]
pub struct Album {
    artists: Vec<String>, //generally just one artist but with flexibility for more
    title: String,
    songs: Vec<String>, //songs are ordered correctly here.
}
impl Album {
    pub fn new(artists: Vec<&str>, title: &str, songs: Vec<&str>) -> Self {
        Album {
            artists: artists.iter().map(|x| String::from(*x)).collect(),
            title: String::from(title),
            songs: songs.iter().map(|x| String::from(*x)).collect(),
        }
    }
}


#[derive(Serialize)]
pub struct GetAlbumsResponse {
    albums: Vec<Album>,
    
}
#[get("/get_albums", format = "json")]
pub fn get_albums(state: &rocket::State<AppState>) -> Json<GetAlbumsResponse> {
    Json(GetAlbumsResponse{albums: state.albums.clone()})
}


