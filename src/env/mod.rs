pub mod db;
pub mod cache;

use std::io;
use std::num;
use std::error;
use std::fmt;

pub fn version() -> String {
    format!(
        "{}/{}\n{}\n{}\n{}",
        env!("CARGO_PKG_NAME").to_uppercase(),
        env!("CARGO_PKG_VERSION"),
        env!("CARGO_PKG_AUTHORS"),
        env!("CARGO_PKG_HOMEPAGE"),
        env!("CARGO_PKG_DESCRIPTION"),
    )
}

#[derive(Debug)]
pub enum Error {
    Io(io::Error),
    Parse(num::ParseIntError),
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Error::Io(ref err) => write!(f, "IO error: {}", err),
            Error::Parse(ref err) => write!(f, "Parse error: {}", err),
        }
    }
}

impl error::Error for Error {
    fn description(&self) -> &str {
        match *self {
            Error::Io(ref err) => err.description(),
            Error::Parse(ref err) => err.description(),
        }
    }

    fn cause(&self) -> Option<&error::Error> {
        match *self {
            Error::Io(ref err) => Some(err),
            Error::Parse(ref err) => Some(err),
        }
    }
}
