use std::fs::{File, OpenOptions};
use std::io::prelude::Read;
use std::io::Write;
use std::os::unix::fs::OpenOptionsExt;
use std::path::Path;

use toml;
use super::errors;

#[derive(Serialize, Deserialize, Debug)]
pub struct Config {
    http: Http,
    secrets: Secrets,
    postgresql: PostgreSql,
    redis: Redis,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Http {
    port: i32,
    ssl: bool,
    host: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct PostgreSql {
    host: String,
    port: i32,
    name: String,
    user: String,
    password: String,
    ssl_mode: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Redis {
    host: String,
    port: i32,
    db: i8,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Secrets {
    jwt: String,
    aes: String,
    hmac: String,
}

impl Config {
    pub fn new() -> Config {
        Config {
            http: Http {
                port: 8080,
                ssl: false,
                host: "www.change-me.com".to_string(),
            },
            secrets: Secrets {
                jwt: "j".to_string(),
                aes: "a".to_string(),
                hmac: "h".to_string(),
            },
            postgresql: PostgreSql {
                host: "localhost".to_string(),
                port: 5432,
                name: "h2o_dev".to_string(),
                user: "postgres".to_string(),
                password: "".to_string(),
                ssl_mode: "disable".to_string(),
            },
            redis: Redis {
                host: "localhost".to_string(),
                port: 6379,
                db: 9,
            },
        }
    }
    pub fn read(name: &'static str) -> errors::Result<Config> {
        let file = Path::new(name).with_extension("toml");
        info!("Reading from {}", file.display());
        let mut fd = try!(File::open(file));
        let mut buf = String::new();
        try!(fd.read_to_string(&mut buf));
        let cfg: Config = try!(toml::from_str(&buf));
        Ok(cfg)
    }

    pub fn write(&self, name: &'static str) -> errors::Result<String> {
        let file = Path::new(name).with_extension("toml");
        println!("Creating file {}", file.display());
        let mut fd = try!(
            OpenOptions::new()
                .write(true)
                .create_new(true)
                .mode(0o644)
                .open(&file)
        );

        let txt = try!(toml::to_string(&self));
        try!(writeln!(&mut fd, "{}", txt));
        Ok(format!("{}", file.display()))
    }
}
