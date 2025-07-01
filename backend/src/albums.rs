use std::{collections::HashMap, sync::{Arc, RwLock}};

use crate::{common::expand_tilde, AppState};

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


pub fn update_albums(state: Arc<RwLock<AppState>>) -> std::io::Result<()> {
    //let mut singles = vec![("".to_string(), "Singles".to_string())];
    //let mut albums = HashMap::new();
    //albums.insert(("", "Singles"), vec![]);

    for dir in &state.read().unwrap().settings.folders {
        let dir = expand_tilde(dir).unwrap(); //safe unwrap i think?
        for entry in fs::read_dir(dir)? {
            let path = entry?.path();

            //// handle files first, implying singles
            //if path.is_file() {
            //    let path_str = path.into_os_string().to_str().unwrap().to_string(); //safe unwrap
            //    let songs = albums.get_mut(&("", "Singles")).unwrap(); //safe unwrap
            //    println!("{}", path_str);

            //    songs.push(path_str.to_string());
            //    continue;
            //} 

            // handle albums
            if path.is_dir() {
                let path_str = path.to_str().unwrap().to_string();
                for sub_entry in fs::read_dir(path)? {
                    let sub_path = sub_entry?.path();

                    if sub_path.is_file() {
                        println!("here1 {:?}", sub_path)
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
    match update_albums(state.inner().clone()) {
        Err(x) => {
            println!("caught error: {}", x);
        }
        Ok(_) => {
            println!("no error");
        }

    }
    Json(GetAlbumsResponse{albums: state.read().unwrap().albums.clone()})
}


