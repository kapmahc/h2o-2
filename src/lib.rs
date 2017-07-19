#![feature(plugin)]
#![plugin(docopt_macros)]
#![plugin(rocket_codegen)]

pub mod app;
pub mod env;

#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate docopt;
extern crate rocket;
