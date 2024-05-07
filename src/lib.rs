use std::ffi::{CStr, CString};
use std::os::raw::c_char;

#[no_mangle]
pub extern "C" fn greet(name: *const c_char) -> *mut c_char {
    let c_str_name = unsafe { CStr::from_ptr(name) };
    let recipient = match c_str_name.to_str() {
        Ok(str) => str,
        Err(_) => "there",
    };
    let greeting = format!("Hello, {}!", recipient);
    CString::new(greeting).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn free_string(s: *mut c_char) {
    unsafe {
        if s.is_null() {
            return;
        }
        CString::from_raw(s)
    };
}
