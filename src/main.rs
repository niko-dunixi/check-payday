extern crate chrono;
extern crate clap;

use crate::chrono::Datelike;
use crate::chrono::Timelike;
use chrono::{Duration, Local, NaiveDate, NaiveDateTime, NaiveTime};
use clap::{App, Arg};
use indicatif::{ProgressBar, ProgressStyle};
use std::error::Error;
use std::ops::{Add, Sub};
use std::{thread, time};
// https://mattgathu.github.io/writing-cli-app-rust/

fn main() -> Result<(), Box<dyn Error>> {
    // Set the CLI options
    let matches = App::new("check-payday")
        .version("0.2")
        .author("Paul Freakn Baker <paul.nelson.baker@gmail.com>")
        .about("Prints out how many days until we get paid. Exit code will match")
        .arg(Arg::with_name("interactive")
            .short("i")
            .long("interactive")
            .takes_value(false)
            .help("indicates whether or not to show the payday progress bar"))
        .arg(Arg::with_name("base-date")
            .short("d")
            .long("base-date")
            .takes_value(true)
            .default_value("2018-01-05")
            .help("the date that calculations will be based off of. Format must be given in: %Y-%m-%d"))
        .get_matches();

    // Start performing the actual date-maths
    let base_date_string = matches.value_of("base-date").ok_or("base-date was bad?")?;
    let base_date = NaiveDate::parse_from_str(base_date_string, "%Y-%m-%d")?;
    let today_date = date_today();
    let tuple = payday_tuple(base_date, today_date);

    // Actually handle the output to ourselves
    let is_interactive = matches.occurrences_of("interactive") > 0;
    if is_interactive {
        // If we want something interactive, we create a progress bar!
        let midnight = NaiveTime::from_hms(0, 0, 0);
        let date_time_a = tuple.0.and_time(midnight);
        let date_time_b = tuple.1.and_time(midnight);
        let total_seconds = date_time_b.sub(date_time_a).num_seconds() as u64;

        let pb = ProgressBar::new(total_seconds);
        pb.set_prefix("🕰");
        pb.set_message("seconds to next payday... 👀 ");
        pb.set_style(
            ProgressStyle::default_bar().template("{prefix} {pos:>7}/{len:7} {msg} [{wide_bar:40.cyan/blue}]")
                .progress_chars("=> ")
        );
        let five_hundred_millis = time::Duration::from_millis(500);
        while datetime_today() < date_time_b {
            let current_seconds = datetime_today().sub(date_time_a).num_seconds() as u64;
            pb.set_position(current_seconds);
            thread::sleep(five_hundred_millis);
        }
        println!("IT'S PAYDAY!! 😃🎉🎉🎉🎉")
    } else {
        // Print out how many days are left until payday!
        let days = tuple.1.sub(today_date).num_seconds() / 60 / 60 / 24;
        if days > 0 {
            println!("{} days until we get paid 💸", days);
        } else {
            println!("0 days until we get paid. Wait, IT'S PAYDAY!! 😃🎉🎉🎉🎉")
        }
    }
    Ok(())
}

fn payday_tuple(base_date: NaiveDate, today_date: NaiveDate) -> (NaiveDate, NaiveDate) {
    let fortnight_duration = Duration::days(14);
    let mut next_payday = base_date;
    while next_payday < today_date {
        next_payday = next_payday.add(fortnight_duration);
    }
    let last_payday = next_payday.sub(fortnight_duration);
    return (last_payday, next_payday);
}

fn date_today() -> NaiveDate {
    let local = Local::today();
    NaiveDate::from_ymd(local.year(), local.month(), local.day())
}

fn datetime_today() -> NaiveDateTime {
    let date = date_today();
    let now = Local::now();
    date.and_hms(now.hour(), now.minute(), now.second())
}
