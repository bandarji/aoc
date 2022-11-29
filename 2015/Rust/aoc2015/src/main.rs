use aoc2015::input::read_file;
use aoc2015::day01::{day011, day012};
use aoc2015::day02::{day021, day022};

fn day01() {
    let filename: &str = "day01_input.txt";
    let input = read_file(filename);
    println!("Answer [1-1]: {}", day011(input));
    let input = read_file(filename);
    println!("Answer [1-2]: {}", day012(input));
}

fn day02() {
    let filename: &str = "day02_input.txt";
    let input = read_file(filename);
    println!("Answer [2-1]: {}", day021(input));
    let input = read_file(filename);
    println!("Answer [2-2]: {}", day022(input));
}

fn main() {
    day01();
    day02();
}
