# golox

golox is a Tree Walk interpreter written in Go.

### Why?

Well, interpreters are interesting. Maybe not tree walk interpreters themeselves (I implemented one for Shell a while back in C), but I'm looking at this as way to play around a bit with Go.
<br />
I'm reading [Crafting Interpreters](https://craftinginterpreters.com/), which implements it in Java - this prospect didn't exactly have me leaping from enthusiasm, hence Go.
<br />
I don't expect this little project to expose me to the most interesting Go features (gorountines, for instance), but anyways.

### Next steps

Once this is done the plan is to look more into bytecode interpreters, before eventually graduating to the big-boy league of Compilers.

### Notes

#### AST Design

The book implements the AST and its traversal with the Visitor Pattern: each node is a derived class from an abstract expression, each implementing the Accept interface. This Accept interface defines a function (aptly) called `accept`, which
takes a Visitor as an argument. This Visitor is itself an interface which defines n methods called accept{nodeType}, where nodeType is the type of the node to visit. This design pattern ensures proper separation of concerns by having all of the logic
associated with a Visitor implementor's operations on a certain node all in one place.<br />
The Visitor pattern plays well with Java, but the lack of explicit inheritance and "normal" dynamic dispatch one would expect from an OOP language in Go makes this implementation a little less natural.
<br />
Instead, I'll opt for defining an Expr interface with just one marker method, with no operation.

#### Rationale

- From my brief reading on Go's type system and interface representation, in memory an interface implementation is what in Rust would be called a "fat pointer": it's essentialy a two word struct, where the first word is a pointer to the implementation information
of the interface (type information + vtable, if any methods are defined there), and the second word is either an unsafe pointer to the actual data, or a copy of the data itself (apparently the Go compiler is able to nicely optimize data locality for small values).
- Given that Go doesn't really lend itself to inheritance and that it supports type-switching at run time, the interface approach **_without the indirection of the Visitor pattern's methods_** seems more idiomatic (and slightly more performant).
- Having a marker method in our trait provides more type safety (an empty interface is just an `any`, afterall), and marginally better performance when resolving the types at run-time.
- Because I don't expect to add much functionality to the AST (will likely only want to either evaluate the nodes or pretty-print them for debugging), it feels like the extensibility provided by the Visitor pattern hardly justifies the extra allocation and boilerplate for our use-case.
