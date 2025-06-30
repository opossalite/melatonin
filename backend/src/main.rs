mod history;
mod get_albums;
mod experimentation;
mod albums;

#[macro_use] extern crate rocket;
use experimentation::{roll, upper};
use rocket::serde::json::Json;
use rocket_cors::CorsOptions;
use serde::{Deserialize, Serialize};
use rand::{self, Rng};



#[launch]
fn rocket() -> _ {
    let cors = CorsOptions::default()
        .to_cors()
        .expect("error building CORS");

    rocket::build()
        .configure(rocket::Config::figment().merge(("port", 8800)))
        //.mount("/", routes![roll, roll_options, upper, upper_options])
        .mount("/", routes![roll, upper, get_albums::get_albums])
        .attach(cors)
}


