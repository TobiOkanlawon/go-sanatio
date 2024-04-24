# go-sanatio

Go-sanatio is a lightweight validation library written in Golang.

I began it as a way to validate a simple web application that I was making.

# Implementation

Go-sanatio is implemented a validation chain. You create a validator and then tack on the validations you want onto the validation chain.

# Roadmap

- Early returns. It'll be really nice to have early returns. The ability to have the validation stop if it fails at any point.

- More kinds of inbuilt validators. We currently only have strings, but that's because that's what I am working with currently. As I get more use cases, I will make more validators
