pub fn day011(input: String) -> i32 {
    answer011(input)
}

pub fn day012(input: String) -> i32 {
    answer012(input)
}

fn answer011(input: String) -> i32 {
    let mut floor: i32 = 0;
    for char in input.chars() {
        floor += if char == '(' {
            1
        } else if char == ')' {
            -1
        } else {
            0
        }
    }
    floor
}

fn answer012(input: String) -> i32 {
    let mut position: i32 = 0;
    let mut floor: i32 = 0;
    for char in input.chars() {
        position += 1;
        let delta: i32 = if char == '(' {
            1
        } else if char == ')' {
            -1
        } else {
            0
        };
        floor += delta;
        if floor == -1 {
            return position;
        };
    }
    -1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer011() {
        assert!(answer011("(())".to_string()) == 0);
        assert!(answer011("()()".to_string()) == 0);
        assert!(answer011("(((".to_string()) == 3);
        assert!(answer011("(()(()(".to_string()) == 3);
        assert!(answer011("))(((((".to_string()) == 3);
        assert!(answer011("())".to_string()) == -1);
        assert!(answer011("))(".to_string()) == -1);
        assert!(answer011(")))".to_string()) == -3);
        assert!(answer011(")())())".to_string()) == -3);
    }

    #[test]
    fn test_answer012() {
        assert!(answer012(")".to_string()) == 1);
        assert!(answer012("()())".to_string()) == 5);
    }
}
