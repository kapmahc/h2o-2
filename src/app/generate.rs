use std::path::Path;
use std::fs::OpenOptions;
use std::os::unix::fs::OpenOptionsExt;

use mustache::{compile_path, MapBuilder};

pub fn nginx() {
    let file = Path::new("etc").join("nginx.conf");
    println!("Creating file {}", file.display());
    let mut fd = OpenOptions::new()
        .write(true)
        .create_new(true)
        .mode(0o644)
        .open(file)
        .unwrap();

    let tpl = compile_path(Path::new("templates").join("nginx.conf")).unwrap();
    let data = MapBuilder::new()
        .insert_str("hostname", "www.change-me.com")
        .insert("port", &3000)
        .expect("port")
        .insert_bool("ssl", false)
        .build();
    tpl.render_data(&mut fd, &data).unwrap();
}

pub fn ssl() {}

pub fn locale() {}

pub fn migration() {}
