use std::io;
use postgres::{Connection, TlsMode};
use super::errors;


#[derive(Serialize, Deserialize, Debug)]
pub struct PostgreSql {
    host: String,
    port: i32,
    name: String,
    user: String,
    password: String,
}

impl PostgreSql {
    pub fn new() -> PostgreSql {
        PostgreSql {
            host: "localhost".to_string(),
            port: 5432,
            name: "h2o_dev".to_string(),
            user: "postgres".to_string(),
            password: "".to_string(),
        }
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
