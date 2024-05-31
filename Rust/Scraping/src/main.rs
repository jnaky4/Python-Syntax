use reqwest::blocking::Client;
use select::document::Document;
use select::node::Node;
use select::predicate::{Name, Predicate}; // Add this import to bring Predicate into scope
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader, BufWriter, Write};

fn scrape_gameinformer() -> HashMap<String, String> {
    let base_url = "https://www.gameinformer.com";
    let routes = vec!["2024", "2023", "2022", "2021", "2020"];

    let mut game_releases = HashMap::new();

    for r in routes {
        let url = format!("{}/{}", base_url, r);
        let response = Client::new().get(&url).send();

        if let Ok(response) = response {
            if response.status().is_success() {
                let mut gr = HashMap::new();
                let document = Document::from_read(response).expect("Failed to parse response");

                for node in document.find(Name("span").and(|n: &Node| n.attr("class").unwrap_or("") == "calendar_entry")) {
                    let game_name_tag = node.find(Name("a")).next();
                    let release_date_tag = node.find(Name("time")).next();

                    if let (Some(game_name_tag), Some(release_date_tag)) = (game_name_tag, release_date_tag) {
                        let game_name = game_name_tag.text();
                        let release_date = release_date_tag.attr("datetime").unwrap_or("").to_string();
                        gr.insert(game_name, release_date);
                    }
                }
                game_releases.extend(gr);
            } else {
                println!("Failed to retrieve data.");
            }
        }
    }

    game_releases
}

fn write_to_file(game_releases: &HashMap<String, String>) {
    let file = File::create("game_releases.bin").expect("Failed to create file");
    let mut writer = BufWriter::new(file);

    bincode::serialize_into(&mut writer, game_releases).expect("Failed to write to file");

    println!("Game releases saved to 'game_releases.bin'");
}

fn read_from_file() -> HashMap<String, String> {
    let file = match File::open("game_releases.bin") {
        Ok(file) => file,
        Err(_) => {
            println!("No game releases found.");
            return HashMap::new();
        }
    };

    let reader = BufReader::new(file);

    bincode::deserialize_from(reader).unwrap_or_else(|_| {
        println!("Failed to read game releases.");
        HashMap::new()
    })
}
fn update_game_releases(old: &mut HashMap<String, String>, new: HashMap<String, String>) {
    old.extend(new);
}

fn main() {
    let mut old_releases = read_from_file();

    // for (game_name, release_date) in &old_releases {
    //     println!("Game: {}, Release Date: {}", game_name, release_date);
    // }


    let new_releases = scrape_gameinformer();

    update_game_releases(&mut old_releases, new_releases);
    write_to_file(&old_releases);
}
