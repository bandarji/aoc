pub fn day021(input: String) -> i32 {
    answer021(input)
}

pub fn day022(input: String) -> i32 {
    answer022(input)
}

fn answer021(input: String) -> i32 {
    let mut sqft: i32 = 0;
    // let mut dims: Vec<i32> = [0, 0, 0].to_vec();
    for line in input.split('\n') {
        let dims: Vec<i32> = answer021_getdims(line.to_string());
        sqft += answer021_onebox(dims[0], dims[1], dims[2]);
    }
    sqft
}

fn answer021_onebox(l: i32, w: i32, h: i32) -> i32 {
    let mut dims: [i32; 3] = [l, w, h];
    dims.sort();
    dims[0] * dims[1] + (2 * l * w) + (2 * w * h) + (2 * l * h)
}

fn answer021_getdims(s: String) -> Vec<i32> {
    if s.len() < 2 {
        return [0, 0, 0].to_vec();
    }
    s.split('x').map(|x| x.parse::<i32>().unwrap()).collect()
}

fn answer022(input: String) -> i32 {
    let mut ribbon: i32 = 0;
    for line in input.split('\n') {
        let mut dims: Vec<i32> = answer021_getdims(line.to_string());
        dims.sort();
        let shortest_perimeter: i32 = 2 * (dims[0] + dims[1]);
        let volume: i32 = dims[0] * dims[1] * dims[2];
        ribbon += shortest_perimeter + volume
    }
    ribbon
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_answer021_onebox() {
        assert!(answer021_onebox(2, 3, 4) == 58);
        assert!(answer021_onebox(1, 1, 10) == 43);
    }

    #[test]
    fn test_answer021_getdims() {
        assert!(answer021_getdims("2x3x4".to_string()) == [2, 3, 4])
    }
}
