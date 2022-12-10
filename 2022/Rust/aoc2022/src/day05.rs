use std::collections::VecDeque; // Should have used deque for Python too

struct Move {
    count: usize,
    from: usize,
    to: usize,
}

pub fn day051(input: String) -> String {
    answer051(input, 8, 9)
}

pub fn day052(input: String) -> String {
    answer052(input)
}

impl Move {
    fn from_line(line: &str) -> Self {
        let mut s = line.split(|c| c == ' ');
        let count = s.nth(1).unwrap().parse::<usize>().unwrap();
        let from = s.nth(1).unwrap().parse::<usize>().unwrap();
        let to = s.nth(1).unwrap().parse::<usize>().unwrap();
        Self { count, from, to }
    }
}

// fn answer051x(input: String) -> String {
//     let mut stacks: Vec<VecDeque<u8>> = vec![VecDeque::new(); COLUMNS as usize];
//     input.lines().take(8).for_each(|line| {
//         let mut boxes = [None; header + 1];
//         for (index, position) in (1..line.len()).step_by(4).enumerate() {
//             let a_box = line.as_bytes()[position];
//             boxes[index] = if a_box == b' ' {
//                 None
//             } else {
//                 Some(a_box)
//             };
//         }
//         for (index, a_box) in boxes.into_iter().enumerate()
//                                  .filter(|(_, a_box)| a_box.is_some()) {
//                                     stacks[index].push_back(a_box.unwrap());
//                                  }
//     });

//     input.lines().skip((header + 2) as usize).map(Move::from_line).for_each(|m| {
//         for _ in 0..m.count {
//             let a_box = stacks[m.from - 1].pop_front().unwrap();
//             stacks[m.to - 1].push_front(a_box);
//         }
//     });
//     let mut answer = String::with_capacity(stacks.len());
//     stacks.iter().for_each(|s| answer.push(*s.front().unwrap() as char));
//     answer
// }

// fn answer052x(input: String) -> String {
//     "0".to_string()
// }

fn answer051(input: String, header: usize, columns: usize) -> String {
    let mut stacks: Vec<VecDeque<u8>> = vec![VecDeque::new(); 1];
    for _ in 0..columns {
        stacks.push(VecDeque::new());
    }
    input.lines().take(header).for_each(|line| {
        let mut boxes = [None; 9];
        for (index, position) in (1..line.len()).step_by(4).enumerate() {
            let a_box = line.as_bytes()[position];
            boxes[index] = if a_box == b' ' {
                None
            } else {
                Some(a_box)
            };
        }
        for (index, a_box) in boxes.into_iter().enumerate()
            .filter(|(_, a_box)| a_box.is_some()) {
                stacks[index].push_back(a_box.unwrap());
            }
    });
    input.lines().skip((header + 2) as usize).map(Move::from_line).for_each(|m| {
        for _ in 0..m.count {
            let a_box = stacks[m.from - 1].pop_front().unwrap();
            stacks[m.to - 1].push_front(a_box);
        }
    });
    let mut answer = String::with_capacity(stacks.len());
    stacks.iter().for_each(|s| answer.push(*s.front().unwrap() as char));
    // "CMZ".to_string()
    answer
}

fn answer052(input: String) -> String {
    "CMZ".to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const TEST_INPUT: &str = "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2";

    #[test]
    fn test_answer051() {
        assert!(answer051(TEST_INPUT.to_string(), 3, 3) == "CMZ");
    }
}
