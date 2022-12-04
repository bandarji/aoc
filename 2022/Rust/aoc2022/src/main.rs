use aoc2022::input::read_file;
use aoc2022::day01::{day011, day012};
use aoc2022::day02::{day021, day022};
use aoc2022::day03::{day031, day032};

fn day01() {
    let filename: &str = "day01_input.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [1-1]: {}", day011(input1));
    println!("Answer [1-2]: {}", day012(input2));
}

fn day02() {
    let filename: &str = "input_day02.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [2-1]: {}", day021(input1));
    println!("Answer [2-2]: {}", day022(input2));
}

fn day03() {
    let filename: &str = "input_day03.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [3-1]: {}", day031(input1));
    println!("Answer [3-2]: {}", day032(input2));
}

fn main() {
    day01();
    day02();
    day03();
}
