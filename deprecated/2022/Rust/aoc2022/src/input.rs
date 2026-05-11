use std::fs::File;
use std::io::prelude::*;
use std::path::Path;

pub fn read_file(filename: &str) -> String {
    let path = Path::new(filename);
    let expanded = path.display();

    let mut file = match File::open(&path) {
        Err(why) => panic!("Could not open {}: {}", expanded, why),
        Ok(file) => file,
    };

    let mut s = String::new();
    match file.read_to_string(&mut s) {
        Err(why) => panic!("Could not read {}: {}", expanded, why),
        Ok(_file) => s,
    }
}
