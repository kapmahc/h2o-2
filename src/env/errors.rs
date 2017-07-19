use std::{io, num, error, fmt, env};

use time;
use mustache;
use postgres;

#[derive(Debug)]
pub enum Error {
    Io(io::Error),
    ParseInt(num::ParseIntError),
    EnvVar(env::VarError),
    ParseTime(time::ParseError),
    Template(mustache::Error),
    Postgres(postgres::error::Error),
    PostgresConnect(postgres::error::ConnectError),
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Error::Io(ref err) => write!(f, "IO error: {}", err),
            Error::ParseInt(ref err) => write!(f, "Parse error: {}", err),
            Error::ParseTime(ref err) => write!(f, "Parse time error: {}", err),
            Error::Template(ref err) => write!(f, "Template error: {}", err),
            Error::EnvVar(ref err) => write!(f, "ENV var error: {}", err),
            Error::Postgres(ref err) => write!(f, "PGSQL error: {}", err),
            Error::PostgresConnect(ref err) => write!(f, "PGSQL connect error: {}", err),
        }
    }
}

impl error::Error for Error {
    fn description(&self) -> &str {
        match *self {
            Error::Io(ref err) => err.description(),
            Error::ParseInt(ref err) => err.description(),
            Error::ParseTime(ref err) => err.description(),
            Error::Template(ref err) => err.description(),
            Error::EnvVar(ref err) => err.description(),
            Error::Postgres(ref err) => err.description(),
            Error::PostgresConnect(ref err) => err.description(),
        }
    }

    fn cause(&self) -> Option<&error::Error> {
        match *self {
            Error::Io(ref err) => Some(err),
            Error::ParseInt(ref err) => Some(err),
            Error::ParseTime(ref err) => Some(err),
            Error::Template(ref err) => Some(err),
            Error::EnvVar(ref err) => Some(err),
            Error::Postgres(ref err) => Some(err),
            Error::PostgresConnect(ref err) => Some(err),
        }
    }
}
