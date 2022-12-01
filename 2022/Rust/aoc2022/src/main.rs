use aoc2022::input::read_file;
use aoc2022::day01::{day011, day012};

fn day01() {
    let filename: &str = "day01_input.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [1-1]: {}", day011(input1));
    println!("Answer [1-2]: {}", day012(input2));
}

fn main() {
    day01();
}
