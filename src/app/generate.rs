use std;
use std::path::Path;
use std::fs::{create_dir_all, OpenOptions};
use std::os::unix::fs::OpenOptionsExt;
use std::io::Write;
use time;

use mustache::{compile_path, MapBuilder};
use super::super::env;

pub fn config() -> env::errors::Result<bool> {
    let cfg = env::config::Config::new();
    try!(cfg.write("config"));
    Ok(true)
}

pub fn nginx() -> env::errors::Result<bool> {
    let root = "etc";
    try!(create_dir_all(root));
    let file = Path::new(root).join("nginx.conf");
    println!("Creating file {}", file.display());

    let mut fd = try!(
        OpenOptions::new()
            .write(true)
            .create_new(true)
            .mode(0o644)
            .open(file)
    );

    let tpl = try!(compile_path(Path::new("templates").join("nginx.conf")));
    let data = MapBuilder::new()
        .insert_str("name", try!(std::env::var("H2O_SERVER_NAME")))
        .insert_str(
            "port",
            try!(try!(std::env::var("ROCKET_PORT")).parse::<u32>()),
        )
        .insert_bool(
            "ssl",
            try!(try!(std::env::var("H2O_SERVER_SSL")).parse::<bool>()),
        )
        .build();
    try!(tpl.render_data(&mut fd, &data));
    Ok(true)
}

pub fn locale(name: &str) -> env::errors::Result<bool> {
    let root = "locales";
    try!(create_dir_all(root));

    let file = Path::new(root).join(name).with_extension("yml");
    println!("Creating file {}", file.display());
    let mut fd = try!(
        OpenOptions::new()
            .write(true)
            .create_new(true)
            .mode(0o644)
            .open(file)
    );
    try!(writeln!(&mut fd, "{}:", name));
    Ok(true)
}

pub fn migration(name: &str) -> env::errors::Result<bool> {
    let now = time::now();
    let ts = try!(time::strftime("%Y%m%d%H%M%S", &now));

    let root = Path::new("db")
        .join("migrations")
        .join(format!("{}-{}", ts, name));
    try!(create_dir_all(&root));
    let files = vec!["up", "down"];
    for n in &files {
        let file = Path::new(&root).join(n).with_extension("sql");
        println!("Creating file {}", file.display());
        try!(
            OpenOptions::new()
                .write(true)
                .create_new(true)
                .mode(0o600)
                .open(file)
        );
    }
    Ok(true)
}
