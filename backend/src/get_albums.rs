use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};

use crate::albums::Album;



#[derive(Serialize)]
pub struct GetAlbumsResponse {
    albums: Vec<Album>,
    
}
#[get("/get_albums", format = "json")]
pub fn get_albums() -> Json<GetAlbumsResponse> {
    Json(GetAlbumsResponse {
        albums: vec![Album::new(
            vec!["Ellise"],
            "BAD EVIL",
            vec!["PRETTY", "EVIL"],
        )]
    })
}


