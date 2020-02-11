# check-payday

### Synopsis
Written in Rust! `check-payday` is a really simple and easy to use terminal application that tells
me if it's payday or not. Running `check-payday` will print out how many days until payday. Executing
in `--interactive` mode (or with the `-i` flag) will print out a progress bar that fills in as the
seconds count down towards the next payday!

### Installation
Dead simple because it's cargo. Install by running:
```bash
cargo install --git https://github.com/paul-nelson-baker/check-payday
```

### Usage
`$ check-payday --help`
```
check-payday 0.2
Paul Freakn Baker <paul.nelson.baker@gmail.com>
Prints out how many days until we get paid. Exit code will match

USAGE:
    check-payday [FLAGS] [OPTIONS]

FLAGS:
    -h, --help           Prints help information
    -i, --interactive    indicates whether or not to show the payday progress bar
    -V, --version        Prints version information

OPTIONS:
    -d, --base-date <base-date>    the date that calculations will be based off of. Format must be given in: %Y-%m-%d
                                   [default: 2018-01-05]
```
