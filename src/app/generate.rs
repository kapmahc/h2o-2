use std::path::Path;
use std::fs::{create_dir_all, OpenOptions};
use std::os::unix::fs::OpenOptionsExt;
use std::io::Write;
use time;

use mustache::{compile_path, MapBuilder};
use super::super::env;

pub fn nginx() -> Result<bool, env::Error> {
    let root = "etc";
    try!(create_dir_all(root).map_err(env::Error::Io));
    let file = Path::new(root).join("nginx.conf");
    info!("Creating file {}", file.display());

    let mut fd = try!(OpenOptions::new()
        .write(true)
        .create_new(true)
        .mode(0o644)
        .open(file)
        .map_err(env::Error::Io));

    let tpl = compile_path(Path::new("templates").join("nginx.conf")).unwrap();
    let data = MapBuilder::new()
        .insert_str("hostname", "www.change-me.com")
        .insert("port", &3000)
        .expect("port")
        .insert_bool("ssl", false)
        .build();
    tpl.render_data(&mut fd, &data).unwrap();
    Ok(true)
}

pub fn locale(name: &str) -> Result<bool, env::Error> {
    let root = "locales";
    try!(create_dir_all(root).map_err(env::Error::Io));

    let file = Path::new(root).join(name).with_extension("yml");
    info!("Creating file {}", file.display());
    let mut fd = try!(OpenOptions::new()
        .write(true)
        .create_new(true)
        .mode(0o644)
        .open(file)
        .map_err(env::Error::Io));
    try!(writeln!(&mut fd, "{}:", name).map_err(env::Error::Io));
    Ok(true)
}

pub fn migration(name: &str) -> Result<bool, env::Error> {
    let now = time::now();
    let ts = try!(time::strftime("%Y%m%d%H%M%S", &now).map_err(env::Error::ParseTime));

    let root = Path::new("db").join("migrations").join(format!("{}-{}", ts, name));
    try!(create_dir_all(&root).map_err(env::Error::Io));
    let files = vec!["up", "down"];
    for n in &files {
        let file = Path::new(&root).join(n).with_extension("sql");
        info!("Creating file {}", file.display());
        try!(OpenOptions::new()
            .write(true)
            .create_new(true)
            .mode(0o600)
            .open(file)
            .map_err(env::Error::Io));
    }
    Ok(true)
}
