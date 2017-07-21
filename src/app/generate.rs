use std::path::Path;
use std::fs::{create_dir_all, OpenOptions};
use std::os::unix::fs::OpenOptionsExt;
use std::io::Write;

use time;
use mustache::{compile_path, MapBuilder};
use super::super::env::{config, errors};

pub fn config() -> errors::Result<bool> {
    let cfg = config::Config::new();
    try!(cfg.write("config"));
    Ok(true)
}

pub fn nginx(cfg: &config::Config) -> errors::Result<bool> {
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
        .insert_str("name", cfg.http.host())
        .insert_str("port", cfg.http.port())
        .insert_bool("ssl", cfg.http.ssl())
        .build();
    try!(tpl.render_data(&mut fd, &data));
    Ok(true)
}

pub fn locale(name: &str) -> errors::Result<bool> {
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

pub fn migration(name: &str) -> errors::Result<bool> {
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
