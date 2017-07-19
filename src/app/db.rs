
use super::super::env;

pub fn create() -> Result<bool, env::errors::Error> {
    try!(try!(env::db::Config::new()).create());
    Ok(true)
}

pub fn drop() -> Result<bool, env::errors::Error> {
    try!(try!(env::db::Config::new()).drop());
    Ok(true)
}

pub fn migrate() -> Result<bool, env::errors::Error> {
    // let db = try!(try!(env::db::Config::new()).open());
    Ok(true)
}

pub fn rollback() -> Result<bool, env::errors::Error> {
    Ok(true)
}

pub fn status() -> Result<bool, env::errors::Error> {
    Ok(true)
}


pub fn connect() -> Result<bool, env::errors::Error> {
    Ok(true)
}
