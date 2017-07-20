use std::{env, io};

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
        let password = try!(env::var("H2O_DATABASE_PASSWORD").map_err(errors::Error::EnvVar));
        let port = try!(try!(env::var("H2O_DATABASE_PORT").map_err(errors::Error::EnvVar))
            .parse::<i32>()
            .map_err(errors::Error::ParseInt));
        Ok(Config {
            host: host,
            port: port,
            name: name,
            user: user,
            password: password,
        })
    }

    fn exec(&self, sql: String) -> Result<bool, errors::Error> {
        info!("Open database: postgres@{}:{}", self.host, self.port);
        println!("Please input password:");
        let mut pwd = String::new();
        try!(io::stdin().read_line(&mut pwd).map_err(errors::Error::Io));
        pwd.pop();

        let db = try!(Connection::connect(format!("postgres://postgres:{}@{}:{}",
                                                  pwd,
                                                  self.host,
                                                  self.port),
                                          TlsMode::None)
            .map_err(errors::Error::PostgresConnect));
        try!(db.execute(&sql, &[]).map_err(errors::Error::Postgres));
        Ok(true)
    }

    pub fn create(&self) -> Result<bool, errors::Error> {
        self.exec(format!("CREATE DATABASE {} WITH ENCODING = 'UTF8';", self.name))
    }

    pub fn drop(&self) -> Result<bool, errors::Error> {
        self.exec(format!("DROP DATABASE {};", self.name))
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
