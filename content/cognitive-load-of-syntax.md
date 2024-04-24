---
name: "Thoughts on Cognitive Load and Programming Language Syntax"
date: "February 07 2024"
tags: ["thoughts"]
---
Recently, I started to pick up a new language in my spare time: Go. Was it partially influenced by the surge in HTMX popularity and how often Go is used alongside it? Almost certainly. But at some point along this journey, I noticed something: This is _so_ much more fun and _so_ much less draining. This whole post won't be me gushing about how much I love Go as a language, not directly. I started to notice a few things...

# Oh, what fun...
I come from the JavaScript and TypeScript world with frameworks like Remix and Next.js. And the DX (Developer Experience) for these is lovely! For reasons I could not quite pin down, I was having much more fun with Go + HTMX despite the hacks to get close to the same DX. It came down to how darn easily I could pull off things I had previously relied on services for. Need a message queue to run on the edge? Easy, write a wrapper over Redis to use its PubSub. Need real-time updates to the UI? No problem, write a basic Server-Sent Events system with subscriptions and channels. Anything I wanted, I could build. That, and the way concurrency _"just works"_, and I could make anything I wanted happen. This brings a level of satisfaction that is hard to convey in an article.

# What changed?
So this isn't meant to be a post to bash JavaScript, as much as I do have my reservations about the language. That being said, JS makes basic things like concurrency an afterthought. Go's lightweight threads, called goroutines, make it easy to write concurrent programs without worrying about the overhead of heavyweight threads. Even better, Go has channels that provide a simple and effective way of synchronizing communication between goroutines, making it easier to write concurrent programs free from race conditions. Similar concepts are present in Java, Rust (Tokio), C#, etc. Even better, error handling! Now, this is controversial, but forcing errors to be valued and allowing functions to assert that there _may_ be errors is a game changer for writing safer code. All of this amounts to a much lower cognitive load. I am less worried that my code is inefficient, error-prone, etc. I can just _write_.

# What does this mean?
As a field, we can and should investigate and identify key points in programming languages that increase or decrease the mental load or overhead for developers. Imagine if we could identify these points and work to reduce the overhead. This has been on my mind as I reflect on my experience learning Go. I want to dive deeper into the details of what we can do to reduce this overhead and get back to writing good software without unnecessary restrictions.

# Thanks!
Do you have any thoughts on this? Do you want to have a chat about the topic? Feel free to reach out by email at [me@atri.dad](mailto:me@atri.dad). Until next time! 🫡