pub fn day041(input: String) -> i32 {
    answer041(input)
}

pub fn day042(input: String) -> i32 {
    answer042(input)
}

fn answer041(input: String) -> i32 {
    let mut pairs: i32 = 0;
    for line in input.split('\n') {
        if line.len() < 3 {
            continue
        }
        let mut entry = line.split(|t| t == ',' || t == '-');
        let e1_start = entry.next().unwrap().parse::<i32>().unwrap();
        let e1_end = entry.next().unwrap().parse::<i32>().unwrap();
        let e2_start = entry.next().unwrap().parse::<i32>().unwrap();
        let e2_end = entry.next().unwrap().parse::<i32>().unwrap();
        pairs += if (e1_start - e2_start) * (e1_end - e2_end) <= 0 {
            1
        } else {
            0
        };
    }
    pairs
}

fn answer042(input: String) -> i32 {
    let mut pairs: i32 = 0;
    for line in input.split('\n') {
        if line.len() < 3 {
            continue
        }
        let mut entry = line.split(|t| t == ',' || t == '-');
        let e1_start = entry.next().unwrap().parse::<i32>().unwrap();
        let e1_end = entry.next().unwrap().parse::<i32>().unwrap();
        let e2_start = entry.next().unwrap().parse::<i32>().unwrap();
        let e2_end = entry.next().unwrap().parse::<i32>().unwrap();
        pairs += if (e2_start <= e1_start && e1_start <= e2_end) || (e1_start <= e2_start && e2_start <= e1_end) {
            1
        } else {
            0
        };
    }
    pairs
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer041() {
        assert!(answer041("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8".to_string()) == 2);
    }

    #[test]
    fn test_answer042() {
        assert!(answer042("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8".to_string()) == 4);
    }

}
