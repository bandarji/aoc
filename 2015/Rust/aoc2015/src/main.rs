use aoc2015::input::read_file;
use aoc2015::day01::{day011, day012};
use aoc2015::day02::{day021, day022};
use aoc2015::day03::{day031, day032};
use aoc2015::day04::{day041, day042};

fn day01() {
    let filename: &str = "day01_input.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [1-1]: {}", day011(input1));
    println!("Answer [1-2]: {}", day012(input2));
}

fn day02() {
    let filename: &str = "day02_input.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [2-1]: {}", day021(input1));
    println!("Answer [2-2]: {}", day022(input2));
}

fn day03() {
    let filename: &str = "day03_input.txt";
    let input1 = read_file(filename);
    let input2 = input1.clone();
    println!("Answer [3-1]: {}", day031(input1));
    println!("Answer [3-2]: {}", day032(input2));
}

fn day04() {
    println!("Answer [4-1]: {}", day041(String::from("yzbqklnj")));
    println!("Answer [4-2]: {}", day042(String::from("yzbqklnj")));
}

fn main() {
    day01();
    day02();
    day03();
    day04();
}
