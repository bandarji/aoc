use std::collections::HashSet;

#[derive(Copy, Clone, Hash, Debug, Eq, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

pub fn day031(input: String) -> i32 {
    answer031(input)
}

pub fn day032(input: String) -> i32 {
    answer032(input)
}

fn answer031(input: String) -> i32 {
    let mut positions: HashSet<Point> = HashSet::new();
    let mut current: Point = Point {x: 0, y: 0};
    positions.insert(Point {x: 0, y: 0});
    for char in input.chars() {
        current = if char == '^' {
            Point {x: current.x + 0, y: current.y + 1}
        } else if char == 'v' {
            Point {x: current.x + 0, y: current.y - 1}
        } else if char == '>' {
            Point {x: current.x + 1, y: current.y + 0}
        } else if char == '<' {
            Point {x: current.x - 1, y: current.y + 0}
        } else {
            Point {x: current.x + 0, y: current.y + 0}
        };
        positions.insert(Point {x: current.x, y: current.y});
    }
    positions.len() as i32
}

fn answer032(input: String) -> i32 {
    let mut houses: HashSet<Point> = HashSet::new();
    let mut robot: Point = Point {x: 0, y: 0};
    let mut santa: Point = Point {x: 0, y: 0};
    let mut turn: i32 = 0;
    houses.insert(robot);
    for instruction in input.chars() {
        if turn % 2 == 0 {
            robot = new_position(robot.x, robot.y, instruction);
            houses.insert(robot);
        } else {
            santa = new_position(santa.x, santa.y, instruction);
            houses.insert(santa);
        }
        turn += 1;
    };
    houses.len() as i32
}

fn new_position(x: i32, y: i32, instruction: char) -> Point {
    if instruction == '^' {
        Point {x: x + 0, y: y + 1}
    } else if instruction == 'v' {
        Point {x: x + 0, y: y - 1}
    } else if instruction == '>' {
        Point {x: x + 1, y: y + 0}
    } else if instruction == '<' {
        Point {x: x - 1, y: y + 0}
    } else {
        Point {x: x, y: y}
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer031() {
        assert!(answer031(String::from(">")) == 2);
        assert!(answer031(String::from("^>v<")) == 4);
        assert!(answer031(String::from("^v^v^v^v^v")) == 2);
    }

    #[test]
    fn test_answer032() {
        assert!(answer032(String::from("^v")) == 3);
        assert!(answer032(String::from("^>v<")) == 3);
        assert!(answer032(String::from("^v^v^v^v^v")) == 11);
    }
    // #[test]
    // fn test_answer021_getdims() {
    //     assert!(answer021_getdims("2x3x4".to_string()) == [2, 3, 4])
    // }
}
