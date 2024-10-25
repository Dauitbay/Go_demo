Go has a number of built-in types, including numeric ones (byte, int64, float32, etc.), Booleans, and byte strings (string). Strings are immutable; built-in operators and keywords (rather than functions) provide concatenation, comparison, and UTF-8 encoding/decoding.Record types can be defined with the struct keyword.

For each type T and each non-negative integer constant n, there is an array type denoted [n]T; arrays of differing lengths are thus of different types. Dynamic arrays are available as "slices", denoted []T for some type T. These have a length and a capacity specifying when new memory needs to be allocated to expand the array. Several slices may share their underlying memory.

Pointers are available for all types, and the pointer-to-T type is denoted *T. Address-taking and indirection use the & and * operators, as in C, or happen implicitly through the method call or attribute access syntax.There is no pointer arithmetic,[e] except via the special unsafe.Pointer type in the standard library.

For a pair of types K, V, the type map[K]V is the type mapping type-K keys to type-V values, though Go Programming Language specification does not give any performance guarantees or implementation requirements for map types. Hash tables are built into the language, with special syntax and built-in functions. chan T is a channel that allows sending values of type T between concurrent Go processes.

Aside from its support for interfaces, Go's type system is nominal: the type keyword can be used to define a new named type, which is distinct from other named types that have the same layout (in the case of a struct, the same members in the same order). Some conversions between types (e.g., between the various integer types) are pre-defined and adding a new type may define additional conversions, but conversions between named types must always be invoked explicitly. For example, the type keyword can be used to define a type for IPv4 addresses, based on 32-bit unsigned integers as follows:

type ipv4addr uint32
With this type definition, ipv4addr(x) interprets the uint32 value x as an IP address. Simply assigning x to a variable of type ipv4addr is a type error.

Constant expressions may be either typed or "untyped"; they are given a type when assigned to a typed variable if the value they represent passes a compile-time check.

Function types are indicated by the func keyword; they take zero or more parameters and return zero or more values, all of which are typed. The parameter and return values determine a function type; thus, func(string, int32) (int, error) is the type of functions that take a string and a 32-bit signed integer, and return a signed integer (of default width) and a value of the built-in interface type error.

Any named type has a method set associated with it. The IP address example above can be extended with a method for checking whether its value is a known standard:

// ZeroBroadcast reports whether addr is 255.255.255.255.
func (addr ipv4addr) ZeroBroadcast() bool {
    return addr == 0xFFFFFFFF
}
Due to nominal typing, this method definition adds a method to ipv4addr, but not on uint32. While methods have special definition and call syntax, there is no distinct method type
