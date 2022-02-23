# Go-Kong Bug 
This is a sample Go project to simulate a detected bug in go-kong v0.28.0.

## When the bug happens
When we try to validate a plugin schema using go-kong and the target kong is offline (or even if the Admin API was disabled), we get a reference error.
