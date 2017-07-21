use std::fs::{File, OpenOptions};
use std::io::prelude::Read;
use std::io::Write;
use std::os::unix::fs::OpenOptionsExt;
use std::path::Path;
use std::str::FromStr;
use std::env;

use log::LogLevelFilter;
use toml;
use rocket::config::Environment;
use super::errors;

pub fn log_level() -> LogLevelFilter {
    let lv = LogLevelFilter::max();
    match env::var("RUST_LOG") {
        Ok(lvl) => {
            match LogLevelFilter::from_str(&lvl) {
                Ok(v) => v,
                Err(_) => lv,
            }
        }
        Err(_) => lv,
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Config {
    env: String,
    pub http: Http,
    pub secrets: Secrets,
    pub postgresql: super::db::PostgreSql,
    pub redis: Redis,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Http {
    port: u16,
    ssl: bool,
    host: String,
}

impl Http {
    pub fn host(&self) -> &str {
        &self.host
    }
    pub fn port(&self) -> u16 {
        self.port
    }
    pub fn ssl(&self) -> bool {
        self.ssl
    }
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
    cors: String,
}

impl Secrets {}

impl Config {
    pub fn new() -> Config {
        Config {
            env: format!("{}", Environment::Development),
            http: Http {
                port: 8080,
                ssl: false,
                host: "www.change-me.com".to_string(),
            },
            secrets: Secrets {
                jwt: super::utils::random_string(32),
                aes: super::utils::random_string(32),
                hmac: super::utils::random_string(32),
                cors: super::utils::random_string(32),
            },
            postgresql: super::db::PostgreSql::new(),
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

    pub fn is_prod(&self) -> bool {
        match Environment::from_str(&self.env) {
            Ok(e) => e.is_prod(),
            Err(_) => false,
        }
    }

    pub fn env(&self) -> Environment {
        match self.env.parse::<Environment>() {
            Ok(e) => e,
            Err(_) => Environment::Development,
        }
    }
}
