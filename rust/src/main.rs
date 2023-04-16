use std::fs::{DirBuilder};
use std::result::Result;

fn create_directory_structure(sorting_name: &str, release_name: &str) -> Result<(), std::io::Error> {
    let full_path = String::from(sorting_name) + "/" + release_name;
    DirBuilder::new()
        .recursive(true)
        .create(&full_path).unwrap();

    let subdirs = [
        "setup",
        "dlc",
        "extras",
        "updates",
        "instructions"
    ];

    for dir in subdirs {
        let path = full_path.clone() + "/" + dir;
        DirBuilder::new()
            .recursive(true)
            .create(&path)?;
        println!("Created: {}", path);
    }

    Ok(())
}

fn main() {
    let args: Vec<_> = std::env::args().collect();
    let program = args.last().unwrap().split('/').last().unwrap();
    if args.len() != 3  {
        println!("Usage: {} (Rust) sorting_name release_name", program);
        return;
    }

    match create_directory_structure(args.get(1).unwrap(), args.get(2).unwrap()) {
        Ok(_) => println!("Success"),
        Err(e) => println!("{:?}", e),
    }

    return;
}
