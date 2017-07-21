use rocket_contrib::{Json, Value};

#[get("/site/info")]
pub fn get_site_info() -> Json<Value> {
    Json(json!({
                "status": "error",
                "reason": "ID exists. Try put."
    }))
}

#[post("/users/sign-in")]
pub fn post_users_sign_in() -> Json<Value> {
    Json(json!({
        "ok":true,
    }))
}

#[post("/users/sign-up")]
pub fn post_users_sign_up() -> Json<Value> {
    Json(json!({
        "ok":true,
    }))
}
