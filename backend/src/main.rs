mod history;
mod albums;

#[macro_use] extern crate rocket;
use rocket::serde::json::Json;
use rocket_cors::CorsOptions;
use serde::{Deserialize, Serialize};
use rand::{self, Rng};

//#[get("/")]
//fn index() -> &'static str {
//    "Hello World!"
//}


#[derive(Debug, Deserialize)]
struct UpperReq {
    word: String,
}
#[derive(Serialize)]
struct UpperRes {
    word_upper: String,
}
#[post("/upper", format = "json", data = "<msg>")]
fn upper(msg: Json<UpperReq>) -> Json<UpperRes> {
    Json(UpperRes {word_upper: msg.into_inner().word.to_uppercase()})
}
//#[options("/upper")]
//fn upper_options() -> &'static str {
//    ""
//}

#[derive(Serialize)]
struct RollRes {
    value: u8
}
#[get("/roll")]
fn roll() -> Json<RollRes> {
    let rolled = rand::rng().random_range(1..7);
    Json(RollRes {value: rolled})
}
//#[options("/roll")]
//fn roll_options() -> &'static str { //don't question this, it's better this way
//    ""
//}

#[launch]
fn rocket() -> _ {
    let cors = CorsOptions::default()
        .to_cors()
        .expect("error building CORS");

    rocket::build()
        .configure(rocket::Config::figment().merge(("port", 8800)))
        //.mount("/", routes![roll, roll_options, upper, upper_options])
        .mount("/", routes![roll, upper])
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

