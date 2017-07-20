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
    pub fn new() -> errors::Result<Config> {
        let host = try!(env::var("H2O_DATABASE_HOST"));
        let name = try!(env::var("H2O_DATABASE_NAME"));
        let user = try!(env::var("H2O_DATABASE_USER"));
        let password = try!(env::var("H2O_DATABASE_PASSWORD"));
        let port = try!(try!(env::var("H2O_DATABASE_PORT")).parse::<i32>());
        Ok(Config {
            host: host,
            port: port,
            name: name,
            user: user,
            password: password,
        })
    }

    fn exec(&self, sql: String) -> errors::Result<bool> {
        info!("Open database: postgres@{}:{}", self.host, self.port);
        println!("Please input password:");
        let mut pwd = String::new();
        try!(io::stdin().read_line(&mut pwd));
        pwd.pop();

        let db = try!(Connection::connect(
            format!("postgres://postgres:{}@{}:{}", pwd, self.host, self.port),
            TlsMode::None
        ));
        try!(db.execute(&sql, &[]));
        Ok(true)
    }

    pub fn create(&self) -> errors::Result<bool> {
        self.exec(format!(
            "CREATE DATABASE {} WITH ENCODING = 'UTF8';",
            self.name
        ))
    }

    pub fn drop(&self) -> errors::Result<bool> {
        self.exec(format!("DROP DATABASE {};", self.name))
    }



    pub fn open(&self) -> errors::Result<Connection> {
        let url = format!(
            "postgres://{}:{}@{}:{}/{}",
            self.user,
            self.password,
            self.host,
            self.port,
            self.name
        );
        debug!("Open database: {}", url);
        Ok(try!(Connection::connect(url, TlsMode::None)))
    }
}
