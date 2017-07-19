pub mod db;
pub mod cache;
pub mod errors;

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
