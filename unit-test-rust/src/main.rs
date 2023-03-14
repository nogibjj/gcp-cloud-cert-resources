/*run all functions in lib.rs */

use unit_test_rust::add;
use unit_test_rust::div;
use unit_test_rust::mul;
use unit_test_rust::sub;

fn main() {
    println!("Hello, world!");
    println!("add(1, 2) = {}", add(1, 2));
    println!("sub(1, 2) = {}", sub(1, 2));
    println!("mul(1, 2) = {}", mul(1, 2));
    println!("div(1, 2) = {}", div(1, 2));
}