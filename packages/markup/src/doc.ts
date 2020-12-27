/**
 
 Markup is still under heavy development but it is close to release with some changes still needed!*
 
 Markup is a project geared towards providing a DOM interaction layer between different
 projects built either in WebAssembly, Zig, Golang or Rust. Interacting with the DOM is at times
 not as performant or somewhat impossible due to the nature of whatever language being used.
 
 I think by creating a simple interface library like this, developers can add whatever features
 and methods desired using Markup, which allows them to have a consistent means of interacting both
 with DOM nodes, diffing and event management.
 
 Markup goal is to be the sit in library that any project can include for this purpose, I plan to ensure
 it provides the following:
 
 - Efficient DOM diffing using html strings and JSON.
 - Easy interaction with RequestAnimationFrame for event loops.
 - Simple Event management using live events.
 
 */
