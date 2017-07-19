use std::env;

use postgres::{Connection, TlsMode};

use super::errors;

pub struct Config {
    host: String,
    port: i32,
    name: String,
    user: String,
    password: String,
}

impl Config {
    pub fn new() -> Result<Config, errors::Error> {
        let host = try!(env::var("H2O_DATABASE_HOST").map_err(errors::Error::EnvVar));
        let name = try!(env::var("H2O_DATABASE_NAME").map_err(errors::Error::EnvVar));
        let user = try!(env::var("H2O_DATABASE_USER").map_err(errors::Error::EnvVar));
        Ok(Config {
            host: host,
            port: 5432,
            name: name,
            user: user,
            password: "".to_string(),
        })
    }

    fn exec(&self, password: &'static str, sql: String) -> Result<bool, errors::Error> {
        debug!("Open database: postgres@{}:{}", self.host, self.port);
        let db = try!(Connection::connect(format!("postgres://postgres:{}@{}:{}",
                                                  password,
                                                  self.host,
                                                  self.port),
                                          TlsMode::None)
            .map_err(errors::Error::PostgresConnect));
        try!(db.execute(&sql, &[]).map_err(errors::Error::Postgres));
        Ok(true)
    }

    pub fn create(&self, password: &'static str) -> Result<bool, errors::Error> {
        self.exec(password,
                  format!("CREATE DATABASE {} WITH ENCODING = 'UTF8';", self.name))
    }

    pub fn drop(&self, password: &'static str) -> Result<bool, errors::Error> {
        self.exec(password, format!("DROP DATABASE {};", self.name))
    }

    pub fn open(&self) -> Result<Connection, errors::Error> {
        let url = format!("postgres://{}:{}@{}:{}/{}",
                          self.user,
                          self.password,
                          self.host,
                          self.port,
                          self.name);
        debug!("Open database: {}", url);
        Ok(try!(Connection::connect(url, TlsMode::None).map_err(errors::Error::PostgresConnect)))
    }
}
