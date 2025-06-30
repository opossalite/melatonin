use rand::Rng;
use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};



#[derive(Debug, Deserialize)]
pub struct UpperReq {
    word: String,
}
#[derive(Serialize)]
pub struct UpperRes {
    word_upper: String,
}
//#[post("/upper", format = "json", data = "<msg>")] //NOTE: the msg here ties it to the msg in the parameters
#[post("/upper", data = "<msg>")] //NOTE: the msg here ties it to the msg in the parameters
pub fn upper(msg: Json<UpperReq>) -> Json<UpperRes> {
    Json(UpperRes {word_upper: msg.into_inner().word.to_uppercase()})
}



#[derive(Serialize)]
pub struct RollRes {
    value: u8
}
#[get("/roll")]
pub fn roll() -> Json<RollRes> {
    let rolled = rand::rng().random_range(1..7);
    Json(RollRes {value: rolled})
}


