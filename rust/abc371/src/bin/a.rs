use proconio::input;

fn main() {
    input! {
        ab: char,
        ac: char,
        bc: char,
    }

    let ans = match (ab, ac, bc) {
        ('<', '>', _) | ('>', '<', _) => "A",
        ('>', _, '>') | ('<', _, '<') => "B",
        (_, '>', '<') | (_, '<', '>') => "C",
        _ => unreachable!(),
    };

    println!("{ans}");
}
