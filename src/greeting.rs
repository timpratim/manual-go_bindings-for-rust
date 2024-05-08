pub fn send_greeting(recipient: &str) -> String {
    let greeting = format!("Greetings from Rustonia to {}!", recipient);
    greeting
}