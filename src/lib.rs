#![feature(plugin)]
#![plugin(docopt_macros)]
#![plugin(rocket_codegen)]

pub mod app;
pub mod env;
pub mod plugins;

extern crate time;
#[macro_use]
extern crate log;
#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate rocket_contrib;
extern crate serde;
extern crate docopt;
extern crate rocket;
extern crate mustache;
extern crate rustc_serialize;
extern crate postgres;
extern crate toml;
extern crate rand;
