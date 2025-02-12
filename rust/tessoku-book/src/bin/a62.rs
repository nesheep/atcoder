use proconio::input;

fn do_dfs(graph: &Vec<Vec<usize>>, v: usize, visited: &mut Vec<bool>) {
    visited[v] = true;
    for &u in graph[v].iter() {
        if visited[u] {
            continue;
        }
        do_dfs(graph, u, visited);
    }
}

fn dfs(graph: &Vec<Vec<usize>>, n: usize) -> Vec<bool> {
    let mut visited = vec![false; n];
    do_dfs(graph, 1, &mut visited);
    visited
}

fn main() {
    input! {
        n: usize,
        m: usize,
        edges: [(usize, usize); m],
    }

    let mut graph = vec![vec![]; n + 1];
    for &(a, b) in edges.iter() {
        graph[a].push(b);
        graph[b].push(a);
    }

    let visited = dfs(&graph, n + 1);
    let ans = visited[1..].iter().all(|&x| x);

    println!(
        "{}",
        if ans {
            "The graph is connected."
        } else {
            "The graph is not connected."
        }
    );
}
