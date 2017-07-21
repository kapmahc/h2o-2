use std::path::Path;
use std::vec::Vec;
use std::fs::{self, File};
use std::io::prelude::Read;
use std::collections::BTreeMap;

use time;
use postgres;

use super::super::env;

pub fn create(cfg: &env::config::Config) -> env::errors::Result<bool> {
    try!(cfg.postgresql.create());
    Ok(true)
}

pub fn drop(cfg: &env::config::Config) -> env::errors::Result<bool> {
    try!(cfg.postgresql.drop());
    Ok(true)
}

pub fn migrate(cfg: &env::config::Config) -> env::errors::Result<bool> {
    let con = try!(cfg.postgresql.open());
    try!(check(&con));
    let db = try!(con.transaction());

    for mig in try!(migrations()) {
        info!("Find migration {}", mig);
        let stmt = try!(db.prepare(
            "SELECT COUNT(*) FROM schema_migrations WHERE version = $1"
        ));
        let rows = try!(stmt.query(&[&mig]));
        let count: i64 = rows.get(0).get(0);
        if count == 0 {
            println!("Up migration {}", mig);
            let mut file = try!(File::open(
                Path::new("db")
                    .join("migrations")
                    .join(&mig)
                    .join("up")
                    .with_extension("sql")
            ));
            let mut sql = String::new();
            try!(file.read_to_string(&mut sql));
            try!(db.batch_execute(&sql));
            try!(db.execute(
                "INSERT INTO schema_migrations(version) VALUES($1)",
                &[&mig]
            ));
        }
    }
    try!(db.commit());
    Ok(true)
}

pub fn rollback(cfg: &env::config::Config) -> env::errors::Result<bool> {
    let con = try!(cfg.postgresql.open());
    try!(check(&con));
    let db = try!(con.transaction());

    let stmt = try!(db.prepare(
        "SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1"
    ));
    let rows = try!(stmt.query(&[]));
    if rows.len() == 0 {
        error!("Empty database.");
        return Ok(true);
    }
    let mig: String = rows.get(0).get(0);
    println!("Down migration {}", mig);

    let mut file = try!(File::open(
        Path::new("db")
            .join("migrations")
            .join(&mig)
            .join("down")
            .with_extension("sql")
    ));
    let mut sql = String::new();
    try!(file.read_to_string(&mut sql));
    try!(db.batch_execute(&sql));
    try!(db.execute(
        "DELETE FROM schema_migrations WHERE version = $1",
        &[&mig]
    ));

    try!(db.commit());
    Ok(true)
}

pub fn status(cfg: &env::config::Config) -> env::errors::Result<bool> {
    let db = try!(cfg.postgresql.open());
    try!(check(&db));

    let mut items = BTreeMap::new();
    for mig in try!(migrations()) {
        let stmt = try!(db.prepare(
            "SELECT created_at FROM schema_migrations WHERE version = $1 LIMIT 1"
        ));
        let rows = try!(stmt.query(&[&mig]));
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
    Ok(false)
}



fn migrations() -> env::errors::Result<Vec<String>> {
    let root = Path::new("db").join("migrations");
    let mut items = Vec::new();
    if root.is_dir() {
        for entry in try!(fs::read_dir(&root)) {
            let dir = try!(entry).path();
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

fn check(db: &postgres::Connection) -> env::errors::Result<()> {
    try!(db.execute(
        "CREATE TABLE IF NOT EXISTS schema_migrations(version VARCHAR(255) PRIMARY \
         KEY, created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now())",
        &[]
    ));
    Ok(())
}
