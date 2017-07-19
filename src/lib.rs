#![feature(plugin)]
#![plugin(docopt_macros)]

pub mod app;
pub mod env;

#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate docopt;
