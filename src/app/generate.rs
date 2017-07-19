use std::path::Path;
use std::fs::OpenOptions;
use std::os::unix::fs::OpenOptionsExt;

use mustache::{compile_path, MapBuilder};
use super::super::env;

pub fn nginx() -> Result<bool, env::Error> {
    let file = Path::new("etc").join("nginx.conf");
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

pub fn ssl() {}

pub fn locale() {}

pub fn migration() {}
