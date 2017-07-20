
use rand::{thread_rng, Rng};

pub fn random_string(len: usize) -> String {
    let s: String = thread_rng().gen_ascii_chars().take(len).collect();
    s
}
