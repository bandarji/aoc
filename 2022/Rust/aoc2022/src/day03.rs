pub fn day031(input: String) -> i32 {
    answer031(input)
}

pub fn day032(input: String) -> i32 {
    answer032(input)
}

fn answer031(input: String) -> i32 {
    let mut total: i32 = 0;
    for line in input.split('\n') {
        let midpoint = line.len() / 2;
        let (left, right) = line.split_at(midpoint);
        let duped_item = find_duped_item(left, right);
        total += item_value(duped_item);
    }
    total
}

fn answer032(input: String) -> i32 {
    let mut total: i32 = 0;
    let mut rucksacks: Vec<&str> = input.split('\n').collect();
    rucksacks.retain(|s| s.len() > 1);

    for (i, _) in rucksacks.iter().enumerate() {
        if i % 3 == 0 {
            let common: char = find_common_item(
                rucksacks[i as usize],
                rucksacks[(i + 1) as usize],
                rucksacks[(i + 2) as usize]
            );
            total += item_value(common);
        }
    }
    total
}

fn find_common_item(r1: &str, r2: &str, r3: &str) -> char {
    // let r2_chars: Vec<char> = r2.chars().collect();
    for letter in r1.chars() {
        if r2.contains(letter) && r3.contains(letter) {
            return letter;
        }
    }
    '∅'
}

fn find_duped_item(left: &str, right: &str) -> char {
    for l_left in left.chars() {
        for l_right in right.chars() {
            if l_left == l_right {
                return l_left;
            }
        }
    }
    '∅'
}

fn item_value(key: char) -> i32 {
    let alphabet: &str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut i: i32 = 0;
    for letter in alphabet.chars() {
        i += 1;
        if key == letter {
            return i;
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer031() {
        assert!(answer031("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw".to_string()) == 157);
    }

    #[test]
    fn test_answer032() {
        assert!(answer032("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw".to_string()) == 70);
    }

}
