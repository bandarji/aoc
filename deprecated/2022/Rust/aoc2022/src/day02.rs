pub fn day021(input: String) -> i32 {
    answer021(input)
}

pub fn day022(input: String) -> i32 {
    answer022(input)
}

fn answer021(input: String) -> i32 {
    let mut score: i32 = 0;
    for line in input.split('\n') {
        if line.len() < 3 {
            return score;
        }
        // let ch = line.chars().nth(0).unwrap();
        let win_loss = if line == "A Z" || line == "B X" || line == "C Y" {
            0
        } else if line == "A X" || line == "B Y" || line == "C Z" {
            3
        } else {
            6
        };
        let myself = &line[2..3];
        let point = if myself == "X" {
            1
        } else if myself == "Y" {
            2
        } else {
            3
        };
        score += win_loss + point
    }
    score
}

fn answer022(input: String) -> i32 {
    let mut score: i32 = 0;
    for line in input.split('\n') {
        if line.len() < 3 {
            return score;
        }
        let strategy = if line == "A Y" || line == "B X" || line == "C Z" {
            1
        } else if line == "A Z" || line == "B Y" || line == "C X" {
            2
        } else {
            3
        };
        let myself = &line[2..3];
        let point = if myself == "X" {
            0
        } else if myself == "Y" {
            3
        } else {
            6
        };
        score += strategy + point
    }
    score
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer021() {
        assert!(answer021("A Y\nB X\nC Z".to_string()) == 15);
    }

    #[test]
    fn test_answer022() {
        assert!(answer022("A Y\nB X\nC Z".to_string()) == 12);
    }
}
