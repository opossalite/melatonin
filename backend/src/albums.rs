use std::{collections::HashMap, sync::{Arc, RwLock}};

use crate::AppState;

use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};
//use tokio::fs;
use std::fs;



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


pub async fn update_albums(state: Arc<RwLock<AppState>>) -> std::io::Result<()> {
    //let mut singles = vec![("".to_string(), "Singles".to_string())];
    let mut albums = HashMap::new();
    albums.insert(("", "Singles"), vec![]);

    for dir in &state.read().unwrap().settings.folders {
        for entry in fs::read_dir(dir)? {
            let path = entry?.path();

            // handle files first, implying singles
            if path.is_file() {
                match path.into_os_string().to_str() {
                    Some(file) => {
                        let songs = albums.get_mut(&("", "Singles")).unwrap();
                        todo!();
                        // want to break the file name into artists and title, maybe via metadata
                        songs.push(file.to_string());
                    }
                    None => unreachable!()
                }
                continue;
            } 

            // handle albums now
            if path.is_dir() {
                for sub_entry in fs::read_dir(path)? {
                    let sub_path = sub_entry?.path();
                    match sub_path.into_os_string().to_str() {
                        Some(file) => {
                            // break into artist name and album name, maybe via internal metadata
                            todo!();
                        }
                        None => unreachable!()
                    }
                    

                    if sub_path.is_file() {

                    }

                }
            }
        }
    }

    Ok(())
}


#[derive(Serialize)]
pub struct GetAlbumsResponse {
    albums: Vec<Album>,
    
}
#[get("/get_albums", format = "json")]
pub fn get_albums(state: &rocket::State<Arc<RwLock<AppState>>>) -> Json<GetAlbumsResponse> {
    Json(GetAlbumsResponse{albums: state.read().unwrap().albums.clone()})
}


