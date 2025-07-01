mod history;
mod albums;
mod experimentation;
mod settings;

use std::sync::{Arc, RwLock};

use albums::{get_albums, Album};
use experimentation::{roll, upper};
use history::History;
use settings::{get_settings, Settings};

#[macro_use] extern crate rocket;
use rocket_cors::CorsOptions;



#[derive(Debug)]
pub struct AppState {
    pub albums: Vec<Album>,
    pub settings: Settings,
    pub history: History,
}
//type SharedState = Arc<RwLock<AppState>>; //optional for multithreading



#[launch]
fn rocket() -> _ {
    // just ensures the frontend is happy
    let cors = CorsOptions::default()
        .to_cors()
        .expect("error building CORS");

    // establish persistent state in the backend
    let state = AppState {
        albums: vec![Album::new( //TEMPORARY
            vec!["Ellise"],
            "BAD EVIL",
            vec!["PRETTY", "EVIL"],
        )],
        settings: Settings::new_default(),
        history: History::new(), //TODO make new_default
    };

    rocket::build()
        .configure(rocket::Config::figment().merge(("port", 8800)))
        .manage(Arc::new(RwLock::new(state)))
        .mount("/", routes![roll, upper, get_albums, get_settings])
        .attach(cors)
}


