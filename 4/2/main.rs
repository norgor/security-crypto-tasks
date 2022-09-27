fn html_escape(s: &str) -> String {
    s.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
}

fn main() {
    let mut line = String::new();
    std::io::stdin().read_line(&mut line).unwrap();

    let escaped = html_escape(&line);
    println!("{}", escaped)
}
