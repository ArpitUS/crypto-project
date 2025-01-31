use sha2::{Sha256, Digest};
use hex;

pub fn hash_data(input: &str) -> String {
    let mut hasher = Sha256::new();
    hasher.update(input);
    hex::encode(hasher.finalize())
}
