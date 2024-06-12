use dioxus::prelude::*;

pub fn Loading() -> Element {
    rsx! {
        div { class: "absolute z-50 w-full h-full flex flex-col items-center justify-center bg-primary-500",
            Throbbler {}
        }
    }
}

pub fn Throbbler() -> Element {
    rsx! {
        div { class: "oxygen bg-primary-500 rounded-full",
            div { class: "orbit outer-orbit",
                div { class: "electron one" }
                div { class: "electron two" }
                div { class: "electron three" }
                div { class: "electron four" }
                div { class: "electron five" }
                div { class: "electron six" }
                div { class: "orbit inner-orbit",
                    div { class: "electron seven" }
                    div { class: "electron eight" }
                    div { class: "core" }
                }
            }
        }
    }
}
