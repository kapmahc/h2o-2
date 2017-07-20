
use super::super::env;

pub fn run(name: &str, threads: u32) -> env::errors::Result<bool> {
    println!("Start worker {}[{}]", name, threads);
    Ok(true)
}
