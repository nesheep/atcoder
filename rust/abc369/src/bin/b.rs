use proconio::input;

fn ini_pos(cmds: &[(i64, char)], s: char) -> i64 {
    cmds.iter().find(|(_, c)| *c == s).unwrap_or(&(0, s)).0
}

fn main() {
    input! {
        n: usize,
        cmds: [(i64, char); n],
    }

    let l_ini = ini_pos(&cmds, 'L');
    let r_ini = ini_pos(&cmds, 'R');

    let ans = cmds
        .iter()
        .fold(((l_ini, r_ini), 0), |((l, r), ans), &(a, s)| match s {
            'L' => ((a, r), ans + (l - a).abs()),
            'R' => ((l, a), ans + (r - a).abs()),
            _ => unreachable!(),
        })
        .1;

    println!("{ans}");
}
