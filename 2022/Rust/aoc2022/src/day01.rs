pub fn day011(input: String) -> i32 {
    answer011(input)
}

pub fn day012(input: String) -> i32 {
    answer012(input)
}

fn answer011(input: String) -> i32 {
    let mut elves: Vec<i32> = Vec::new();
    let mut elf: i32 = 0;
    for line in input.split('\n') {
        let calories: i32 = line.trim().parse().unwrap_or(0);
        if calories == 0 {
            elves.push(elf);
            elf = 0;
        } else {
            elf += calories;
        }
    }
    elves.push(elf);
    elves.sort();
    elves.reverse();
    elves[0]
}

fn answer012(input: String) -> i32 {
    let mut elves: Vec<i32> = Vec::new();
    let mut elf: i32 = 0;
    for line in input.split('\n') {
        let calories: i32 = line.trim().parse().unwrap_or(0);
        if calories == 0 {
            elves.push(elf);
            elf = 0;
        } else {
            elf += calories;
        }
    }
    elves.push(elf);
    elves.sort();
    elves.reverse();
    elves[0] + elves[1] + elves[2]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer011() {
        assert!(answer011("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000".to_string()) == 24000);
    }

    #[test]
    fn test_answer012() {
        assert!(answer012("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000".to_string()) == 45000);
    }
}
