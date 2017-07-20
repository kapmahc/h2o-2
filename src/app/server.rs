
// use rocket::custom;
// use rocket::config::{Config, Environment};

use super::super::{env, plugins};

pub fn run(worker: bool) -> env::errors::Result<bool> {
    if worker {
        // TODO
    }
    // "staging".parse::<Environment>()
    //
    // let cfg = try!(
    //     Config::build(Environment::Development)
    //         .address("localhost")
    //         .port(try!(try!(std::env::var("ROCKET_PORT")).parse::<u32>()))
    //         .finalize()
    // );
    //
    // rocket::custom(cfg, false)
    //     .mount("/", routes![plugins::auth::controllers::get_site_info])
    //     .launch();
    Ok(true)
}
