use rocket::{custom, config};
use super::super::env;
use super::super::plugins::auth::controllers as auth;

pub fn run(cfg: &env::config::Config, worker: bool) -> env::errors::Result<bool> {
    if worker {
        // TODO
    }

    let cfg = try!(
        config::Config::build(cfg.env())
            .address("localhost")
            .port(cfg.http.port())
            .finalize()
    );

    custom(cfg, false)
        .mount(
            "/api",
            routes![
                auth::get_site_info,
                auth::post_users_sign_in,
                auth::post_users_sign_up,
            ],
        )
        .launch();
    Ok(true)
}
