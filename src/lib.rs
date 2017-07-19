#![feature(plugin)]
#![plugin(docopt_macros)]
#![plugin(rocket_codegen)]

pub mod app;
pub mod env;

extern crate time;
#[macro_use]
extern crate log;
#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate docopt;
extern crate rocket;
extern crate mustache;
extern crate rustc_serialize;
extern crate postgres;
