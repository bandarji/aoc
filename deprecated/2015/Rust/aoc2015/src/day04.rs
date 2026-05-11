use md5;

pub fn day041(input: String) -> i32 {
    answer04(input, 5)
}

pub fn day042(input: String) -> i32 {
    answer04(input, 6)
}

fn answer04(input: String, zeros: i32) -> i32 {
    // for n in (0..std::u64::MAX) {
    let mut z = String::new();
    for n in 0..zeros {
        z.push('0');
    }
    for n in 0..std::i32::MAX {
        let s = format!("{}{}", input, n);
        let h = format!("{:?}", md5::compute(s.as_bytes()));
        if h.to_string().starts_with(&z) {
            return n
        }
    }
    -1
}
