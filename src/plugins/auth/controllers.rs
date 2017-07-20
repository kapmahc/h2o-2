use rocket_contrib::{Json, Value};

#[get("/api/site/info")]
pub fn get_site_info() -> Json<Value> {
    Json(json!({
                "status": "error",
                "reason": "ID exists. Try put."
    }))
}
