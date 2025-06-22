#[macro_use] extern crate rocket;
use rocket::serde::json::Json;
use rocket_cors::CorsOptions;
use serde::Serialize;
use rand::{self, Rng};

//#[get("/")]
//fn index() -> &'static str {
//    "Hello World!"
//}

#[derive(Serialize)]
struct RollResult {
    value: u8
}

#[post("/roll")]
fn roll() -> Json<RollResult> {
    let rolled = rand::rng().random_range(1..7);
    Json(RollResult {value: rolled})
}
#[options("/roll")]
fn roll_options() -> &'static str {
    "" // Needed for CORS preflight to succeed
}

#[launch]
fn rocket() -> _ {
    let cors = CorsOptions::default()
        .to_cors()
        .expect("error building CORS");

    rocket::build()
        .configure(rocket::Config::figment().merge(("port", 8800)))
        .mount("/", routes![roll])
        .attach(cors)
}



//#[derive(Serialize)]
//struct DiceResult {
//    value: u8,
//}
//
//#[post("/roll")]
//fn roll_die() -> Json<DiceResult> {
//    let rolled = rand::thread_rng().gen_range(1..=6);
//    Json(DiceResult { value: rolled })
//}
//
//#[launch]
//fn rocket() -> _ {
//    rocket::build().mount("/", routes![roll_die])
//}

