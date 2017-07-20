use std::path::Path;
use std::vec::Vec;
use std::fs::{self, File};
use std::io::prelude::Read;
use std::collections::BTreeMap;

use time;
use postgres;

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
    let con = try!(open());
    let db = try!(con.transaction().map_err(env::errors::Error::Postgres));

    for mig in try!(migrations()) {
        info!("Find migration {}", mig);
        let stmt = try!(db.prepare("SELECT COUNT(*) FROM schema_migrations WHERE version = $1")
            .map_err(env::errors::Error::Postgres));
        let rows = try!(stmt.query(&[&mig]).map_err(env::errors::Error::Postgres));
        let count: i64 = rows.get(0).get(0);
        if count == 0 {
            let mut file = try!(File::open(Path::new("db")
                    .join("migrations")
                    .join(&mig)
                    .join("up")
                    .with_extension("sql"))
                .map_err(env::errors::Error::Io));
            let mut sql = String::new();
            try!(file.read_to_string(&mut sql).map_err(env::errors::Error::Io));
            try!(db.batch_execute(&sql).map_err(env::errors::Error::Postgres));
            try!(db.execute("INSERT INTO schema_migrations(version) VALUES($1)", &[&mig])
                .map_err(env::errors::Error::Postgres));
        }
    }
    try!(db.commit().map_err(env::errors::Error::Postgres));
    Ok(true)
}

pub fn rollback() -> Result<bool, env::errors::Error> {
    let con = try!(open());
    let db = try!(con.transaction().map_err(env::errors::Error::Postgres));
    let stmt =
        try!(db.prepare("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1")
            .map_err(env::errors::Error::Postgres));
    let rows = try!(stmt.query(&[]).map_err(env::errors::Error::Postgres));
    if rows.len() == 0 {
        error!("Empty database.");
        return Ok(true);
    }
    let mig: String = rows.get(0).get(0);
    info!("Find migration {}", mig);

    let mut file = try!(File::open(Path::new("db")
            .join("migrations")
            .join(&mig)
            .join("down")
            .with_extension("sql"))
        .map_err(env::errors::Error::Io));
    let mut sql = String::new();
    try!(file.read_to_string(&mut sql).map_err(env::errors::Error::Io));
    try!(db.batch_execute(&sql).map_err(env::errors::Error::Postgres));
    try!(db.execute("DELETE FROM schema_migrations WHERE version = $1", &[&mig])
        .map_err(env::errors::Error::Postgres));

    try!(db.commit().map_err(env::errors::Error::Postgres));
    Ok(true)
}

pub fn status() -> Result<bool, env::errors::Error> {
    let db = try!(open());
    let mut items = BTreeMap::new();
    for mig in try!(migrations()) {
        let stmt =
            try!(db.prepare("SELECT created_at FROM schema_migrations WHERE version = $1 LIMIT 1")
                .map_err(env::errors::Error::Postgres));
        let rows = try!(stmt.query(&[&mig]).map_err(env::errors::Error::Postgres));
        if rows.len() == 0 {
            items.insert(mig, "None".to_string());
        } else {
            let created: time::Timespec = rows.get(0).get(0);
            items.insert(mig, format!("{}", time::at(created).rfc822()));
        }
    }

    println!("{:<32}\t{}", "NAME", "TIMESTAMP");
    for (k, v) in items {
        println!("{:<32}\t{}", k, v);
    }
    Ok(true)
}



fn migrations() -> Result<Vec<String>, env::errors::Error> {
    let root = Path::new("db").join("migrations");
    let mut items = Vec::new();
    if root.is_dir() {
        for entry in try!(fs::read_dir(&root).map_err(env::errors::Error::Io)) {
            let dir = try!(entry.map_err(env::errors::Error::Io)).path();
            if dir.is_dir() {
                match dir.file_name() {
                    Some(n) => {
                        match n.to_str() {
                            Some(s) => items.push(s.to_string()),
                            None => {}
                        }
                    }
                    None => {}
                }
            }
        }
    }
    items.sort();
    Ok(items)
}

fn open() -> Result<postgres::Connection, env::errors::Error> {
    let db = try!(try!(env::db::Config::new()).open());
    try!(db.execute("CREATE TABLE IF NOT EXISTS schema_migrations(version VARCHAR(255) PRIMARY \
                  KEY, created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now())",
                 &[])
        .map_err(env::errors::Error::Postgres));
    Ok(db)
}
